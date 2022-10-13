package transaction

import (
	"context"
	"pet-sitting-backend/datasource"
	"pet-sitting-backend/utils/errors"
	"pet-sitting-backend/utils/logger"

	"github.com/randallmlough/pgxscan"
)

var (
	queryAddTransaction    = "insert into transaction(user_id,sitter_id,sitter_req_id,amount,charges,transaction_date) values($1,$2,$3,$4,$5,$6)"
	queryGetGeneralExpense = "select sum(amount+charges) as expense,pet_name from transaction where user_id=$1 group by pet_name"
    queryGetCustomExpense = "select sum(amount+charges) as expense,pet_name from transaction where user_id=$1 and transaction_date between $2 and $3 group by pet_name"
    queryGetGeneralEarnings = "select count(*) as transaction_count,sum(amount) as total_amount from transaction where sitter_id=$1 group by sitter_id"
    queryGetCustomEarnings = "select count(*) as transaction_count,sum(amount) as total_amount from transaction where sitter_id=$1 and transaction_date between $2 and $3 group by sitter_id"
    queryGetRevenueInfo = "select transaction_id,amount,charges,transaction_date,sum(charges) as total_revenue from transaction group by transaction_id order by transaction_id desc"
    queryGetAllTransactions ="select * from transaction"
)

func (trn *Transaction) AddTransactionToDB() *errors.RestErr {
	_, err := datasource.Client.Query(
		context.Background(),
		queryAddTransaction,
		trn.UserId,
		trn.SitterId,
		trn.SitterReqId,
		trn.Amount,
		trn.Charges,
		trn.TransactionDate,
	)
	if err != nil {
		logger.Error.Println(err)
		return errors.NewBadRequestError("Database error")
	}
	return nil
}

func (trn *Transaction) FetchGeneralExpense() (*[]GeneralExpense, *errors.RestErr) {
	result, err := datasource.Client.Query(context.Background(), queryGetGeneralExpense, trn.UserId)
	if err != nil {
		logger.Error.Println(err)
		return nil, errors.NewBadRequestError("Database error")
	}
	var expense []GeneralExpense
	logger.Error.Println(err)
	if err := pgxscan.NewScanner(result).Scan(&expense); err != nil {
		return nil, errors.NewBadRequestError("Failed to scan")
	}
	return &expense, nil

}

func (cust *CustomExpenseInfo) FetchCustomExpense() (*[]GeneralExpense, *errors.RestErr) {
	result, err := datasource.Client.Query(context.Background(), queryGetCustomExpense, cust.UserId,cust.FromDate,cust.ToDate)
	if err != nil {
		logger.Error.Println(err)
		return nil, errors.NewBadRequestError("Database error")
	}
	var expense []GeneralExpense
	logger.Error.Println(err)
	if err := pgxscan.NewScanner(result).Scan(&expense); err != nil {
		return nil, errors.NewBadRequestError("Failed to scan")
	}
	return &expense, nil
}


func (trn *Transaction) FetchGeneralEarnings() (*GeneralEarnings, *errors.RestErr) {
	result, err := datasource.Client.Query(context.Background(), queryGetGeneralEarnings, trn.SitterId,)
	if err != nil {
		logger.Error.Println(err)
		return nil, errors.NewBadRequestError("Database error")
	}
	var earnings GeneralEarnings
	logger.Error.Println(err)
	if err := pgxscan.NewScanner(result).Scan(&earnings); err != nil {
		return nil, errors.NewBadRequestError("Failed to scan")
	}
	return &earnings, nil

}
func (cust *CustomEarningsInfo) FetchCustomEarnings() (*GeneralEarnings, *errors.RestErr) {
	result, err := datasource.Client.Query(context.Background(), queryGetCustomEarnings, cust.SitterId,cust.FromDate,cust.ToDate)
	if err != nil {
		logger.Error.Println(err)
		return nil, errors.NewBadRequestError("Database error")
	}
	var earnings GeneralEarnings
	logger.Error.Println(err)
	if err := pgxscan.NewScanner(result).Scan(&earnings); err != nil {
		return nil, errors.NewBadRequestError("Failed to scan")
	}
	return &earnings, nil
}

func (revenue *RevenueInfo) FetchRevenueInfo() (*[]RevenueInfo,*errors.RestErr){
	result, err := datasource.Client.Query(context.Background(), queryGetRevenueInfo)
	if err != nil {
		logger.Error.Println(err)
		return nil, errors.NewBadRequestError("Database error")
	}
	var revInfo []RevenueInfo
	logger.Error.Println(err)
	if err := pgxscan.NewScanner(result).Scan(&revInfo); err != nil {
		return nil, errors.NewBadRequestError("Failed to scan")
	}
	return &revInfo, nil
}

func (trn *Transaction) FetchAllTransactions() (*[]Transaction,*errors.RestErr){
	result, err := datasource.Client.Query(context.Background(), queryGetAllTransactions)
	if err != nil {
		logger.Error.Println(err)
		return nil, errors.NewBadRequestError("Database error")
	}
	var transactions []Transaction
	logger.Error.Println(err)
	if err := pgxscan.NewScanner(result).Scan(&transactions); err != nil {
		return nil, errors.NewBadRequestError("Failed to scan")
	}
	return &transactions, nil
}
