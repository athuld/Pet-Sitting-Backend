package sitterreq

import (
	"context"
	"pet-sitting-backend/datasource"
	"pet-sitting-backend/utils/errors"
	"pet-sitting-backend/utils/logger"

	"github.com/randallmlough/pgxscan"
)

var (
	queryAddRequest    = "insert into sitter_reqs (pet_id,user_id,date,time,instructions,base_prize) values ($1,$2,$3,$4,$5,$6) returning req_id,pet_id,user_id,date,time,instructions,base_prize,is_accepted"
	queryDeleteRequest = "delete from sitter_reqs where req_id=$1"
	queryActiveRequest = "select * from sitter_reqs s inner join pets p on s.pet_id=p.id where user_id=$1 and is_accepted=false;"
	queryInActiveRequest = "select s.*,p.*,ud.name,ud.phone,ud.pincode from sitter_reqs s inner join pets p on s.pet_id=p.id inner join userdetails ud on ud.user_id=s.sitter_id where s.user_id=$1 and is_accepted=true;"
)

func (sitter_req *SitterReq) AddRequestToDB() *errors.RestErr {

	result, err := datasource.Client.Query(
		context.Background(),
		queryAddRequest,
		sitter_req.PetId,
		sitter_req.UserId,
		sitter_req.Date,
		sitter_req.Time,
		sitter_req.Instructions,
		sitter_req.BasePrize,
	)
	if err != nil {
		logger.Error.Println(err)
		return errors.NewBadRequestError("Cannot insert values")
	}
	if scanErr := pgxscan.NewScanner(result).Scan(&sitter_req); scanErr != nil {
		logger.Error.Println(scanErr)
		return errors.NewBadRequestError("Cannot scan struct")
	}
	return nil
}

func (sitter_req *SitterReq) DeleteRequestFromDB() *errors.RestErr {
	_, err := datasource.Client.Exec(context.Background(), queryDeleteRequest, sitter_req.ReqId)
	if err != nil {
		logger.Error.Println(err)
		return errors.NewBadRequestError("Cannot delete values")
	}
	return nil
}

func (sitter_req *SitterReq) GetActiveRequestsFromDB() (*[]SitterPets, *errors.RestErr) {
	result, err := datasource.Client.Query(
		context.Background(),
		queryActiveRequest,
		sitter_req.UserId,
	)
	if err != nil {
		logger.Error.Println(err)
		return nil, errors.NewBadRequestError("Database error")
	}
	var activer_reqs []SitterPets
	if err := pgxscan.NewScanner(result).Scan(&activer_reqs); err != nil {
		logger.Error.Println(err)
		return nil, errors.NewBadRequestError("Failed to scan")
	}
	return &activer_reqs, nil
}

func (sitter_req *SitterReq) GetInActiveRequestsFromDB() (*[]SitterPetsUsers, *errors.RestErr) {
	result, err := datasource.Client.Query(
		context.Background(),
		queryInActiveRequest,
		sitter_req.UserId,
	)
	if err != nil {
		logger.Error.Println(err)
		return nil, errors.NewBadRequestError("Database error")
	}
	var activer_reqs []SitterPetsUsers
	if err := pgxscan.NewScanner(result).Scan(&activer_reqs); err != nil {
		logger.Error.Println(err)
		return nil, errors.NewBadRequestError("Failed to scan")
	}
	return &activer_reqs, nil
}
