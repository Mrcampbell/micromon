package main

import (
	"log"
	"net"

	"github.com/Mrcampbell/pgo2/battle-service/config"
	// "github.com/Mrcampbell/pgo2/battle-service/psql"
	"github.com/Mrcampbell/pgo2/protorepo/pokemon"
	"google.golang.org/grpc"
)

func main() {
	// Set up a connection to the server.
	lis, err := net.Listen("tcp", config.GRPCPort)
	if err != nil {
		log.Fatalf("Failed to listen on: %v", config.GRPCPort)
	}

	pokemonConn, err := grpc.Dial(config.PokemonServiceAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	pokemonClient := pokemon.NewPokemonServiceClient(pokemonConn)

	moveConn, err := grpc.Dial(config.MoveServiceAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	moveClient := pokemon.NewMoveServiceClient(moveConn)

	service, err := psql.NewBattleService(pokemonClient, moveClient)
	if err != nil {
		log.Fatal(err)
	}
	server := grpc.NewServer()
	pokemon.RegisterBattleServiceServer(server, service)

	log.Println("Starting Pokemon Service..")
	server.Serve(lis)
}
