package services

import (
	"pet-sitting-backend/domain/pets"
	"pet-sitting-backend/utils/errors"
	"pet-sitting-backend/utils/logger"
)

func SavePet(pet pets.Pet) (*pets.Pet, *errors.RestErr) {
	if err := pet.SavePetToDB(); err != nil {
		logger.Error.Println(err)
		return nil, errors.NewBadRequestError("Database error")
	}
	return &pet, nil
}

func DeletePetByID(id int64) *errors.RestErr {
	pet := &pets.Pet{ID: id}
	if err := pet.DeletePetFromDB(); err != nil {
		logger.Error.Println(err)
		return errors.NewBadRequestError("Database error")
	}
	return nil
}

func FetchAllPets(owner_id int64) (*[]pets.Pet, *errors.RestErr) {
	pet := &pets.Pet{OwnerID: owner_id}
	var result *[]pets.Pet

	result, err := pet.GetAllPetsFromDB()
	if err != nil {
		logger.Error.Println(err)
		return nil, errors.NewBadRequestError("Database error")
	}

	return result, nil
}
