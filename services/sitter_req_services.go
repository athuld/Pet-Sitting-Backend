package services

import (
	sitterreq "pet-sitting-backend/domain/sitter_req"
	"pet-sitting-backend/utils/errors"
)

func CreateSitterRequest(sitter_req sitterreq.SitterReq) (*sitterreq.SitterReq, *errors.RestErr) {
	if err := sitter_req.AddRequestToDB(); err != nil {
		return nil, err
	}
	return &sitter_req, nil
}

func FetchActiveRequests(
	sitter_req sitterreq.SitterReq,
) (*[]sitterreq.SitterPets, *errors.RestErr) {
	request := &sitterreq.SitterReq{UserId: sitter_req.UserId}
	result, err := request.GetActiveRequestsFromDB()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func FetchInActiveRequests(
	sitter_req sitterreq.SitterReq,
) (*[]sitterreq.SitterPetsUsers, *errors.RestErr) {
	request := &sitterreq.SitterReq{UserId: sitter_req.UserId}
	result, err := request.GetInActiveRequestsFromDB()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func RemoveRequest(sitter_req sitterreq.SitterReq) *errors.RestErr {
	request := &sitterreq.SitterReq{ReqId: sitter_req.ReqId}
	if err := request.DeleteRequestFromDB(); err != nil {
		return err
	}
	return nil
}
