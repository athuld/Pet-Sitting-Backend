package services

import (
	"os"
	sitterreq "pet-sitting-backend/domain/sitter_req"
	"pet-sitting-backend/domain/users"
	"pet-sitting-backend/utils/errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.ValidateUser(); err != nil {
		return nil, err
	}

	pwdSlice, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return nil, errors.NewBadRequestError("failed to encrypt password")
	}
	user.Password = string(pwdSlice[:])

	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUser(user users.User) (*users.User, *errors.RestErr) {
	result := &users.User{Email: user.Email}
	if err := result.GetByEmail(); err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(user.Password)); err != nil {
		return nil, errors.NewBadRequestError("Passwords do not match")
	}

	resultWpwd := &users.User{ID: result.ID, Username: result.Username, Email: result.Email}
	return resultWpwd, nil
}

func GetUserFromJwt(c *gin.Context) (*users.User, *errors.RestErr) {
	cookie, err := c.Cookie("accessToken")
	if err != nil {
		return nil, errors.NewBadRequestError("Cannot get cookie")
	}
	token, err := jwt.ParseWithClaims(
		cookie,
		&jwt.StandardClaims{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET_KEY")), nil
		},
	)

	if err != nil {
		return nil, errors.NewBadRequestError("error parsing cookie")
	}
	claims := token.Claims.(*jwt.StandardClaims)
	issuer, err := strconv.ParseInt(claims.Issuer, 10, 64)
	if err != nil {
		return nil, errors.NewBadRequestError("id should be an int")
	}

	result := &users.User{ID: issuer}
	if err := result.GetById(); err != nil {
		return nil, err
	}
	return result, nil
}

func AddUserDetails(userDetails users.UserDetails) (*errors.RestErr){
    err:= userDetails.AddDetails()
    if err!=nil{
        return err
    }
    return nil
}

func GetUserDetails(userId int64) (*users.UserDetails,*errors.RestErr){
    result := &users.UserDetails{UserID: userId}
    if err:= result.GetDetailsByID();err!=nil{
        return nil,err
    }
    return result,nil
}

func FetchActiveRequestsByPincode(pincode int)(*[]sitterreq.SitterPetsUsers,*errors.RestErr){
    request:= &users.UserDetails{Pincode: pincode}
    result,err := request.GetActiverRequestsByPinFromDB()
    if err!=nil{
        return nil,err
    }
    return result,nil
}
