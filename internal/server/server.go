package server

import (
	"github.com/google/uuid"
	"io"
	"messageBox/internal/box"
	"messageBox/pkg/connector"
)

type Server struct {
	connector.UnimplementedMessageBoxServer
	box *box.Box
}

func NewServer() *Server {
	return &Server{
		box: box.New(),
	}
}

func (s *Server) SendMessages(stream connector.MessageBox_SendMessagesServer) error {
	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&connector.SendMessageResponse{Success: true})
		}

		if err != nil {
			return err
		}

		m := box.Message{
			Uuid:    uuid.New().String(),
			Message: msg.Message,
		}

		s.box.Push(m)
	}
}

func (s *Server) ReceiveMessages(_ *connector.Void, stream connector.MessageBox_ReceiveMessagesServer) error {
	for msg := s.box.Pop(); msg.Uuid != ""; msg = s.box.Pop() {
		if err := stream.Send(&connector.Message{Uuid: msg.Uuid, Message: msg.Message}); err != nil {
			return err
		}
	}

	return nil
}
