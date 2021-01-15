package app

import (
	"errors"
	"golang.org/x/crypto/bcrypt"

	"gopkg.in/mgo.v2/bson"
	"time"
)

type User struct {
	ID            bson.ObjectId `bson:"_id,omitempty" json:"id"`
	CreatedOn     time.Time     `json:"createdon,omitempty"`
	Username      string        `json:"username"`
	Password      string        `json:"password"`
	Email         string        `json:"email"`
	Bio           string        `json:"bio"`
	EmailVerified bool          `json:"email_verified"`
	IsActive      bool          `json:"is_active"`
	//Image         *string       `json:"image"`
}

func (u *User) SetPassword(password string) error {
	if len(password) == 0 {
		return errors.New("password should not be empty")
	}
	bytePassword := []byte(password)
	passwordHash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	u.Password = string(passwordHash)
	return nil
}

func (u *User) CheckPassword(password string) error {
	bytePassword := []byte(password)
	byteHashedPassword := []byte(u.Password)
	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}
