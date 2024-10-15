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

func convUpdateRequest(ur *pb.UpdateRequest, userUUID uuid.UUID) database.User {
	user := database.User{
		AccountID: ur.AccountId,
		ID:        userUUID,
	}
	if ur.FirstName != nil {
		user.FirstName = *ur.FirstName
	}
	if ur.LastName != nil {
		user.LastName = *ur.LastName
	}
	if ur.Email != nil && *ur.Email != "" {
		user.Email = ur.Email
	} else {
		user.Email = nil
	}
	if ur.Phone != nil && *ur.Phone != "" {
		user.Phone = ur.Phone
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
