package sitterresp

import (
	"context"
	"pet-sitting-backend/datasource"
	"pet-sitting-backend/utils/errors"
	"pet-sitting-backend/utils/logger"

	"github.com/randallmlough/pgxscan"
)

var (
	queryInsertResponse   = "insert into sitter_resps (sitter_id,sitter_req_id,prize,response) values($1,$2,$3,$4) returning resp_id,sitter_id,sitter_req_id,prize,response"
	queryGetResponsesById = "select * from sitter_resps sr inner join userdetails ud on sr.sitter_id=ud.user_id where sitter_req_id=$1"
	queryAcceptResponse   = "update sitter_reqs set is_accepted=true,base_prize=$1,sitter_id=$2 where req_id=$3"
)

func (response *SitterResponse) AddResponseToDB() *errors.RestErr {
	result, err := datasource.Client.Query(
		context.Background(),
		queryInsertResponse,
		response.SitterID,
		response.SitterReqId,
		response.Prize,
		response.Response,
	)
	if err != nil {
		logger.Error.Println(err)
		return errors.NewBadRequestError("Database error")
	}
	if err := pgxscan.NewScanner(result).Scan(&response); err != nil {
		return errors.NewBadRequestError("Failed to scan")
	}
	return nil
}

func (response *SitterResponse) GetResponsesFromDB() (*[]SitterResponseUsers, *errors.RestErr) {
	result, err := datasource.Client.Query(
		context.Background(),
		queryGetResponsesById,
		response.SitterReqId,
	)
	if err != nil {
		logger.Error.Println(err)
		return nil, errors.NewBadRequestError("Database error")
	}
	var responses []SitterResponseUsers
	if err := pgxscan.NewScanner(result).Scan(&responses); err != nil {
		logger.Error.Println(err)
		return nil, errors.NewBadRequestError("Failed to scan")
	}
	return &responses, nil
}
func (response *SitterResponse) AcceptResponseToDB() *errors.RestErr {
	_, err := datasource.Client.Query(
		context.Background(),
		queryAcceptResponse,
		response.Prize,
		response.SitterID,
		response.SitterReqId,
	)
	if err != nil {
		logger.Error.Println(err)
		return errors.NewBadRequestError("Cannot update data")
	}
	return nil
}
