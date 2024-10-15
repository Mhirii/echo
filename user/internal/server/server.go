package server

import (
	"context"
	"errors"

	"user/internal/database"
	pb "user/proto"

	"github.com/google/uuid"
	"github.com/gookit/slog"
)

type Server struct {
	pb.UserServer
	AuthAddr string
}

func (s *Server) Create(ctx context.Context, in *pb.CreateRequest) (*pb.CreateResponse, error) {
	authClient := &AuthClient{addr: &s.AuthAddr}
	tokenData, err := authClient.ParseToken(ctx, in.Token)
	if err != nil {
		slog.Error(err)
		return nil, err
	}

	user := convSignupRequest(in, tokenData.Id)

	err = user.Create()
	if err != nil {
		return nil, err
	}

	res := convSignupResponse(&user)

	return res, nil
}

func (s *Server) Update(ctx context.Context, in *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	authClient := &AuthClient{addr: &s.AuthAddr}
	tokenData, err := authClient.ParseToken(ctx, in.Token)
	if err != nil {
		slog.Error(err)
		return nil, err
	}

	userUUID, err := uuid.Parse(in.UserId)
	if err != nil {
		slog.Error(err)
		return nil, err
	}

	user := convUpdateRequest(in, tokenData.Id, userUUID)

	err = user.UpdatePartial()
	if err != nil {
		return nil, err
	}

	res := convUpdateResponse(&user)

	return res, nil
}

func (s *Server) Read(ctx context.Context, in *pb.ReadRequest) (*pb.ReadResponse, error) {
	slog.Info("Not implemented")
	return nil, errors.New("Read not implemented")
}

func (s *Server) Self(ctx context.Context, in *pb.SelfRequest) (*pb.SelfResponse, error) {
	authClient := &AuthClient{addr: &s.AuthAddr}
	tokenData, err := authClient.ParseToken(ctx, in.Token)
	if err != nil {
		slog.Error(err)
		return nil, err
	}

	user, err := database.FindByAccountId(tokenData.Id)
	if err != nil {
		return nil, err
	}

	res := convSelfResponse(user)

	return res, nil
}
