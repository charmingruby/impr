package server

import (
	"github.com/charmingruby/impr/lib/proto/gen/pb"
	"github.com/charmingruby/impr/service/identity/internal/account/transport/shared/client"
	"google.golang.org/grpc"
)

type Server struct {
	gRPCServer *grpc.Server
	service    *Service
}

type Service struct {
	pb.UnimplementedIdentityServiceServer

	identityProviderClient *client.CognitoIdentityProvider
}

func New(
	gRPCServer *grpc.Server,
	identityProviderClient *client.CognitoIdentityProvider,
) *Server {
	return &Server{
		gRPCServer: gRPCServer,
		service: &Service{
			identityProviderClient: identityProviderClient,
		},
	}
}

func (s *Server) Register() {
	pb.RegisterIdentityServiceServer(s.gRPCServer, s.service)
}
