package server

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/gookit/slog"
	pb "user/proto"
)

type Server struct {
	pb.userServer
} 

func (s *Server) Create(ctx context.Context, in *pb.CreateRequest) (*pb.CreateResponse, error) {
	return nil, errors.New("Create not implemented")
}

func (s *Server) Update(ctx context.Context, in *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	return nil, errors.New("Update not implemented")
}

func (s *Server) Read(ctx context.Context, in *pb.ReadRequest) (*pb.ReadResponse, error) {
	return nil, errors.New("Read not implemented")
}

func (s *Server) Info(ctx context.Context, in *pb.InfoRequest) (*pb.InfoResponse, error) {
	return nil, errors.New("Info not implemented")
}

func (s *Server) Self(ctx context.Context, in *pb.SelfRequest) (*pb.SelfResponse, error) {
	return nil, errors.New("Self not implemented")
}

