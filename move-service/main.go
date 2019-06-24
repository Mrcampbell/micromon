package main

import (
	"context"
	"log"
	"net"

	"github.com/Mrcampbell/pgo2/move-service/config"
	"github.com/Mrcampbell/pgo2/move-service/dgraph"
	"github.com/Mrcampbell/pgo2/protorepo/pokemon"
	"google.golang.org/grpc"
)

func main() {
	// Set up a connection to the server.
	lis, err := net.Listen("tcp", config.GRPCPort)
	if err != nil {
		log.Fatalf("Failed to listen on: %v", config.GRPCPort)
	}

	ctx := context.Background()

	service, err := dgraph.NewMoveService(ctx)
	if err != nil {
		log.Fatal(err)
	}
	server := grpc.NewServer()
	pokemon.RegisterMoveServiceServer(server, service)

	log.Println("Starting Pokemon Service..")
	server.Serve(lis)
}
