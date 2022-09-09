package services

import (
	sitterresp "pet-sitting-backend/domain/sitter_resp"
	"pet-sitting-backend/utils/errors"
)

func SaveResponse(
	response sitterresp.SitterResponse,
) (*sitterresp.SitterResponse, *errors.RestErr) {
	if err := response.AddResponseToDB(); err != nil {
		return nil, err
	}
	return &response, nil
}

func FetchResponses(req_id int64) (*[]sitterresp.SitterResponseUsers, *errors.RestErr) {
	result := &sitterresp.SitterResponse{SitterReqId: req_id}
	responses, err := result.GetResponsesFromDB()
	if err != nil {
		return nil, err
	}
	return responses, nil
}

func UpdateRequestSitter(response sitterresp.SitterResponse) *errors.RestErr{
    if err:= response.AcceptResponseToDB();err!=nil{
        return err
    }
    return nil
}
