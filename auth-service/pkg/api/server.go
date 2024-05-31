package server

import (
	"fmt"
	"net"

	"Auth/pkg/config"
	pb "Auth/pkg/pb/auth"

	"google.golang.org/grpc"
)

type Server struct {
	server   *grpc.Server
	listener net.Listener
}

func NewGRPCServer(cfg config.Config, adminServer pb.AdminServer, employerServer pb.EmployerServer, jobseekerServer pb.JobSeekerServer) (*Server, error) {
	lis, err := net.Listen("tcp", cfg.Port)
	if err != nil {
		return nil, err
	}

	newServer := grpc.NewServer()
	pb.RegisterAdminServer(newServer, adminServer)
	pb.RegisterEmployerServer(newServer, employerServer)
	pb.RegisterJobSeekerServer(newServer, jobseekerServer)

	return &Server{
		server:   newServer,
		listener: lis,
	}, nil
}

func (c *Server) Start() error {
	fmt.Println("grpc server listening on port :50052")
	return c.server.Serve(c.listener)
}
