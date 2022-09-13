package sitterreqs

import (
	"net/http"
	sitterreq "pet-sitting-backend/domain/sitter_req"
	"pet-sitting-backend/services"
	"pet-sitting-backend/utils/errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddRequest(c *gin.Context) {
	var sitter_req sitterreq.SitterReq

	user, err := services.GetUserFromJwt(c)
	if err != nil {
		err := errors.NewBadRequestError("Unable to decrypt")
		c.JSON(err.Status, err)
		return
	}

	if err := c.ShouldBindJSON(&sitter_req); err != nil {
		err := errors.NewBadRequestError("Json error")
		c.JSON(err.Status, err)
		return
	}
	sitter_req.UserId = user.ID

	result, getErr := services.CreateSitterRequest(sitter_req)
	if getErr != nil {
		err := errors.NewBadRequestError("Database error")
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

func GetActiveRequests(c *gin.Context) {
	var sitter_req sitterreq.SitterReq

	user, err := services.GetUserFromJwt(c)
	if err != nil {
		err := errors.NewBadRequestError("Unable to decrypt")
		c.JSON(err.Status, err)
		return
	}

	sitter_req.UserId = user.ID
	result, getErr := services.FetchActiveRequests(sitter_req)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, result)
}


func GetInActiveRequests(c *gin.Context) {
	var sitter_req sitterreq.SitterReq

	user, err := services.GetUserFromJwt(c)
	if err != nil {
		err := errors.NewBadRequestError("Unable to decrypt")
		c.JSON(err.Status, err)
		return
	}

	sitter_req.UserId = user.ID
	result, getErr := services.FetchInActiveRequests(sitter_req)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, result)
}


func DeleteRequest(c *gin.Context) {
	var sitter_req sitterreq.SitterReq
	req_id, err := strconv.ParseInt(c.Query("req_id"), 10, 64)
	if err != nil {
		err := errors.NewBadRequestError("Query param error")
		c.JSON(err.Status, err)
		return
	}
	sitter_req.ReqId = req_id
	if getErr := services.RemoveRequest(sitter_req); getErr != nil {
		c.JSON(getErr.Status, err)
		return
	}
}
