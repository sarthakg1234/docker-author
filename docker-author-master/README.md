# docker-author

This is an interactive tool to add the docker-author field with the the corresponding author name in the docker history field in config.json inside docker image.

Use the binary based on the host operating system among MACOS, LINUX and WINDOWS.

## Usage

using binary
```
chmod 777 $binary
./$binary -image $name-of-docker-image -author $name-of-author
```
using go-lang

```
go run main.go -image $name-of-docker-image -author $name-of-author
```

### Detailed explanation for MACOS

1. Pull a docker image from the repository

```
docker pull $name-of-docker-image
```
![docker-pull](https://github.wdf.sap.corp/sfappsec/docker-author/blob/master/images/docker_pull.png "docker pull")

2. Execute the binary to add the docker author tag

```
chmod 777 $binary
./$binary -image $name-of-docker-image -author $name-of-author
```
![execute-binary](https://github.wdf.sap.corp/sfappsec/docker-author/blob/master/images/docker_author.png "execute binary")

### Verify Solution

1. Save the docker image as a tar file

```
docker save $name-of-docker-image -o $name-of-tar.tar
```
![docker-save](https://github.wdf.sap.corp/sfappsec/docker-author/blob/master/images/tar_from_docker.png "save docker image")

2. Untar the tar file 

```
tar -xvf $name-of-tar.tar
```
![untar](https://github.wdf.sap.corp/sfappsec/docker-author/blob/master/images/untar.png "untar the file")


3. Read the $hash.json file inside after untar to see the modification inside history tag

```
cat $hash.json | jq .
```
![display_output](https://github.wdf.sap.corp/sfappsec/docker-author/blob/master/images/display_output.png "Output display")

![final_output](https://github.wdf.sap.corp/sfappsec/docker-author/blob/master/images/added_author_tag.png "Output")




