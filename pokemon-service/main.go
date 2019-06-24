package main

import (
	"log"
	"net"

	"github.com/Mrcampbell/pgo2/pokemon-service/config"
	"github.com/Mrcampbell/pgo2/pokemon-service/psql"
	"github.com/Mrcampbell/pgo2/protorepo/pokemon"
	"google.golang.org/grpc"
)

func main() {
	// Set up a connection to the server.
	lis, err := net.Listen("tcp", config.GRPCPort)
	if err != nil {
		log.Fatalf("Failed to listen on: %v", config.GRPCPort)
	}

	service, err := psql.NewPokemonService()
	if err != nil {
		log.Fatal(err)
	}
	server := grpc.NewServer()
	pokemon.RegisterPokemonServiceServer(server, service)

	log.Println("Starting Pokemon Service..")
	server.Serve(lis)
}
