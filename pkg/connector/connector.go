package connector

import (
	"google.golang.org/grpc"
	"time"
)

type Connector struct {
	Client     MessageBoxClient
	Connection *grpc.ClientConn
}

func New(addr string) (*Connector, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(time.Second*5))

	if err != nil {
		return &Connector{}, err
	}

	return &Connector{
		Client:     NewMessageBoxClient(conn),
		Connection: conn,
	}, nil
}

func (c *Connector) CloseConnection() error {
	if c.Connection == nil {
		return nil
	}

	return c.Connection.Close()
}
