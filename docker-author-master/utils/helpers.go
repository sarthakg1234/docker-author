package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

func FilteJSONFilesIn(root string) ([]string, error) {
	var matches []string
	re := regexp.MustCompile(`\b[A-Fa-f0-9]{64}\b\.json`)
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if !re.MatchString(path) {
			return err
		}
		matches = append(matches, path)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return matches, nil
}

func GetJsonDataFrom(file string) (*ImageJSON, error) {
	fileContents, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("Error occured while reading file %s : %s", file, err)
	}
	jsonData := &ImageJSON{}
	err = json.Unmarshal(fileContents, &jsonData)
	return jsonData, err
}

func EditJsonData(jsonData *ImageJSON, authorName string) error {
	for index, _ := range jsonData.History {
		jsonData.History[index].Author = authorName
	}
	return nil
	// if value, ok := jsonData["history"]; ok {
	// 	jsonData["history"] = editHistory(value, authorName)
	// } else {
	// 	return nil, fmt.Errorf("Couldn't find the history")
	// }
	// return jsonData, nil
}

func editHistory(value interface{}, authorName string) interface{} {
	historySlice := value.([]interface{})
	for index := 0; index < len(historySlice); index++ {
		mapOfHistory := historySlice[index].(map[string]interface{})
		if _, ok := mapOfHistory["author"]; ok {
			continue
		}
		mapOfHistory["author"] = authorName
		historySlice[index] = mapOfHistory
	}
	return historySlice
}
