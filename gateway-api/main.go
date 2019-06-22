package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Mrcampbell/pgo2/protorepo/pokemon"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

var (
	pokemonServerEndpoint = "pokemon-service:9090"
)

func main() {
	if err := run(); err != nil {
		glog.Fatal(err)
	}
}

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := pokemon.RegisterPokemonServiceHandlerFromEndpoint(ctx, mux, pokemonServerEndpoint, opts)
	if err != nil {
		return err
	}
	fmt.Println("Registering and Listening to Pokemon Service...")
	return http.ListenAndServe(":8081", mux)
}
