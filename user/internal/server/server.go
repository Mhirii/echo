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

	res := &pb.CreateResponse{
		AccountId: user.AccountID,
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}

	if user.Email != nil {
		res.Email = user.Email
	}
	if user.Phone != nil {
		res.Phone = user.Phone
	}

	return res, nil
}

func (s *Server) Update(ctx context.Context, in *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	userUUID, err := uuid.Parse(in.UserId)
	if err != nil {
		return nil, err
	}
	user := database.User{
		AccountID: in.AccountId,
		ID:        userUUID,
	}

	if in.FirstName != nil {
		user.FirstName = *in.FirstName
	}

	if in.LastName != nil {
		user.LastName = *in.LastName
	}

	if in.Email != nil && *in.Email != "" {
		user.Email = in.Email
	}

	if in.Phone != nil && *in.Phone != "" {
		user.Phone = in.Phone
	}

	err = user.UpdatePartial()
	if err != nil {
		return nil, err
	}

	res := &pb.UpdateResponse{
		AccountId: user.AccountID,
		UserId:    user.ID.String(),
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone,
	}

	return res, nil
}

func (s *Server) Read(ctx context.Context, in *pb.ReadRequest) (*pb.ReadResponse, error) {
	slog.Info("Not implemented")
	return nil, errors.New("Read not implemented")
}

func (s *Server) Self(ctx context.Context, in *pb.SelfRequest) (*pb.SelfResponse, error) {
	slog.Info("Not implemented")
	return nil, errors.New("Self not implemented")
}
