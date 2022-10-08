package sitterreq

import (
	"pet-sitting-backend/domain/pets"
)

type SitterReq struct {
	ReqId        int64  `json:"req_id"`
	PetId        int64  `json:"pet_id"`
	UserId       int64  `json:"user_id"`
	Date         string `json:"date"`
	Time         string `json:"time"`
	Instructions string `json:"instructions"`
	BasePrize    int    `json:"base_prize"`
	IsAccepted   bool   `json:"is_accepted"`
	SitterId     int64  `json:"sitter_id"`
	IsPersonal   bool   `json:"is_personal"`
}

type UserDetails struct {
	Name      string `json:"name"`
	Address   string `json:"address"`
	Pincode   int    `json:"pincode"`
	Phone     int64  `json:"phone"`
	AvatarIMG string `json:"avatar_img"`
}

type SitterPets struct {
	SitterReq
	pets.Pet
}

type SitterPetsUsers struct {
	SitterReq
	pets.Pet
	UserDetails
}

type AcceptedRequests struct {
	UserDetails
	PetImg string `json:"pet_img"`
	Date   string `json:"date"`
	Time   string `json:"time"`
	Prize  int    `json:"prize"`
}
