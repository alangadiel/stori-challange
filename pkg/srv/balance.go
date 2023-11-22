package srv

import (
	"context"
	"fmt"

	"github.com/alangadiel/stori-challenge/pkg/model"
)

func (s Service) PostBalance(ctx context.Context, fileName, email string) error {
	var transactions []model.Transaction
	{
		var err error
		if transactions, err = readTransactionsFile(fileName); err != nil {
			return fmt.Errorf("error reading transactions file: %w", err)
		}
	}

	// Save transactions
	if err := s.Repository.SaveTransactions(ctx, transactions); err != nil {
		return fmt.Errorf("error saving transactions: %w", err)
	}

	// Get transactions balance
	var balance model.Balance
	{
		var err error
		if balance, err = s.Repository.GetBalance(ctx); err != nil {
			return fmt.Errorf("error getting balance: %w", err)
		}
	}

	fmt.Println(balance)

	// Send email

	return nil
}
