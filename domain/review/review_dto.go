package review

type Reviews struct {
	ReviewId int64  `json:"review_id"`
	SitterId int64  `json:"sitter_id"`
	UserId   int64  `json:"user_id"`
	Rating   int    `json:"rating"`
	Review   string `json:"review"`
}

type UserDetails struct {
	Name      string `json:"name"`
	Address   string `json:"address"`
	Pincode   int    `json:"pincode"`
	Phone     int64  `json:"phone"`
	AvatarIMG string `json:"avatar_img"`
}

type ReviewsWithSitter struct {
	Reviews
	AvgRating float32 `json:"avg_rating"`
	UserDetails
}
