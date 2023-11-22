package model

type BalanceRequest struct {
	FileName string `json:"file_name"`
	Email    string `json:"email"`
}

type Balance struct {
	TotalAmount     float64        `json:"total_amount"`
	MonthlyBalances []MonthBalance `json:"monthly_balances"`
}

type MonthBalance struct {
	Month                string  `json:"month"`
	NumberOfTransactions int     `json:"number_of_transactions"`
	AvgDebit             float64 `json:"avg_debit"`
	AvgCredit            float64 `json:"avg_credit"`
}
