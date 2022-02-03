package utils

import(
	"time"
)

type ImageJSON struct {
	Architecture string `json:"architecture"`
	Config       struct {
		Hostname     string      `json:"Hostname"`
		Domainname   string      `json:"Domainname"`
		User         string      `json:"User"`
		Attachstdin  bool        `json:"AttachStdin"`
		Attachstdout bool        `json:"AttachStdout"`
		Attachstderr bool        `json:"AttachStderr"`
		Tty          bool        `json:"Tty"`
		Openstdin    bool        `json:"OpenStdin"`
		Stdinonce    bool        `json:"StdinOnce"`
		Env          []string    `json:"Env"`
		Cmd          []string    `json:"Cmd"`
		Image        string      `json:"Image"`
		Volumes      interface{} `json:"Volumes"`
		Workingdir   string      `json:"WorkingDir"`
		Entrypoint   interface{} `json:"Entrypoint"`
		Onbuild      interface{} `json:"OnBuild"`
		Labels       interface{} `json:"Labels"`
	} `json:"config"`
	Container       string `json:"container"`
	ContainerConfig struct {
		Hostname     string      `json:"Hostname"`
		Domainname   string      `json:"Domainname"`
		User         string      `json:"User"`
		Attachstdin  bool        `json:"AttachStdin"`
		Attachstdout bool        `json:"AttachStdout"`
		Attachstderr bool        `json:"AttachStderr"`
		Tty          bool        `json:"Tty"`
		Openstdin    bool        `json:"OpenStdin"`
		Stdinonce    bool        `json:"StdinOnce"`
		Env          []string    `json:"Env"`
		Cmd          []string    `json:"Cmd"`
		Image        string      `json:"Image"`
		Volumes      interface{} `json:"Volumes"`
		Workingdir   string      `json:"WorkingDir"`
		Entrypoint   interface{} `json:"Entrypoint"`
		Onbuild      interface{} `json:"OnBuild"`
		Labels       struct {
		} `json:"Labels"`
	} `json:"container_config"`
	Created       time.Time `json:"created"`
	DockerVersion string    `json:"docker_version"`
	History       []struct {
		Author     string    `json:"author"`
		Created    time.Time `json:"created"`
		CreatedBy  string    `json:"created_by"`
		EmptyLayer bool      `json:"empty_layer,omitempty"`
	} `json:"history"`
	Os     string `json:"os"`
	Rootfs struct {
		Type    string   `json:"type"`
		DiffIds []string `json:"diff_ids"`
	} `json:"rootfs"`
}