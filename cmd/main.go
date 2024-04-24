package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/jarium/messageBox/internal/server"
	pb "github.com/jarium/messageBox/pkg/connector"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"net"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterMessageBoxServer(s, server.NewServer())

	hs := health.NewServer()
	grpc_health_v1.RegisterHealthServer(s, hs)
	hs.SetServingStatus("", grpc_health_v1.HealthCheckResponse_SERVING)

	fmt.Println("info: listening grpc connections")

	go func() {
		// create a context that will be canceled when SIGINT or SIGTERM is received
		ctxListen, cancelListen := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
		defer cancelListen()

		// Block until a signal is received
		<-ctxListen.Done()

		fmt.Println("info: shutting down gRPC server")
		s.GracefulStop()

		// Close the listener
		if err := lis.Close(); err != nil {
			log.Fatalf("failed to close listener: %v", err)
		}

		// add timeout to wait for existing connections to finish
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		<-ctx.Done()

		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			fmt.Println("info: shutdown timed out, forcing stop")
			s.Stop()
		} else {
			fmt.Println("info: graceful shutdown completed")
		}
	}()

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
