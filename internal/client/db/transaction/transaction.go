package transaction

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/kom1ssar/grpc-auth/internal/client/db"
	"github.com/kom1ssar/grpc-auth/internal/client/db/pg"
	"github.com/pkg/errors"
)

type txManager struct {
	db db.Transaction
}

func NewTransactionManager(db db.Transaction) db.TxManager {
	return &txManager{db: db}
}

func (m *txManager) transaction(ctx context.Context, opts pgx.TxOptions, fn db.HandlerTx) (err error) {
	tx, ok := ctx.Value(pg.TxKey).(pgx.Tx)
	if ok {
		return fn(ctx)
	}

	tx, err = m.db.BeginTx(ctx, opts)
	if err != nil {
		return errors.Wrap(err, "begin transaction")
	}

	ctx = pg.MakeContextTx(ctx, tx)

	defer func() {
		if r := recover(); r != nil {
			err = errors.Errorf("panic recovered: %v", r)
		}

		if err != nil {
			if errRollback := tx.Rollback(ctx); errRollback != nil {
				err = errors.Wrapf(err, "err_rollback: %v", errRollback)
			}
			return
		}

		if err == nil {
			err = tx.Commit(ctx)
			if err != nil {
				err = errors.Wrapf(err, "tx commit error")
			}
		}
	}()

	if err = fn(ctx); err != nil {
		err = errors.Wrap(err, "failed executing code inside transaction")

	}

	return err
}

func (m *txManager) ReadCommitted(ctx context.Context, f db.HandlerTx) error {
	txOptions := pgx.TxOptions{IsoLevel: pgx.ReadCommitted}

	return m.transaction(ctx, txOptions, f)
}
