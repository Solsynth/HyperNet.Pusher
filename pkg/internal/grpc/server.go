package grpc

import (
	"git.solsynth.dev/hypernet/pusher/pkg/proto"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	health "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"net"
)

type Server struct {
	proto.UnimplementedPusherServiceServer
	health.UnimplementedHealthServer

	srv *grpc.Server
}

func NewServer() *Server {
	server := &Server{
		srv: grpc.NewServer(),
	}

	proto.RegisterPusherServiceServer(server.srv, server)
	health.RegisterHealthServer(server.srv, server)

	reflection.Register(server.srv)

	return server
}

func (v *Server) Listen() error {
	listener, err := net.Listen("tcp", viper.GetString("grpc_bind"))
	if err != nil {
		return err
	}

	return v.srv.Serve(listener)
}
