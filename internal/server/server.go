package server

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jarium/messageBox/internal/box"
	"github.com/jarium/messageBox/pkg/connector"
	"io"
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
		m, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&connector.SendMessageResponse{Success: true})
		}

		if err != nil {
			return err
		}

		if m.Uuid == "" {
			m.Uuid = uuid.New().String()
		} else { //existing message, sent from consumer
			//@todo: save unprocessable message to a data source
			fmt.Printf("message sent back again by consumer, unprocessable message:%s id:%s\n", m.Message, m.Uuid)
			continue
		}

		fmt.Printf("info: msg received with id:%s msg:%s\n", m.Uuid, m.Message)

		s.box.Push(m)
	}
}

func (s *Server) ReceiveMessages(_ *connector.Void, stream connector.MessageBox_ReceiveMessagesServer) error {
	for {
		m := s.box.Pop()

		if err := stream.Send(m); err != nil {
			fmt.Printf("send msg to consumer err:%s", err)

			s.box.Push(m)
			return err
		}

		fmt.Printf("info: msg sent to consumer with id:%s msg:%s\n", m.Uuid, m.Message)
	}
}
