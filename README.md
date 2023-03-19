# gophKeeper
### Safe storage of your confidential information.
[![Golang](https://img.shields.io/badge/-Go-0000FF?style=flat-square&logo=Go)](https://go.dev)
[![gRPC](https://img.shields.io/badge/-gRPC-0000FF?style=flat-square&logo=gRPC)](https://grpc.io)
[![protobuf](https://img.shields.io/badge/-protobuf-0000FF?style=flat-square&logo=protobuf)](https://protobuf.dev)
[![PostgreSQL](https://img.shields.io/badge/-PostgreSQL-0000FF?style=flat-square&logo=PostgreSQL)](https://www.postgresql.org/)
[![Goose](https://img.shields.io/badge/-Goose-0000FF?style=flat-square&logo=Goose)](https://github.com/pressly/goose)
[![Docker](https://img.shields.io/badge/-Docker-0000FF?style=flat-square&logo=Docker)](https://docker.com/)
[![Linux](https://img.shields.io/badge/-linux-0000FF?style=flat-square&logo=linux)](https://www.kernel.org/)
[![Openssl](https://img.shields.io/badge/-Openssl-0000FF?style=flat-square&logo=Openssl)](https://www.openssl.org)
[![MIRO](https://img.shields.io/badge/-MIRO-0000FF?style=flat-square&logo=MIRO)](https://miro.com/ru/)

![gophkeeper](https://github.com/ervand7/gophKeeper/blob/master/images/gophkeeper.png)

Launch: 
 1) Open first terminal window and enter `make runserver`
 2) Open second terminal window and enter one of the following commands depending on your operating system:
    - `make runclient-darwin-amd64`
    - `make runclient-darwin-arm64`
    - `make runclient-linux-amd64`
    - `make runclient-windows-amd64`

Read [interaction scenarios.](https://github.com/ervand7/gophKeeper/tree/master/diagrams)

App can store, read, edit and delete next kind of data: credentials, bank card data, arbitrary text, binary data.
After launching client app there creates protected gRPC connection between
client and server via TLS certificates. Also, all sensitive data are passed to server and stored in 
it database as garbled rows which can be decrypted only with special key,
which user sets while launching client app.

User interacts with client app via TUI (terminal user interface), which shows
list of available commands, prompts and autocomplete.
There is an opportunity of using client app simultaneously with many devises of one user. User data 
from each device synchronize each other every one second via open gRPC streaming.
Server sends data from his database to each device. And there data are stored im RAM.

Tests coverage - 80%.