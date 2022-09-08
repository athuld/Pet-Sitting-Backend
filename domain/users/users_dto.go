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
	Gender      string `json:"gender"`
	Age         int    `json:"age"`
	Address     string `json:"address"`
	Pincode     int    `json:"pincode"`
	IsPetsitter bool   `json:"is_petsitter"`
	IsDogwalker bool   `json:"is_dogwalker"`
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
