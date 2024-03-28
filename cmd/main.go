package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"messageBox/internal/server"
	pb "messageBox/pkg/connector"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterMessageBoxServer(s, server.NewServer())

	fmt.Println("listening grpc connections")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
