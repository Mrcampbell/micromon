package main

import (
	"log"
	"net"

	"github.com/Mrcampbell/pgo2/battle-service/config"
	pgrpc "github.com/Mrcampbell/pgo2/battle-service/grpc"
	"github.com/Mrcampbell/pgo2/battle-service/psql"
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

	battleService, err := psql.NewBattleService()
	battleServer, err := pgrpc.NewBattleServer(*battleService, pokemonClient, moveClient)
	if err != nil {
		log.Fatal(err)
	}
	server := grpc.NewServer()
	pokemon.RegisterBattleServiceServer(server, battleServer)

	log.Println("Starting Battle Service..")
	server.Serve(lis)
}
