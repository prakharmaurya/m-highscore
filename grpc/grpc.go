package grpc

import (
	"context"
	"net"

	"github.com/pkg/errors"
	"github.com/prakharmaurya/m-highscore/api"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

type Grpc struct {
	address string
	srv     *grpc.Server
}

var HighScore = 99999999999.0

func NewServer(address string) *Grpc {
	return &Grpc{
		address: address,
	}
}

func (g *Grpc) SetHighScore(ctx context.Context, input *api.SetHighScoreRequest) (*api.SetHighScoreResponse, error) {
	log.Info().Msg("SetHighScore in m-highscore called")
	HighScore = input.HighScore
	return &api.SetHighScoreResponse{
		Set: true,
	}, nil
}

func (g *Grpc) GetHighScore(ctx context.Context, input *api.GetHighScoreRequest) (*api.GetHighScoreResponse, error) {
	log.Info().Msg("GetHighScore in m-highscore called")
	return &api.GetHighScoreResponse{
		HighScore: HighScore,
	}, nil
}

func (g *Grpc) ListenAndServe() error {
	lis, err := net.Listen("tcp", g.address)
	if err != nil {
		return errors.Wrap(err, "failed to open tcp port")
	}
	serverOpts := []grpc.ServerOption{}
	g.srv = grpc.NewServer(serverOpts...)

	api.RegisterGameServer(g.srv, g)
	log.Info().Str("address", g.address).Msg("starting grpc server m-highscore microservice")

	err = g.srv.Serve(lis)
	if err != nil {
		return errors.Wrap(err, "failed to start grpc server m-highscore microservice")
	}
	return nil
}
