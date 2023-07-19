## Prototype - Swagger

<br>

Translations:

* [繁體中文](README.md)

---

<br>

# Introduction

- Location of Swagger Specification 

<br><br>

# Organizational

- Swagger UI Server independent from Api Server, reduces unnecessary setups for deploying services

<br><br>

# CORS

- Because of Swagger UI and Api Server isn't the same domain, there are cors problem
- Using nginx for CORS

<br><br>

# Generate Swagger Specification 

- [Swaggo/swag](https://github.com/swaggo/swag)

- Installing


1.change directory to cmd (main.go in directory cmd)

```console
cd cmd
```

2.install [Swaggo/swag](https://github.com/swaggo/swag)

```console
 go get -u github.com/swaggo/swag/cmd/swag 
 go install github.com/swaggo/swag/cmd/swag
```

3.after installing, can type below command for check is completed

```console
swag --version
```

4.Go back to root directory of Api Server project

```console
cd ../
```
 
5.At entrance of main function and all Api handler，write [Swaggo/swag](https://github.com/swaggo/swag) format comment


```go
// @title Prototype Api Server
// @version 1.0
// @description Standard Api Server

// @contact.name sigma
// @contact.url https://github.com/sigmalin/Prototype

// @host 127.0.0.1:443
// @schemes https
func main() {
}
```

6.Type below commond，generate swagger docs to ```backend/swagger/```

```console
swag init -g cmd/main.go -o ../swagger/doc/
```

<br><br>

# Code generate form swagger

- [Swagger Codegen](https://github.com/swagger-api/swagger-codegen)

- Running by docker image

- File will save in convert folder

| SHELL | Feature |
|:-:|:--|
|openapi|convert openAPI 2.0 to openAPI 3.0|
|csharp|generate csharp client source code|

command

```console
sh openapi.sh
```

<br><br>

# TODO

- The csharp code generated can't be used in prototype frontend (need to install extra packages & variable name is not clear), so we keep the conversion function for now and expand it later.
    
