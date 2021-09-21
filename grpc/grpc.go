package grpcServer

import (
	"context"
	"net"

	"github.com/pkg/errors"
	"github.com/prakharmaurya/m-game-engine/api"
	"github.com/prakharmaurya/m-game-engine/logic"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

type Grpc struct {
	address string
	srv     *grpc.Server
}

func NewServer(address string) *Grpc {
	return &Grpc{
		address: address,
	}
}

func (g *Grpc) GetSize(ctx context.Context, input *api.GetSizeRequest) (*api.GetSizeResponse, error) {
	log.Info().Msg("Get size called in game engine")
	return &api.GetSizeResponse{
		Size: logic.GetSize(),
	}, nil
}

func (g *Grpc) SetScore(ctx context.Context, input *api.SetScoreRequest) (*api.SetScoreResponse, error) {
	log.Info().Msg("Set Score called in game engine ")
	set := logic.SetScore(input.Score)
	return &api.SetScoreResponse{
		Set: set,
	}, nil
}

func (g *Grpc) ListenAndServe() error {
	lis, err := net.Listen("tcp", g.address)
	if err != nil {
		return errors.Wrap(err, "failed to open tcp port")
	}
	serverOpts := []grpc.ServerOption{}
	g.srv = grpc.NewServer(serverOpts...)

	api.RegisterGameEngineServer(g.srv, g)
	log.Info().Str("address", g.address).Msg("starting grpc server m-game-engine microservice")

	err = g.srv.Serve(lis)
	if err != nil {
		return errors.Wrap(err, "failed to start grpc server m-game-engine microservice")
	}
	return nil
}
