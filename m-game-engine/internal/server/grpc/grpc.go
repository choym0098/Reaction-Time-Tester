package grpc

import (
	"context"
	"net"

	pbgameengine "github.com/choym0098/Reaction-Time-Trainer/m-apis/m-game-engine/v1"
	"github.com/choym0098/Reaction-Time-Trainer/m-game-engine/internal/server/logic"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

type Grpc struct {
	address string
	server  *grpc.Server
}

func NewServer(address string) *Grpc {
	return &Grpc{
		address: address,
	}
}

func (g *Grpc) GetSize(ctx context.Context, input *pbgameengine.GetSizeRequest) (*pbgameengine.GetSizeResponse, error) {
	log.Info().Msg("GetSize in m-game-engine called")
	return &pbgameengine.GetSizeResponse{
		Size: logic.GetSize(),
	}, nil
}

func (g *Grpc) SetScore(ctx context.Context, input *pbgameengine.SetScoreRequest) (*pbgameengine.SetScoreResponse, error) {
	log.Info().Msg("SetScore in m-game-engine called")
	set := logic.SetScore(input.Score)
	return &pbgameengine.SetScoreResponse{
		Set: set,
	}, nil
}

func (g *Grpc) ListenAndServe() error {
	listener, err := net.Listen("tcp", g.address)
	if err != nil {
		return errors.Wrap(err, "failed to open tcp port")
	}

	serverOptions := []grpc.ServerOption{}

	g.server = grpc.NewServer(serverOptions...)

	pbgameengine.RegisterGameServer(g.server, g)

	log.Info().Str("address", g.address).Msg("starting gRPC server for highscore microservice")

	err = g.server.Serve(listener)
	if err != nil {
		return errors.Wrap(err, "failed to start gRPC server for highscore microservice")
	}

	return nil
}
