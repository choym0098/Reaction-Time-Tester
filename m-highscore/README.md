# High Score Microservice
This microservice provides the current highest score during the game.

## Compile proto
```
cd ./v1
protoc --go_out=plugins=grpc:. *.proto
```

## Running Server/Client
```
// run highscore microservice
go run cli/server/main.go

// run test client for highscore microservice
go run cli/client/main.go
```

