package repo

import (
	"context"
	"fmt"

	"github.com/alangadiel/stori-challenge/pkg/model"
	"github.com/jackc/pgx/v5"
)

func (r *Repository) SaveTransactions(ctx context.Context, trs []model.Transaction) error {
	_, err := r.dbConn.CopyFrom(ctx,
		pgx.Identifier{"transactions"},
		[]string{"id", "date", "amount"},
		CopyFromTransactions(trs))

	if err != nil {
		return fmt.Errorf("error saving transactions: %w", err)
	}

	return nil
}

// CopyFromRows returns a CopyFromSource interface over the provided rows slice
// making it usable by *Conn.CopyFrom.
func CopyFromTransactions(trs []model.Transaction) pgx.CopyFromSource {
	return &copyFromTransactions{trs: trs, idx: -1}
}

type copyFromTransactions struct {
	trs []model.Transaction
	idx int
}

func (ctr *copyFromTransactions) Next() bool {
	ctr.idx++
	return ctr.idx < len(ctr.trs)
}

func (ctr *copyFromTransactions) Values() ([]any, error) {
	t := ctr.trs[ctr.idx]
	return []any{t.ID, t.Date, t.Amount}, nil
}

func (ctr *copyFromTransactions) Err() error {
	return nil
}
