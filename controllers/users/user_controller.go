package users

import (
	"log"
	"net/http"
	"os"
	"pet-sitting-backend/domain/users"
	"pet-sitting-backend/services"
	"pet-sitting-backend/utils/errors"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func Register(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		err := errors.NewInternalServerError("invalid json body")
		c.JSON(err.Status, err)
		return
	}

	result, err := services.CreateUser(user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

func Login(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		err := errors.NewBadRequestError("invalid json")
		c.JSON(err.Status, err)
		return
	}

	result, err := services.GetUser(user)
	if err != nil {
		c.JSON(err.Status, err)
        return
	}
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(result.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
	})

	token, signErr := claims.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if signErr != nil {
		err := errors.NewBadRequestError("login failed")
		c.JSON(err.Status, err)
        return
	}
	c.SetCookie("accessToken", token, 3600, "/", "localhost", false, true)
	c.JSON(http.StatusOK, result)
}

func Logout(c *gin.Context) {
	c.SetCookie("accessToken", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "successfully logged out",
	})
}

func AddUserDetails(c *gin.Context) {
	var userDetails users.UserDetails

	if err := c.ShouldBindJSON(&userDetails); err != nil {
		err := errors.NewBadRequestError("Json is incorrect")
		c.JSON(err.Status, err)
        return
	}

	user, err := services.GetUserFromJwt(c)
	if err != nil {
		err := errors.NewBadRequestError("Failed to find user")
		c.JSON(err.Status, err)
        return
	}
	userDetails.UserID = user.ID
	if err := userDetails.AddDetails(); err != nil {
        log.Fatal(err)
		err := errors.NewBadRequestError("Unable to insert values")
		c.JSON(err.Status, err)
        return
	}
	c.JSON(http.StatusOK, gin.H{"message": "userdetails inserted"})
}

func GetUserDetails(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Query("userID"), 10, 64)
	if err != nil {
		err := errors.NewBadRequestError("Query param error")
		c.JSON(err.Status, err)
        return
	}
	result, getErr := services.GetUserDetails(userID)
	if getErr != nil {
		err := errors.NewBadRequestError("Database error")
		c.JSON(err.Status, err)
        return
	}
	c.JSON(http.StatusOK, result)
}
