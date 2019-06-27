package main

import (
	"log"
	"net"

	"github.com/Mrcampbell/pgo2/breed-move-service/config"
	bms "github.com/Mrcampbell/pgo2/breed-move-service/grpc"
	"github.com/Mrcampbell/pgo2/breed-move-service/psql"
	"github.com/Mrcampbell/pgo2/protorepo/pokemon"
	"google.golang.org/grpc"
)

func main() {
	// Set up a connection to the server.
	lis, err := net.Listen("tcp", config.GRPCPort)
	if err != nil {
		log.Fatalf("Failed to listen on: %v", config.GRPCPort)
	}

	bmservice, err := psql.NewBreedMoveService()
	if err != nil {
		log.Fatal(err)
	}
	server := grpc.NewServer()

	bmserver := bms.NewBreedMoveServer(*bmservice)

	pokemon.RegisterBreedMoveServiceServer(server, bmserver)

	log.Println("Starting Breed Move Service..")
	server.Serve(lis)
}
