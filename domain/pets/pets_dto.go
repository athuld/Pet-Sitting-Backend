package pets

type Pet struct {
	ID        int64  `json:"id"`
	OwnerID   int64  `json:"owner_id"`
	PetImg    string `json:"pet_img"`
	PetType   string `json:"pet_type"`
	PetGender string `json:"pet_gender"`
	PetWeight string `json:"pet_weight"`
	PetDesc   string `json:"pet_desc"`
}
