package grpc

import (
	"context"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	*grpc.Server
	lis     net.Listener
	err     error
	baseCtx context.Context
	address string
}

func NewServer(address string, serverOption ...grpc.ServerOption) *Server {
	srv := &Server{
		address: address,
	}

	srv.Server = grpc.NewServer(serverOption...)
	srv.err = srv.Listen()

	return srv
}
func (s *Server) Start(ctx context.Context) error {
	if s.err != nil {
		return s.err
	}
	s.baseCtx = ctx
	return s.Serve(s.lis)
}

func (s *Server) Stop(ctx context.Context) error {
	s.GracefulStop()
	return nil
}

func (s *Server) Listen() error {
	if s.lis == nil {
		lis, err := net.Listen("tcp", s.address)
		if err != nil {
			return err
		}
		s.lis = lis
	}
	return nil
}
