package main

import (
	"log"
	"net"

	"github.com/Mrcampbell/pgo2/breed-service/boltdb"
	"github.com/Mrcampbell/pgo2/breed-service/config"
	"github.com/Mrcampbell/pgo2/protorepo/pokemon"
	"google.golang.org/grpc"
)

func main() {
	// Set up a connection to the server.
	lis, err := net.Listen("tcp", config.GRPCPort)
	if err != nil {
		log.Fatalf("Failed to listen on: %v", config.GRPCPort)
	}

	conn, err := grpc.Dial(config.BreedMoveService, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	breedMoveClient := pokemon.NewBreedMoveServiceClient(conn)

	service, err := boltdb.NewBreedService(breedMoveClient)

	if err != nil {
		log.Fatal(err)
	}
	server := grpc.NewServer()
	pokemon.RegisterBreedServiceServer(server, service)

	log.Println("Starting Breed Service..")
	server.Serve(lis)
}
