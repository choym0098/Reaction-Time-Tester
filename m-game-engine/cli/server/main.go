package main

import (
	"flag"

	grpcSetup "github.com/choym0098/Reaction-Time-Trainer/m-game-engine/internal/server/grpc"
	"github.com/rs/zerolog/log"
)

func main() {
	var addressPtr = flag.String("address", ":60051", "address where you can connect with m-highscore service")
	flag.Parse()

	s := grpcSetup.NewServer(*addressPtr)

	err := s.ListenAndServe()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to start grpc server of m-highscore")
	}

}
