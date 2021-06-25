package main

import (
	"context"
	"flag"
	"time"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"

	pgameengine "github.com/choym0098/Reaction-Time-Trainer/m-apis/m-game-engine/v1"
)

func main() {
	var addressPtr = flag.String("address", "localhost:60051", "address to connect")
	flag.Parse()

	// Use grpc.WithInsecure for testing purposes
	conn, err := grpc.Dial(*addressPtr, grpc.WithInsecure())
	if err != nil {
		log.Fatal().Err(err).Str("address", *addressPtr).Msg("failed to dial m-game-engine gRPC service")
	}

	defer func() {
		err := conn.Close()
		if err != nil {
			log.Fatal().Err(err).Str("address", *addressPtr).Msg("failed to close connection")
		}
	}()

	timeoutCtx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	c := pgameengine.NewGameClient(conn)
	if c == nil {
		log.Info().Msg("Client nil")
	}

	r, err := c.GetSize(timeoutCtx, &pgameengine.GetSizeRequest{})
	if err != nil {
		log.Fatal().Err(err).Str("address", *addressPtr).Msg("failed to get a response from GetHighScore call")
	}

	if r != nil {
		log.Info().Interface("size", r.GetSize()).Msg("Size from m-game-engine microservice")
	} else {
		log.Error().Msg("Couldn't get size")
	}

}
