package transaction

import "github.com/jackc/pgtype"

type Transaction struct {
	TransactionId   int64       `json:"transaction_id"`
	UserId          int64       `json:"user_id"`
	SitterId        int64       `json:"sitter_id"`
	SitterReqId     int64       `json:"sitter_req_id"`
	Amount          int         `json:"amount"`
	Charges         int         `json:"charges"`
	TransactionDate pgtype.Date `json:"transaction_date"`
	PetName         string      `json:"pet_name"`
}

type GeneralExpense struct {
	PetName string `json:"pet_name"`
	Expense int    `json:"expense"`
}

type GeneralEarnings struct {
	TransactionCount int `json:"transaction_count"`
	TotalAmount      int `json:"total_amount"`
}

type CustomExpenseInfo struct {
	UserId   int64  `json:"user_id"`
	FromDate string `json:"from_date"`
	ToDate   string `json:"to_date"`
}

type CustomEarningsInfo struct {
	SitterId int64  `json:"sitter_id"`
	FromDate string `json:"from_date"`
	ToDate   string `json:"to_date"`
}

type RevenueInfo struct {
	TransactionId   int64       `json:"transaction_id"`
	Amount          int         `json:"amount"`
	Charges         int         `json:"charges"`
	TransactionDate pgtype.Date `json:"transaction_date"`
	TotalRevenue    int         `json:"total_revenue"`
}
