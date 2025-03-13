package client

import (
	"github.com/charmingruby/impr/lib/proto/gen/pb"
	"google.golang.org/grpc"
)

type Client struct {
	conn *grpc.ClientConn

	Service *Service
}

type Service struct {
	pb.IdentityServiceClient
}

func New(conn *grpc.ClientConn) *Client {
	return &Client{
		conn: conn,
	}
}

func (c *Client) Register() {
	c.Service = &Service{
		IdentityServiceClient: pb.NewIdentityServiceClient(c.conn),
	}
}
