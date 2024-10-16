package server

import (
	"user/internal/database"
	pb "user/proto"

	"github.com/google/uuid"
)

func convSignupRequest(sr *pb.CreateRequest, accountId string) database.User {
	user := database.User{
		AccountID: accountId,
		Username:  sr.Username,
		FirstName: sr.FirstName,
		LastName:  sr.LastName,
	}

	if sr.Email != nil && *sr.Email != "" {
		user.Email = sr.Email
	} else {
		user.Email = nil
	}

	if sr.Phone != nil && *sr.Phone != "" {
		user.Phone = sr.Phone
	} else {
		user.Phone = nil
	}

	return user
}

func convSignupResponse(user *database.User) *pb.CreateResponse {
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

	return res
}

func convUpdateRequest(in *pb.UpdateRequest, accountId string, userUUID uuid.UUID) database.User {
	user := database.User{
		AccountID: accountId,
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
	} else {
		user.Email = nil
	}
	if in.Phone != nil && *in.Phone != "" {
		user.Phone = in.Phone
	} else {
		user.Phone = nil
	}
	return user
}

func convUpdateResponse(user *database.User) *pb.UpdateResponse {
	res := &pb.UpdateResponse{
		AccountId: user.AccountID,
		UserId:    user.ID.String(),
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
	return res
}

func convSelfResponse(user database.User) *pb.SelfResponse {
	return &pb.SelfResponse{
		AccountId: user.AccountID,
		UserId:    user.ID.String(),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone,
	}
}

func convInfoByIdResponse(user database.User) *pb.InfoByIdResponse {
	return &pb.InfoByIdResponse{
		UserId:    user.ID.String(),
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone,
	}
}

func convInfoByUsernameResponse(user database.User) *pb.InfoByUsernameResponse {
	return &pb.InfoByUsernameResponse{
		UserId:    user.ID.String(),
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone,
	}
}
