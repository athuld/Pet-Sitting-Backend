package sitterresps

import (
	"net/http"
	sitterresp "pet-sitting-backend/domain/sitter_resp"
	"pet-sitting-backend/services"
	"pet-sitting-backend/utils/errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddResponse(c *gin.Context) {
	var response sitterresp.SitterResponse

	user, err := services.GetUserFromJwt(c)
	if err != nil {
		err := errors.NewBadRequestError("Unable to decrypt")
		c.JSON(err.Status, err)
		return
	}

	req_id, conErr := strconv.ParseInt(c.Query("req_id"), 10, 64)
	if conErr != nil {
		err := errors.NewBadRequestError("Query error")
		c.JSON(err.Status, err)
		return
	}

	if err := c.ShouldBindJSON(&response); err != nil {
		err := errors.NewBadRequestError("Json error")
		c.JSON(err.Status, err)
		return
	}
	response.SitterID = user.ID
	response.SitterReqId = req_id

	result, getErr := services.SaveResponse(response)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, result)
}

func GetResponsesById(c *gin.Context) {
	req_id, err := strconv.ParseInt(c.Query("req_id"), 10, 64)
	if err != nil {
		err := errors.NewBadRequestError("Query Error")
		c.JSON(err.Status, err)
		return
	}
	result, getErr := services.FetchResponses(req_id)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, result)
}

func AcceptResponse(c *gin.Context) {
	var response sitterresp.SitterResponse

	if err := c.ShouldBindJSON(&response); err != nil {
		err := errors.NewBadRequestError("Json error")
		c.JSON(err.Status, err)
		return
	}
	if upErr := services.UpdateRequestSitter(response); upErr != nil {
		c.JSON(upErr.Status, upErr)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "response accepted"})
}
