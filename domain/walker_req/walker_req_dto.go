package walkerreq

type WalkerReq struct {
	ReqId        int64  `json:"req_id"`
	PetId        int64  `json:"pet_id"`
	OwnerId      int64  `json:"owner_id"`
	Date         string `json:"date"`
	Time         string `json:"time"`
	Instructions string `json:"instructions"`
	BasePrize    int    `json:"base_prize"`
	IsAccepted   bool   `json:"is_accepted"`
	WalkerId     int64  `json:"sitter_id"`
}
