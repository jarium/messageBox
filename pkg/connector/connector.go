package connector

import (
	"google.golang.org/grpc"
	"time"
)

type Connector struct {
	Client     MessageBoxClient
	connection *grpc.ClientConn
}

func New(addr string) (*Connector, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(time.Second*5))

	if err != nil {
		return nil, err
	}

	return &Connector{
		Client:     NewMessageBoxClient(conn),
		connection: conn,
	}, nil
}

func (c *Connector) CloseConnection() error {
	if c.connection == nil {
		return nil
	}

	return c.connection.Close()
}
