package main

import (
	"context"
	"flag"
	"time"

	"github.com/prakharmaurya/m-highscore/api"
	"github.com/rs/zerolog/log"

	"google.golang.org/grpc"
)

func main() {
	var addressPtr = flag.String("address", "localhost:50051", "address to connect")
	flag.Parse()

	con, err := grpc.Dial(*addressPtr, grpc.WithInsecure())
	if err != nil {
		log.Fatal().Err(err).Str("address", *addressPtr).Msg("failed to dial m-highscore gRPC service")
	}

	defer func() {
		err := con.Close()
		if err != nil {
			log.Fatal().Err(err).Str("address", *addressPtr).Msg("Failed to close the connecton")
		}
	}()

	c := api.NewGameClient(con)

	if c == nil {
		log.Info().Msg("Client nil")
	}

	timeoutCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	r, err := c.GetHighScore(timeoutCtx, &api.GetHighScoreRequest{})
	if err != nil {
		log.Fatal().Err(err).Str("address", *addressPtr).Msg("faild to get a response")
	}
	if r != nil {
		log.Info().Interface("interface", r.GetHighScore()).Msg("Highscore from m-hoghscore service")
	} else {
		log.Error().Msg("Couldn't get highscore")
	}
}
