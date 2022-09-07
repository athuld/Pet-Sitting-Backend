package users

import (
	"context"
	"log"
	"pet-sitting-backend/datasource"
	"pet-sitting-backend/utils/errors"
)

var (
	queryInsertUser     = "insert into users(username,email,password) values($1,$2,$3) returning id,username,email,password"
	queryGetUserByEmail = "select id,username,email,password from users where email=$1"
	queryGetUserById    = "select id,username,email from users where id=$1"
)

func (user *User) Save() *errors.RestErr {
	result, err := datasource.Client.Query(
		context.Background(),
		queryInsertUser,
		user.Username,
		user.Email,
		user.Password,
	)
    for result.Next(){
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

    for result.Next(){
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
    for result.Next(){
        if getErr := result.Scan(&user.ID, &user.Username, &user.Email); getErr != nil {
            log.Fatal(getErr)
            return errors.NewInternalServerError("Databae error")
        }
    }
	return nil
}
