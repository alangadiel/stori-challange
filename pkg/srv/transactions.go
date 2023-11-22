package srv

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/alangadiel/stori-challenge/pkg/model"
)

const (
	transactionsFile = "transactions.csv"
)

func (s Service) PostTransactions(ctx context.Context) error {
	var transactions []model.Transaction
	{
		var err error
		if transactions, err = readTransactionsFile(); err != nil {
			return fmt.Errorf("error reading transactions file: %w", err)
		}
	}

	// Save transactions
	if err := s.Repository.SaveTransactions(ctx, transactions); err != nil {
		return fmt.Errorf("error saving transactions: %w", err)
	}

	// Get transactions balance

	// Send email

	return nil
}

func readTransactionsFile() ([]model.Transaction, error) {
	// open file transactions.csv
	file, err := os.Open(transactionsFile)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	var transactions []model.Transaction

	// read csv
	reader := csv.NewReader(file)

	// skip header
	if _, err := reader.Read(); err != nil {
		return nil, fmt.Errorf("error reading csv: %w", err)
	}

	for {
		var record []string
		{
			var err error
			if record, err = reader.Read(); err != nil {
				if err == io.EOF {
					break
				}
				return nil, fmt.Errorf("error reading csv: %w", err)
			}
		}

		var id int
		{
			var err error
			if id, err = strconv.Atoi(record[0]); err != nil {
				return nil, fmt.Errorf("error converting ID to int: %w", err)
			}
		}

		var date time.Time
		{
			var err error
			dateStr := fmt.Sprintf("%v/%s", time.Now().Year(), record[1])
			if date, err = time.Parse("2006/01/02", dateStr); err != nil {
				return nil, fmt.Errorf("error converting date to time.Time: %w", err)
			}
		}

		var amount float64
		{
			var err error
			if amount, err = strconv.ParseFloat(record[2], 64); err != nil {
				return nil, fmt.Errorf("error converting amount to float64: %w", err)
			}
		}

		transactions = append(transactions, model.Transaction{
			ID:     id,
			Date:   date,
			Amount: amount,
		})
	}

	return transactions, nil
}
