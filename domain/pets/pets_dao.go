package pets

import (
	"context"
	"pet-sitting-backend/datasource"
	"pet-sitting-backend/utils/errors"

	"github.com/randallmlough/pgxscan"
)

var (
	queryInsertPet  = "insert into pets (owner_id,pet_img,pet_type,pet_gender,pet_weight,pet_desc) values ($1,$2,$3,$4,$5,$6) returning id,owner_id,pet_type,pet_gender,pet_weight,pet_desc"
	queryDeletePet  = "delete from pets where id=$1"
	queryGetAllPets = "select * from pets where owner_id=$1"
)

func (pet *Pet) SavePetToDB() *errors.RestErr {
	result, err := datasource.Client.Query(
		context.Background(),
		queryInsertPet,
		pet.OwnerID,
		pet.PetImg,
		pet.PetType,
		pet.PetGender,
		pet.PetWeight,
		pet.PetDesc,
	)
	if err != nil {
		return errors.NewBadRequestError("Failed to insert data")
	}
	if err := pgxscan.NewScanner(result).Scan(&pet); err != nil {
		return errors.NewBadRequestError("Failed to scan data")
	}
	return nil
}

func (pet *Pet) DeletePetFromDB() *errors.RestErr {
	_, err := datasource.Client.Exec(context.Background(), queryDeletePet, pet.ID)
	if err != nil {
		return errors.NewBadRequestError("Cannot delete pet record")
	}
	return nil
}

func (pet *Pet) GetAllPetsFromDB() (*[]Pet, *errors.RestErr) {
	result, err := datasource.Client.Query(context.Background(), queryGetAllPets, pet.OwnerID)
	var pets []Pet
	if err != nil {
		return nil, errors.NewBadRequestError("Cannot fetch data")
	}
	if err := pgxscan.NewScanner(result).Scan(&pets); err != nil {
		return nil, errors.NewBadRequestError("Failed to scan data")
	}
	return &pets, nil
}
