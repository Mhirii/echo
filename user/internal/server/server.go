package server

import (
	"context"
	"errors"

	"user/internal/database"
	pb "user/proto"

	"github.com/gookit/slog"
)

type Server struct {
	pb.UserServer
}

func (s *Server) Create(ctx context.Context, in *pb.CreateRequest) (*pb.CreateResponse, error) {
	slog.Info("Not implemented")
	userInput := database.User{
		AccountID: in.AccountId,
		Username:  in.Username,
		FirstName: in.FirstName,
		LastName:  in.LastName,
	}

	if in.Email != nil && *in.Email != "" {
		userInput.Email = in.Email
	} else {
		userInput.Email = nil
	}

	if in.Phone != nil && *in.Phone != "" {
		userInput.Phone = in.Phone
	} else {
		userInput.Phone = nil
	}

	err := userInput.Create()
	if err != nil {
		return nil, err
	}

	res := &pb.CreateResponse{
		AccountId: in.AccountId,
		Username:  in.Username,
		FirstName: in.FirstName,
		LastName:  in.LastName,
	}

	if userInput.Email != nil {
		res.Email = userInput.Email
	}
	if userInput.Phone != nil {
		res.Phone = userInput.Phone
	}

	return res, nil
}

func (s *Server) Update(ctx context.Context, in *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	slog.Info("Not implemented")
	return nil, errors.New("Update not implemented")
}

func (s *Server) Read(ctx context.Context, in *pb.ReadRequest) (*pb.ReadResponse, error) {
	slog.Info("Not implemented")
	return nil, errors.New("Read not implemented")
}

func (s *Server) Self(ctx context.Context, in *pb.SelfRequest) (*pb.SelfResponse, error) {
	slog.Info("Not implemented")
	return nil, errors.New("Self not implemented")
}
