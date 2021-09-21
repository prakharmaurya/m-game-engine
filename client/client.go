package main

import (
	"context"
	"flag"
	"time"

	"github.com/prakharmaurya/m-game-engine/api"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

func main() {
	var addressPtr = flag.String("address", "localhost:50051", "address to connect")
	flag.Parse()

	con, err := grpc.Dial(*addressPtr, grpc.WithInsecure())
	if err != nil {
		log.Fatal().Err(err).Str("address", *addressPtr).Msg("failed to dial m-game-engine gRPC service")
	}

	defer func() {
		err := con.Close()
		if err != nil {
			log.Fatal().Err(err).Str("address", *addressPtr).Msg("Failed to close the connecton")
		}
	}()

	c := api.NewGameEngineClient(con)

	if c == nil {
		log.Info().Msg("Client nil")
	}

	timeoutCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	s, err := c.SetScore(timeoutCtx, &api.SetScoreRequest{Score: 1})
	if err != nil {
		log.Fatal().Err(err).Str("address", *addressPtr).Msg("failed to set score")
	}
	if s != nil {
		log.Info().Interface("interface", s.GetSet()).Msg("Set from Set Score service")
	} else {
		log.Error().Msg("Couldn't set score")
	}

	r, err := c.GetSize(timeoutCtx, &api.GetSizeRequest{})
	if err != nil {
		log.Fatal().Err(err).Str("address", *addressPtr).Msg("faild to get a response")
	}
	if r != nil {
		log.Info().Interface("interface", r.GetSize()).Msg("GetSize from  service")
	} else {
		log.Error().Msg("Couldn't get size")
	}

}
