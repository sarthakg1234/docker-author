package main

/***
 * @author sarthak.goyal@sap.com
 */

import (
	"docker-modifier/utils"
	"flag"
	"fmt"
	"log"
)

func main() {
	var imageName, authorName string 
	flag.StringVar(&imageName, "image", "", "image name")
	flag.StringVar(&authorName, "author", "", "author name")
	flag.Parse()
	tarName := fmt.Sprintf("%s.tar", imageName)
	modifiedTarName := fmt.Sprintf("%s_modified", imageName)
	defer utils.CleanDirectory(imageName)
	saveToDocker(imageName, tarName)
	unTarFile(tarName, imageName)
	modifyAndWriteJSONData(imageName, authorName)
	createNewTar(imageName, modifiedTarName)
	loadTarToDocker(modifiedTarName)
}

func saveToDocker(imageName, tarName string) {
	err := utils.SaveTarFromDockerImage(imageName, tarName)
	if err != nil {
		log.Fatalf("Error occured while saving image to file %s Error is %s", tarName, err)
	}
}

func unTarFile(tarName, imageName string) {
	err := utils.Untar(tarName, imageName)
	if err != nil {
		log.Fatalf("Error occured while unarchiving file %s Error is %s", tarName, err)
	}
}

func modifyAndWriteJSONData(imageName string, authorName string) {
	jsonFiles, err := utils.FilteJSONFilesIn(imageName)
	if err != nil {
		log.Fatalf("Error occured while filtering json files in %s folder. Error is %s", imageName, err)
	}
	jsonData, err := utils.GetJsonDataFrom(jsonFiles[0])
	if err != nil {
		log.Fatalf("Error occured while unmarshaling file contents for file %s : %s", jsonFiles[0], err)
	}
	err = utils.EditJsonData(jsonData, authorName)
	if err != nil {
		log.Fatalf("Error occured while editing file contents for file %s : %s", jsonFiles[0], err)
	}
	err = utils.WriteNewJsonDataTo(jsonFiles[0], jsonData)
	if err != nil {
		log.Fatalf("Error occured while writing file contents for file %s : %s", jsonFiles[0], err)
	}
}

func createNewTar(imageName, modifiedTarName string) {
	err := utils.CreateTarFromFolder(imageName, modifiedTarName)
	if err != nil {
		log.Fatalf("Error occured while archiving file contents for file %s : %s", imageName, err)
	}
}

func loadTarToDocker(modifiedTarName string) {
	output, err := utils.LoadTarToDocker(modifiedTarName)
	if err != nil {
		log.Fatalf("Error occured while loading to docker %s : %s", modifiedTarName, err)
	} else {
		log.Println(output)
	}
}
