# Reaction-Time-Trainer
Fullstack project written in Golang and html + JavaScript that utilizes gRPC, microservices and reverse proxy. 

Project Architecture Diagram:
![alt text](https://github.com/choym0098/Reaction-Time-Trainer/Reaction-Time-Trainer-Architecture-Diagram.png?raw=true)


## Prerequisites
- Golang
- Protobuf and gRPC
```
brew install protobuf
go get -u google.golang.org/protobuf
go get -u google.golang.org/grpc
```

## Dependency Setup
On the root directory, run
```
go get ./...
```

## Running locally
1. Git clone the repository
```
git clone https://github.com/choym0098/Reaction-Time-Trainer.git
```
2. Run high score microservice first by
```
go run m-highscore/cli/server/main.go
```
3. Run game engine microservice by
```
go run m-highscore/cli/server/main.go
```
4. Run reverse proxy by
```
go run m-reverse-proxy/cli/main.go
```
5. Open index.html in a chrome browser (either double click the file in finder or copy and paste the path to index.html into a chrome browser)

