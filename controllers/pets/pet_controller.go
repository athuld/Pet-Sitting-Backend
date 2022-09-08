package pets

import (
	"net/http"
	"pet-sitting-backend/domain/pets"
	"pet-sitting-backend/services"
	"pet-sitting-backend/utils/errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddPet(c *gin.Context) {
	var pet pets.Pet

	if err := c.ShouldBindJSON(&pet); err != nil {
		err := errors.NewBadRequestError("Invalid json")
		c.JSON(err.Status, err)
		return
	}

	user, err := services.GetUserFromJwt(c)
	if err != nil {
		err := errors.NewBadRequestError("Unable to decrypt")
		c.JSON(err.Status, err)
		return
	}
	pet.OwnerID = user.ID

	result, err := services.SavePet(pet)
	if err != nil {
		err := errors.NewBadRequestError("Database error")
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, result)

}

func DeletePet(c *gin.Context) {
	id, err := strconv.ParseInt(c.Query("id"), 10, 64)
	if err != nil {
		err := errors.NewBadRequestError("Query error")
		c.JSON(err.Status, err)
	}
	if err := services.DeletePetByID(id); err != nil {
		err := errors.NewBadRequestError("Database error")
		c.JSON(err.Status, err)
	}
	c.JSON(http.StatusOK, gin.H{"message": "pet record deleted successfully"})
}

func GetAllPets(c *gin.Context) {
	user, err := services.GetUserFromJwt(c)
	if err != nil {
		err := errors.NewBadRequestError("Unable to decrypt")
		c.JSON(err.Status, err)
		return
	}

	var result *[]pets.Pet
    if err != nil {
		err := errors.NewBadRequestError("Query error")
		c.JSON(err.Status, err)
        return
	}

    result,getErr:= services.FetchAllPets(user.ID)
    if getErr!=nil{
        err:= errors.NewBadRequestError("Cannot fetch data")
        c.JSON(err.Status,err)
        return
    }

    c.JSON(http.StatusOK,result)
}
