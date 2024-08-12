package main

import (
	"context"
	"math/rand"
	"time"

	"github.com/tlsh0/protoapi"
	"golang.org/x/exp/rand"
)

type RandomServer struct {
	protoapi.UnimplementedRandomServer
}

func (RandomServer) GetDate(ctx context.Context, r *protoapi.RequestDateTime) (*protoapi.DateTime, error) {
	currentTime := time.Now()
	response := &protoapi.DateTime{
		Value: currentTime.String(),
	}
	return response, nil
}

func (RandomServer) GetRandom(ctx context.Context, r *protoapi.RandomParams) (*protoapi.RandomInt, error) {
	rand.Seed(r.GetSeed())
	place := r.GetPlace()
	temp := random(min, max)

	for {
		place--
		if place <= 0 {
			break
		}
		temp = random(min, max)
	}

	response := &protoapi.RandomInt{
		Value: int64(temp),
	}

	return response, nil
}

func (RandomServer) GetRandomPass(ctx context.Context, r *protoapi.RequestPass) (*protoapi.RandomPass, error) {
	rand.Seed(r.GetSeed())
	temp := getString(r.GetLength())

	response := &protoapi.RandomPass{
		Password: temp,
	}

	return response, nil
}
