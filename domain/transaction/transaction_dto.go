package transaction

type Transaction struct {
	TransactionId   int64  `json:"transaction_id"`
	UserId          int64  `json:"user_id"`
	SitterId        int64  `json:"sitter_id"`
	Amount          int    `json:"amount"`
	Charges         int    `json:"charges"`
	TransactionDate string `json:"transaction_date"`
}
