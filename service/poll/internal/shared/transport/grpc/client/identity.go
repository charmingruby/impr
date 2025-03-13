package client

import (
	"context"
	"time"

	"github.com/charmingruby/impr/lib/proto/gen/pb"
)

type VerifyTokenResponse struct {
	IsValid   bool
	AccountID string
}

func (s *Service) VerifyToken(token string) (VerifyTokenResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	req := &pb.VerifyTokenPayload{
		Token: token,
	}

	res, err := s.IdentityServiceClient.VerifyToken(ctx, req)
	if err != nil {
		return VerifyTokenResponse{
			IsValid:   false,
			AccountID: "",
		}, err
	}

	return VerifyTokenResponse{
		IsValid:   res.IsValid,
		AccountID: res.AccountId,
	}, nil
}
