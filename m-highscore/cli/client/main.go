package main

import (
	"context"
	"flag"
	"time"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"

	pbhighscore "github.com/choym0098/Reaction-Time-Tester/m-apis/m-highscore/v1"
)

func main() {
	var addressPtr = flag.String("address", "localhost:50051", "address to connect")
	flag.Parse()

	// Use grpc.WithInsecure for testing purposes
	conn, err := grpc.Dial(*addressPtr, grpc.WithInsecure())
	if err != nil {
		log.Fatal().Err(err).Str("address", *addressPtr).Msg("failed to dial m-highscore gRPC service")
	}

	defer func() {
		err := conn.Close()
		if err != nil {
			log.Fatal().Err(err).Str("address", *addressPtr).Msg("failed to close connection")
		}
	}()

	timeoutCtx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	c := pbhighscore.NewGameClient(conn)
	if c == nil {
		log.Info().Msg("Client nil")
	}

	r, err := c.GetHighScore(timeoutCtx, &pbhighscore.GetHighScoreRequest{})
	if err != nil {
		log.Fatal().Err(err).Str("address", *addressPtr).Msg("failed to get a response from GetHighScore call")
	}

	if r != nil {
		log.Info().Interface("highscore", r.GetHighScore()).Msg("HighScore from m-highscore microservice")
	} else {
		log.Error().Msg("Couldn't get highscore")
	}

}
