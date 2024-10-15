package server

import (
	"user/internal/database"
	pb "user/proto"
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
