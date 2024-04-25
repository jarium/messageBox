package main

import (
	"context"
	"fmt"
	"github.com/jarium/messageBox/pkg/connector"
	"google.golang.org/grpc/health/grpc_health_v1"
	"os"
)

func main() {
	conn, err := connector.New("localhost:50051")

	if err != nil {
		fmt.Println("Failed to connect:", err)
		os.Exit(1)
	}

	defer func(conn *connector.Connector) {
		err := conn.CloseConnection()
		if err != nil {
			fmt.Println("Failed to close connection:", err)
		}
	}(conn)

	client := grpc_health_v1.NewHealthClient(conn.Connection)
	request := &grpc_health_v1.HealthCheckRequest{}
	response, err := client.Check(context.Background(), request)

	if err != nil || response.Status != grpc_health_v1.HealthCheckResponse_SERVING {
		fmt.Println("Service not healthy:", err)
		os.Exit(1)
	}

	os.Exit(0)
}
