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
