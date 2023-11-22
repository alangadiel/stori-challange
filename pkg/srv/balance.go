package srv

import (
	"context"
	"fmt"
	"strings"

	"github.com/alangadiel/stori-challenge/pkg/model"
	"github.com/alangadiel/stori-challenge/pkg/repo"
)

type BalanceService struct {
	repo.Repository
	EmailService
}

func (s BalanceService) PostBalance(ctx context.Context, fileName, email string) error {
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

	// Send email
	var emailBody strings.Builder
	emailBody.WriteString(fmt.Sprintf("Total Amount: %.2f\n", balance.TotalAmount))
	for _, monthBalance := range balance.MonthlyBalances {
		emailBody.WriteString(fmt.Sprintf("%s: %d transactions, average debit %.2f, average credit %.2f\n",
			monthBalance.Month, monthBalance.NumberOfTransactions, monthBalance.AvgDebit, monthBalance.AvgCredit))
	}
	if err := s.EmailService.SendEmail(email, emailBody.String()); err != nil {
		return fmt.Errorf("error sending email: %w", err)
	}

	return nil
}
