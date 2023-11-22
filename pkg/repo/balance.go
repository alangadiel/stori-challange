package repo

import (
	"context"
	"fmt"

	"github.com/alangadiel/stori-challenge/pkg/model"
)

func (r *Repository) GetBalance(ctx context.Context) (model.Balance, error) {
	var balance model.Balance

	// Get total amount
	{
		var err error
		if err = r.dbConn.QueryRow(ctx, `
			SELECT SUM(amount)
			FROM transactions
		`).Scan(&balance.TotalAmount); err != nil {
			return balance, fmt.Errorf("error getting total amount: %w", err)
		}
	}

	// Get monthly balances
	{
		rows, err := r.dbConn.Query(ctx, `
			SELECT
				TO_CHAR(date, 'YYYY-MM') AS month,
				COUNT(*) AS number_of_transactions,
				AVG(CASE WHEN amount < 0 THEN amount END) AS avg_debit,
				AVG(CASE WHEN amount > 0 THEN amount END) AS avg_credit
			FROM transactions
			GROUP BY month
			ORDER BY month
		`)
		if err != nil {
			return balance, fmt.Errorf("error getting monthly balances: %w", err)
		}
		defer rows.Close()

		for rows.Next() {
			var mb model.MonthBalance
			if err := rows.Scan(&mb.Month, &mb.NumberOfTransactions, &mb.AvgDebit, &mb.AvgCredit); err != nil {
				return balance, fmt.Errorf("error scanning monthly balance: %w", err)
			}

			balance.MonthlyBalances = append(balance.MonthlyBalances, mb)
		}
		if err := rows.Err(); err != nil {
			return balance, fmt.Errorf("error iterating monthly balances: %w", err)
		}
	}

	return balance, nil
}
