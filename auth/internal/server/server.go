package server

import (
	"context"
	"errors"

	"auth/internal/database"
	"auth/internal/lib"
	pb "auth/proto"

	"github.com/google/uuid"
	"github.com/gookit/slog"
)

type Server struct {
	pb.AuthenticationServer
}

func (s *Server) Signup(ctx context.Context, in *pb.SignupRequest) (*pb.SignupResponse, error) {
	user, err := database.ValidateSignup(in)
	if err != nil {
		slog.Error(err)
		return nil, err
	}

	err = user.Create()
	if err != nil {
		slog.Error(err)
		return nil, err
	}

	tokens, err := user.GenTokens()
	if err != nil {
		slog.Error(err)
		return nil, err
	}

	res := &pb.SignupResponse{
		Id:       user.ID.String(),
		Username: user.Username,
		Email:    user.Email,
		Tokens:   tokens,
	}

	return res, nil
}

func (s *Server) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	slog.Info(in.Username)
	user, err := database.ValidateLogin(in)
	if err != nil {
		slog.Error(err)
		return nil, err
	}

	dbUser := database.Accounts{
		Username: user.Username,
	}

	err = dbUser.GetByUsername()
	if err != nil {
		slog.Error(err)
		return nil, err
	}

	err = dbUser.CheckPassword(user.Password)
	if err != nil {
		return nil, err
	}

	tokens, err := dbUser.GenTokens()
	if err != nil {
		slog.Error(err)
		return nil, err
	}

	return &pb.LoginResponse{
		Id:       dbUser.ID.String(),
		Username: dbUser.Username,
		Email:    dbUser.Email,
		Tokens:   tokens,
	}, nil
}

func (s *Server) Logout(ctx context.Context, in *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	return nil, errors.New("not implemented")
}

func (s *Server) Refresh(ctx context.Context, in *pb.RefreshRequest) (*pb.RefreshResponse, error) {
	err := lib.ValidateToken(in.Refresh)
	if err != nil {
		return nil, err
	}
	payload, err := lib.ParseRefreshToken(in.Refresh)
	if err != nil {
		return nil, err
	}

	accessPayload := lib.AccessTokenPayload{
		Username: payload.Username,
		ID:       payload.ID,
	}
	access, err := accessPayload.GenAccessToken()
	if err != nil {
		return nil, err
	}

	return &pb.RefreshResponse{
		Access: access,
	}, nil
}

func (s *Server) AccountInfo(ctx context.Context, in *pb.AccountInfoRequest) (*pb.AccountInfoResponse, error) {
	u, err := uuid.Parse(in.Id)
	if err != nil {
		return nil, err
	}

	user := database.Accounts{ID: u}

	err = user.GetByID()
	if err != nil {
		return nil, err
	}

	return &pb.AccountInfoResponse{
		Id:       user.ID.String(),
		Username: user.Username,
		Email:    user.Email,
	}, nil
}

func (s *Server) VerifyToken(ctx context.Context, in *pb.VerifyTokenRequest) (*pb.VerifyTokenResponse, error) {
	err := lib.ValidateToken(in.Token)
	return &pb.VerifyTokenResponse{Valid: err == nil}, err
}

func (s *Server) ParseToken(ctx context.Context, in *pb.ParseTokenRequest) (*pb.ParseTokenResponse, error) {
	p, err := lib.ParseAccessToken(in.Token)
	if err != nil {
		return nil, err
	}
	data, err := lib.InterfaceToStringMap(p.Data)
	if err != nil {
		return nil, err
	}

	res := &pb.ParseTokenResponse{
		Id:       p.ID,
		Username: p.Username,
		Data:     data,
	}

	return res, nil
}
