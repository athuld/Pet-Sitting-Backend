package transactions

import (
	"net/http"
	"pet-sitting-backend/domain/transaction"
	"pet-sitting-backend/services"
	"pet-sitting-backend/utils/errors"
	"pet-sitting-backend/utils/logger"

	"github.com/gin-gonic/gin"
)

func AddNewTransaction(c *gin.Context) {
	var tran transaction.Transaction

	if err := c.ShouldBindJSON(&tran); err != nil {
		err := errors.NewBadRequestError("Json is incorrect")
		c.JSON(err.Status, err)
		return
	}

	user, err := services.GetUserFromJwt(c)
	if err != nil {
		logger.Error.Println(err)
		err := errors.NewBadRequestError("Failed to find user")
		c.JSON(err.Status, err)
		return
	}
	tran.UserId = user.ID

	if err := tran.AddTransactionToDB(); err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "transaction addedd successfully"})
}

func GetGeneralExpense(c *gin.Context) {
	var tran transaction.Transaction
	user, err := services.GetUserFromJwt(c)
	if err != nil {
		logger.Error.Println(err)
		err := errors.NewBadRequestError("Failed to find user")
		c.JSON(err.Status, err)
		return
	}
	tran.UserId = user.ID
	result, getErr := tran.FetchGeneralExpense()
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, result)

}

func GetCustomExpense(c *gin.Context) {
	var cust transaction.CustomExpenseInfo

	if err := c.ShouldBindJSON(&cust); err != nil {
		err := errors.NewBadRequestError("Json is incorrect")
		c.JSON(err.Status, err)
		return
	}

	user, err := services.GetUserFromJwt(c)
	if err != nil {
		logger.Error.Println(err)
		err := errors.NewBadRequestError("Failed to find user")
		c.JSON(err.Status, err)
		return
	}
	cust.UserId = user.ID
	logger.Info.Println(cust)
	result, getErr := cust.FetchCustomExpense()

	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, result)

}

func GetGeneralEarnings(c *gin.Context) {
	var tran transaction.Transaction
	user, err := services.GetUserFromJwt(c)
	if err != nil {
		logger.Error.Println(err)
		err := errors.NewBadRequestError("Failed to find user")
		c.JSON(err.Status, err)
		return
	}
	tran.SitterId = user.ID
	result, getErr := tran.FetchGeneralEarnings()
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, result)
}

func GetCustomEranings(c *gin.Context) {
	var cust transaction.CustomEarningsInfo

	if err := c.ShouldBindJSON(&cust); err != nil {
		err := errors.NewBadRequestError("Json is incorrect")
		c.JSON(err.Status, err)
		return
	}

	user, err := services.GetUserFromJwt(c)
	if err != nil {
		logger.Error.Println(err)
		err := errors.NewBadRequestError("Failed to find user")
		c.JSON(err.Status, err)
		return
	}
	cust.SitterId = user.ID
	logger.Info.Println(cust)
	result, getErr := cust.FetchCustomEarnings()

	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, result)

}

func GetRevenueInfo(c *gin.Context) {
	var rev transaction.RevenueInfo
	result, getErr := rev.FetchRevenueInfo()
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, result)
}

func GetAllTransactions(c *gin.Context) {
	var tran transaction.Transaction
	result, getErr := tran.FetchAllTransactions()
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, result)
}
