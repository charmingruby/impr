package server

import (
	"context"

	"github.com/charmingruby/impr/lib/pkg/core/id"
	"github.com/charmingruby/impr/lib/proto/gen/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) VerifyToken(ctx context.Context, payload *pb.VerifyTokenPayload) (*pb.VerifyTokenResponse, error) {
	result, err := s.identityProviderClient.RetrieveUserAttributesFromToken(payload.Token)
	if err != nil {
		return &pb.VerifyTokenResponse{
			Id:      id.New(),
			IsValid: false,
		}, status.Errorf(codes.Unauthenticated, "invalid token")
	}

	return &pb.VerifyTokenResponse{
		Id:        id.New(),
		IsValid:   true,
		AccountId: result.ID,
	}, nil
}
