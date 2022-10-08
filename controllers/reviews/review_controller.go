package reviews

import (
	"net/http"
	"pet-sitting-backend/domain/review"
	"pet-sitting-backend/services"
	"pet-sitting-backend/utils/errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddReview(c *gin.Context) {
	var review review.Reviews

	user, err := services.GetUserFromJwt(c)
	if err != nil {
		err := errors.NewBadRequestError("Unable to decrypt")
		c.JSON(err.Status, err)
		return
	}

	if err := c.ShouldBindJSON(&review); err != nil {
		err := errors.NewBadRequestError("Json error")
		c.JSON(err.Status, err)
		return
	}

	review.UserId = user.ID
	if err := services.SaveReview(review); err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "review added"})
}

func GetReviewForOwner(c *gin.Context) {
	var review review.Reviews

	user, err := services.GetUserFromJwt(c)
	if err != nil {
		err := errors.NewBadRequestError("Unable to decrypt")
		c.JSON(err.Status, err)
		return
	}

	sitterId, convErr := strconv.ParseInt(c.Query("sitter_id"), 10, 64)
	if convErr != nil {
		err := errors.NewBadRequestError("Unable to decrypt")
		c.JSON(err.Status, err)
		return
	}

	review.UserId = user.ID
	review.SitterId = sitterId

	result, getErr := services.FetchReviewForOwner(review)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, result)
}


func GetReviewsForSitter(c *gin.Context) {
	user, err := services.GetUserFromJwt(c)
	if err != nil {
		err := errors.NewBadRequestError("Unable to decrypt")
		c.JSON(err.Status, err)
		return
	}

	result, getErr := services.FetchAllReviewsForSitter(user.ID)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, result)
}

func GetReviewsForSitterByID(c *gin.Context) {

	sitterId, convErr := strconv.ParseInt(c.Query("sitter_id"), 10, 64)
	if convErr != nil {
		err := errors.NewBadRequestError("Unable to decrypt")
		c.JSON(err.Status, err)
		return
	}

	result, getErr := services.FetchAllReviewsForSitter(sitterId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, result)
}

func GetAllReviews(c *gin.Context){
    var review review.Reviews

    result,err:= review.GetAllReviewsGroup()
    if err!=nil{
        c.JSON(err.Status,err)
        return
    }
    c.JSON(http.StatusOK,result)
}
