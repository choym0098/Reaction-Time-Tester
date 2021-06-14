# High Score Microservice
This microservice provides the current highest score during the game.

## Compile proto
```
cd ./v1
protoc --go_out=plugins=grpc:. *.proto
```
