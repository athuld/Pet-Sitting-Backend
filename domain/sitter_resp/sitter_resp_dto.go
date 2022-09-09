package sitterresp

import "pet-sitting-backend/domain/users"

type SitterResponse struct {
	RespID      int64  `json:"resp_id"`
	SitterID    int64  `json:"sitter_id"`
	SitterReqId int64  `json:"sitter_req_id"`
	Prize       int    `json:"prize"`
	Response    string `json:"response"`
}

type SitterResponseUsers struct {
	users.UserDetails
	SitterResponse
}
