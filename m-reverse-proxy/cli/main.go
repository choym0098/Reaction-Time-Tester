package main

import (
	"flag"
)

func main() {
	grpcAddressHighScore := flag.String("address-m-highscore", "localhost:50051", "The grpc server address for highscore service")
	grpcAddressGameEngine := flag.String("address-m-game-engine", "localhost:60051", "The grpc server address for game-engine service")

	serverAddress := flag.String("address=http", ":8081", "HTTP server address")

	flag.Parse()

}
