package main

import (
	"flag"

	"github.com/prakharmaurya/m-highscore/grpc"
	"github.com/rs/zerolog/log"
)

func main() {
	var addressPtr = flag.String("address", ":50051", "addresss where you can connect to m-highscore service")
	flag.Parse()

	s := grpc.NewServer(*addressPtr)
	err := s.ListenAndServe()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to start grpc server of m-highscore")
	}
}
