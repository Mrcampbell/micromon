package main

import (
	"log"
	"net"

	"github.com/Mrcampbell/pgo2/move-service/config"
	"github.com/Mrcampbell/pgo2/move-service/moveservice"
	"github.com/Mrcampbell/pgo2/protorepo/pokemon"
	"google.golang.org/grpc"
)

func main() {
	// Set up a connection to the server.
	lis, err := net.Listen("tcp", config.GRPCPort)
	if err != nil {
		log.Fatalf("Failed to listen on: %v", config.GRPCPort)
	}

	service := moveservice.NewMoveService()
	server := grpc.NewServer()
	pokemon.RegisterMoveServiceServer(server, service)

	log.Println("Starting Move Service..")
	server.Serve(lis)
}
