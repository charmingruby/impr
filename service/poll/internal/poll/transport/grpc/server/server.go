package server

import (
	"github.com/charmingruby/impr/lib/proto/gen/pb"
	"github.com/charmingruby/impr/service/poll/internal/poll/core/service"
	"google.golang.org/grpc"
)

type Server struct {
	gRPCServer *grpc.Server
	service    *Service
}

type Service struct {
	pb.UnimplementedPollServiceServer

	domainService *service.Service
}

func New(
	gRPCServer *grpc.Server,
	service *service.Service,
) *Server {
	return &Server{
		gRPCServer: gRPCServer,
		service: &Service{
			domainService: service,
		},
	}
}

func (s *Server) Register() {
	pb.RegisterPollServiceServer(s.gRPCServer, s.service)
}
