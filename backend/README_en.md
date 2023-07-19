## Prototype - backend

<br>

Translations:

* [繁體中文](README.md)

---

<br>

# Introduction

- Backend architecture to provide the necessary backend functions for game development.

<br><br>

# Folders

| Folders | feature |
|:--|:--|
|api|Api Server source code|
|bundles|Provide a location for front-end download resources|
|file|FTP Server source code, providing front-end download bundles|
|nginx|nginx config，for solving cors problem|
|sql|database data and schema|
|ssl|ssl files|
|swagger|swagger scripts|

<br><br>

# How to use

- Installing docker

- type ``` docker-compose up -d ``` , and then press enter to start all services

<br><br>

# docker-compose.yml

- go_api

| Folders | Usage |
|:--|:--|
|API_PORT|Api Server Port|
|SSL_CERTIFICATION|ssl location|
|SSL_PRIVATE_KEY|ssl location|
|JWT_SIGNING_KEY|Signature for Json web token|
|JWT_CLAIMS_KEY|Access key for Json web token claims data|
|SQL_DRIVER|SQL Driver|
|SQL_USERNAME|Username for SQL|
|SQL_PASSWORD|Password for SQL|
|SQL_ADDRESS|SQL IP Address|
|SQL_PORT|SQL Port|
|SQL_DATABASE|table name of SQL|
|SQL_TIMEOUT|Max time(sec) for SQL command wait|
|SQL_MAXLIFETIME|Max time(sec) for SQL connect|
|SQL_MAXOPENCONNECT|Max count for SQL connect|
|SQL_MAXIDLECONNECT|Max count for idle SQL connect|
|REDIS_CACHE_ADDRESS|Redis cache IP Address|
|REDIS_CACHE_PORT|Redis cache Port|
|REDIS_CACHE_EXPIRATION|Life time(sec) of Cahae|

<br><br>

- go_api

| Folders | Usage |
|:--|:--|
|FILE_PORT|FTP Port|

<br><br>

- swagger_ui

| Folders | Usage |
|:--|:--|
|API_URL|location of swagger scripts|

<br><br>

# TODO

- Building Jenkins，provide the automation required for development
- Building gRPC Server，for strong connect service
