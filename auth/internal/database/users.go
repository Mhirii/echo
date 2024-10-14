package database

import (
	"auth/internal/lib"
	pb "auth/proto"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Username string    `gorm:"unique;not null"`
	Email    *string   `gorm:"unique"`
	Password string    `gorm:"not null"`
}

type tokens struct {
	access  string
	refresh string
}

func (u *Users) Create() error {
	hashed, err := lib.Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = hashed
	res := dbInstance.db.Save(&u)
	if res.Error != nil {
		return res.Error
	}
	return res.Error
}

func (u *Users) GetByUsername() error {
	res := dbInstance.db.
		Where("username = ?", u.Username).
		First(&u)

	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (u *Users) GetByID() error {
	res := dbInstance.db.Find(&u, u.ID)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func ValidateSignup(in *pb.SignupRequest) (Users, error) {
	user := Users{}
	usernameErr := lib.ValidateUsername(in.Username)
	if usernameErr != nil {
		return Users{}, usernameErr
	}
	user.Username = in.Username
	passErr := lib.ValidatePassword(in.Password)
	if passErr != nil {
		return Users{}, passErr
	}
	user.Password = in.Password

	if *in.Email != "" {
		err := lib.ValidateEmail(*in.Email)
		if err != nil {
			return Users{}, err
		}
		user.Email = in.Email
	}

	return user, nil
}

func ValidateLogin(in *pb.LoginRequest) (Users, error) {
	user := Users{}
	usernameErr := lib.ValidateUsername(in.Username)
	if usernameErr != nil {
		return Users{}, usernameErr
	}
	user.Username = in.Username
	passErr := lib.ValidatePassword(in.Password)
	if passErr != nil {
		return Users{}, passErr
	}
	user.Password = in.Password

	return user, nil
}

func (u *Users) GenTokens() (*pb.Tokens, error) {
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

func (u *Users) CheckPassword(password string) error {
	return lib.ComparePassword(u.Password, password)
}
