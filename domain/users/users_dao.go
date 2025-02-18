package users

import (
	"context"
	"pet-sitting-backend/datasource"
	sitterreq "pet-sitting-backend/domain/sitter_req"
	"pet-sitting-backend/utils/errors"
	"pet-sitting-backend/utils/logger"

	"github.com/randallmlough/pgxscan"
)

var (
	queryInsertUser         = "insert into users(username,email,password) values($1,$2,$3) returning id,username,email,password"
	queryGetUserByEmail     = "select id,username,email,password from users where email=$1"
	queryGetUserById        = "select id,username,email from users where id=$1"
	queryAddUserDetails     = "insert into userdetails(user_id,name,gender,age,phone,address,pincode,is_petsitter,avatar_img) values ($1,$2,$3,$4,$5,$6,$7,$8,$9)"
	queryGetUserDetails     = "select * from userdetails where user_id=$1"
	queryActiveRequestByPin = "select s.*,p.*,ud.address,ud.pincode,ud.phone from sitter_reqs s inner join pets p on s.pet_id=p.id inner join userdetails ud on s.user_id=ud.user_id where ud.pincode between $1 and $2 and s.user_id!=$3 and s.req_id not in (select sitter_req_id from sitter_resps where sitter_id=$3) and is_accepted=false;"
    queryGetAllUsers        = "select user_id,username,name,email,gender,age,address,pincode,phone,is_petsitter from users u inner join userdetails ud on u.id=ud.user_id"
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
			logger.Error.Println(getErr)
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
		logger.Error.Println(err)
		return errors.NewInternalServerError("Database error")
	}

	for result.Next() {
		if getErr := result.Scan(&user.ID, &user.Username, &user.Email, &user.Password); getErr != nil {
			logger.Error.Println(getErr)
			return errors.NewInternalServerError("Databae error")
		}
	}
	return nil
}

func (user *User) GetById() *errors.RestErr {
	result, err := datasource.Client.Query(context.Background(), queryGetUserById, user.ID)
	if err != nil {
		logger.Error.Println(err)
		return errors.NewBadRequestError("database error")
	}
	for result.Next() {
		if getErr := result.Scan(&user.ID, &user.Username, &user.Email); getErr != nil {
			logger.Error.Println(getErr)
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
        userDetails.Name,
		userDetails.Gender,
		userDetails.Age,
        userDetails.Phone,
		userDetails.Address,
		userDetails.Pincode,
		userDetails.IsPetsitter,
        userDetails.AvatarIMG,
	)
	if err != nil {
		logger.Error.Println(err)
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
		logger.Error.Println(err)
		return errors.NewBadRequestError("database error")
	}
	if getErr := pgxscan.NewScanner(result).Scan(&userDetails); getErr != nil {
		logger.Error.Println(getErr)
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
        user.UserID,
	)
	if err != nil {
		logger.Error.Println(err)
		return nil, errors.NewBadRequestError("Cannot fetch data")
	}
	var activer_reqs_pins []sitterreq.SitterPetsUsers
	if err := pgxscan.NewScanner(result).Scan(&activer_reqs_pins); err != nil {
		logger.Error.Println(err)
		return nil, errors.NewBadRequestError("Failed to scan")
	}
	return &activer_reqs_pins, nil
}

func (user *User) GetAllUsers() (*[]FullUserDetails,*errors.RestErr){
    result,err:= datasource.Client.Query(context.Background(),queryGetAllUsers)
	if err != nil {
		logger.Error.Println(err)
		return nil, errors.NewBadRequestError("Cannot fetch data")
	}
    var allUsers []FullUserDetails
	if err := pgxscan.NewScanner(result).Scan(&allUsers); err != nil {
		logger.Error.Println(err)
		return nil, errors.NewBadRequestError("Failed to scan")
	}
	return &allUsers, nil

}
