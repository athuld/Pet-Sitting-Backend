package users

import (
	"pet-sitting-backend/utils/errors"
	"strings"
)

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserDetails struct {
	ID          int64  `json:"id"`
	UserID      int64  `json:"user_id"`
	Name        string `json:"name"`
	Gender      string `json:"gender"`
	Age         int    `json:"age"`
	Phone       int64  `json:"phone"`
	Address     string `json:"address"`
	Pincode     int    `json:"pincode"`
	IsPetsitter bool   `json:"is_petsitter"`
	AvatarIMG   string `json:"avatar_img"`
}

type FullUserDetails struct {
	UserID      int64  `json:"user_id"`
	Username string `json:"username"`
    Name        string `json:"name"`
	Email    string `json:"email"`
	Gender      string `json:"gender"`
	Age         int    `json:"age"`
    Address     string `json:"address"`
    Pincode     int    `json:"pincode"`
	Phone       int64  `json:"phone"`
	IsPetsitter bool   `json:"is_petsitter"`
}

func (user *User) ValidateUser() *errors.RestErr {
	user.Username = strings.TrimSpace(user.Username)
	user.Email = strings.TrimSpace(user.Email)
	user.Password = strings.TrimSpace(user.Password)

	result := &User{Email: user.Email}

	_ = result.GetByEmail()

	if result.Username != "" {
		return errors.NewBadRequestError("User already exist")
	}

	if user.Password == "" {
		return errors.NewBadRequestError("Invalid password")
	}

	if user.Username == "" {
		return errors.NewBadRequestError("Invalid username")
	}
	return nil
}
