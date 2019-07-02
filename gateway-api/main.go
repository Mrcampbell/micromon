package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/Mrcampbell/pgo2/protorepo/pokemon"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

var (
	pokemonServerEndpoint = "pokemon-service:9090"
	breedServerEndpoint   = "breed-service:9090"
	moveServiceEndpoint   = "move-service:9090"
	battleServiceEndpoint = "battle-service:9090"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
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
	err = pokemon.RegisterBreedServiceHandlerFromEndpoint(ctx, mux, breedServerEndpoint, opts)
	if err != nil {
		return err
	}
	err = pokemon.RegisterMoveServiceHandlerFromEndpoint(ctx, mux, moveServiceEndpoint, opts)
	if err != nil {
		return err
	}
	err = pokemon.RegisterBattleServiceHandlerFromEndpoint(ctx, mux, battleServiceEndpoint, opts)
	if err != nil {
		return err
	}

	fmt.Println("Registering and Listening to Pokemon Service...")
	return http.ListenAndServe(":8081", middlewares(mux))
}

func middlewares(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.String() != "/favicon.ico" {
			fmt.Println(r.URL, r.Method)
		}
		h.ServeHTTP(w, r)
	})
}
