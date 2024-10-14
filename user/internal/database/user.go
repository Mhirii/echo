package database

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	AccountID string    `gorm:"unique;not null"`
	Username  string    `gorm:"unique;not null"`
	FirstName string    `gorm:"not null"`
	LastName  string    `gorm:"not null"`
	Email     *string   `gorm:"unique"`
	Phone     *string   `gorm:"unique"`
}

func (u *User) Create() error {
	res := dbInstance.db.Save(&u)
	return res.Error
}

func (u *User) Update() error {
	res := dbInstance.db.Save(&u)
	return res.Error
}

func (u *User) Delete() error {
	res := dbInstance.db.Delete(&u)
	return res.Error
}

func FindAll() ([]User, error) {
	users := []User{}
	res := dbInstance.db.Find(&users)
	if res.Error != nil {
		return users, res.Error
	}
	return users, nil
}

func FindById(id string) (User, error) {
	user := User{}
	res := dbInstance.db.Where("id = ?", id).First(user)
	if res.Error != nil {
		return user, res.Error
	}
	return user, nil
}

func FindByUsername(username string) (User, error) {
	user := User{}
	res := dbInstance.db.Where("username = ?", username).First(user)
	if res.Error != nil {
		return user, res.Error
	}
	return user, nil
}

func FindByAccountId(accountId string) (User, error) {
	user := User{}
	res := dbInstance.db.Where("account_id = ?", accountId).First(user)
	if res.Error != nil {
		return user, res.Error
	}
	return user, nil
}
