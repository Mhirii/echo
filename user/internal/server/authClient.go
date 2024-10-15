package server

import (
	"context"

	pb "user/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AuthClient struct {
	addr *string
}

func (ac *AuthClient) ParseToken(ctx context.Context, token string) (*pb.ParseTokenResponse, error) {
	conn, err := grpc.NewClient(*ac.addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	c := pb.NewAuthenticationClient(conn)
	in := &pb.ParseTokenRequest{
		Token: token,
	}
	res, err := c.ParseToken(ctx, in)
	return res, err
}
