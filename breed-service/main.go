package main

import (
	"log"
	"net"

	"github.com/Mrcampbell/pgo2/breed-service/mock"
	"github.com/Mrcampbell/pgo2/pokemon-service/config"
	"github.com/Mrcampbell/pgo2/protorepo/pokemon"
	"google.golang.org/grpc"
)

func main() {
	// Set up a connection to the server.
	lis, err := net.Listen("tcp", config.GRPCPort)
	if err != nil {
		log.Fatalf("Failed to listen on: %v", config.GRPCPort)
	}

	service := mock.NewBreedService()
	server := grpc.NewServer()
	pokemon.RegisterBreedServiceServer(server, service)

	log.Println("Starting Breed Service..")
	server.Serve(lis)
}
