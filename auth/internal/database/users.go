package database

import (
	"auth/internal/lib"
	pb "auth/proto"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Accounts struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Username string    `gorm:"unique;not null"`
	Email    *string   `gorm:"unique"`
	Password string    `gorm:"not null"`
}

func (a *Accounts) Create() error {
	hashed, err := lib.Hash(a.Password)
	if err != nil {
		return err
	}
	a.Password = hashed
	res := dbInstance.db.Save(&a)
	if res.Error != nil {
		return res.Error
	}
	return res.Error
}

func (a *Accounts) GetByUsername() error {
	res := dbInstance.db.
		Where("username = ?", a.Username).
		First(&a)

	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (a *Accounts) GetByID() error {
	res := dbInstance.db.Find(&a, a.ID)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func ValidateSignup(in *pb.SignupRequest) (Accounts, error) {
	user := Accounts{}
	usernameErr := lib.ValidateUsername(in.Username)
	if usernameErr != nil {
		return Accounts{}, usernameErr
	}
	user.Username = in.Username
	passErr := lib.ValidatePassword(in.Password)
	if passErr != nil {
		return Accounts{}, passErr
	}
	user.Password = in.Password

	if *in.Email != "" {
		err := lib.ValidateEmail(*in.Email)
		if err != nil {
			return Accounts{}, err
		}
		user.Email = in.Email
	}

	return user, nil
}

func ValidateLogin(in *pb.LoginRequest) (Accounts, error) {
	user := Accounts{}
	usernameErr := lib.ValidateUsername(in.Username)
	if usernameErr != nil {
		return Accounts{}, usernameErr
	}
	user.Username = in.Username
	passErr := lib.ValidatePassword(in.Password)
	if passErr != nil {
		return Accounts{}, passErr
	}
	user.Password = in.Password

	return user, nil
}

func (u *Accounts) GenTokens() (*pb.Tokens, error) {
	accessPl := lib.AccessTokenPayload{
		ID:       u.ID.String(),
		Username: u.Username,
	}
	access, err := accessPl.GenAccessToken()
	if err != nil {
		return &pb.Tokens{}, err
	}

	refreshPl := lib.RefreshTokenPayload{
		ID:       u.ID.String(),
		Username: u.Username,
	}
	refresh, err := refreshPl.GenRefershToken()
	if err != nil {
		return &pb.Tokens{}, err
	}

	return &pb.Tokens{
		Access:  access,
		Refresh: refresh,
	}, nil
}

func (u *Accounts) CheckPassword(password string) error {
	return lib.ComparePassword(u.Password, password)
}
