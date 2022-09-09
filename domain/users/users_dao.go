package users

import (
	"context"
	"log"
	"pet-sitting-backend/datasource"
	sitterreq "pet-sitting-backend/domain/sitter_req"
	"pet-sitting-backend/utils/errors"

	"github.com/randallmlough/pgxscan"
)

var (
	queryInsertUser         = "insert into users(username,email,password) values($1,$2,$3) returning id,username,email,password"
	queryGetUserByEmail     = "select id,username,email,password from users where email=$1"
	queryGetUserById        = "select id,username,email from users where id=$1"
	queryAddUserDetails     = "insert into userdetails(user_id,gender,age,address,pincode,is_petsitter,is_dogwalker) values ($1,$2,$3,$4,$5,$6,$7)"
	queryGetUserDetails     = "select * from userdetails where user_id=$1"
	queryActiveRequestByPin = "select s.*,p.*,ud.address,ud.pincode from sitter_reqs s inner join pets p on s.pet_id=p.id inner join userdetails ud on s.user_id=ud.user_id where ud.pincode between $1 and $2;"
)

func (user *User) Save() *errors.RestErr {
	result, err := datasource.Client.Query(
		context.Background(),
		queryInsertUser,
		user.Username,
		user.Email,
		user.Password,
	)
	for result.Next() {
		if getErr := result.Scan(&user.ID, &user.Username, &user.Email, &user.Password); getErr != nil {
			log.Fatal(getErr)
			return errors.NewInternalServerError("Databae error")
		}
	}
	if err != nil {
		return errors.NewInternalServerError("Database error")
	}

	return nil
}

func (user *User) GetByEmail() *errors.RestErr {
	result, err := datasource.Client.Query(context.Background(), queryGetUserByEmail, user.Email)
	if err != nil {
		log.Fatal(err)
		return errors.NewInternalServerError("Database error")
	}

	for result.Next() {
		if getErr := result.Scan(&user.ID, &user.Username, &user.Email, &user.Password); getErr != nil {
			log.Fatal(getErr)
			return errors.NewInternalServerError("Databae error")
		}
	}
	return nil
}

func (user *User) GetById() *errors.RestErr {
	result, err := datasource.Client.Query(context.Background(), queryGetUserById, user.ID)
	if err != nil {
		return errors.NewBadRequestError("database error")
	}
	for result.Next() {
		if getErr := result.Scan(&user.ID, &user.Username, &user.Email); getErr != nil {
			log.Fatal(getErr)
			return errors.NewInternalServerError("Databae error")
		}
	}
	return nil
}

func (userDetails *UserDetails) AddDetails() *errors.RestErr {
	_, err := datasource.Client.Query(
		context.Background(),
		queryAddUserDetails,
		userDetails.UserID,
		userDetails.Gender,
		userDetails.Age,
		userDetails.Address,
		userDetails.Pincode,
		userDetails.IsPetsitter,
		userDetails.IsDogwalker,
	)
	if err != nil {
		log.Fatal(err)
		return errors.NewBadRequestError("database error")
	}
	return nil
}

func (userDetails *UserDetails) GetDetailsByID() *errors.RestErr {
	result, err := datasource.Client.Query(
		context.Background(),
		queryGetUserDetails,
		userDetails.UserID,
	)
	if err != nil {
		return errors.NewBadRequestError("database error")
	}
	if getErr := pgxscan.NewScanner(result).Scan(&userDetails); getErr != nil {
		log.Fatal(getErr)
		return errors.NewBadRequestError("Error is here")
	}
	return nil
}

func (user *UserDetails) GetActiverRequestsByPinFromDB() (*[]sitterreq.SitterPetsUsers, *errors.RestErr) {
	low_pin := user.Pincode - 2
	high_pin := user.Pincode + 2
	result, err := datasource.Client.Query(
		context.Background(),
		queryActiveRequestByPin,
		low_pin,
		high_pin,
	)
	if err != nil {
		return nil, errors.NewBadRequestError("Cannot fetch data")
	}
	var activer_reqs_pins []sitterreq.SitterPetsUsers
	if err := pgxscan.NewScanner(result).Scan(&activer_reqs_pins); err != nil {
		return nil, errors.NewBadRequestError("Failed to scan")
	}
	return &activer_reqs_pins, nil
}
