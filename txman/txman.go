package txman

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

const baseBackoff = 10 * time.Millisecond

type Postgres interface {
	BeginTx(ctx context.Context, opts pgx.TxOptions) (pgx.Tx, error)
}

type Handler func(context.Context) error

type TxManager struct {
	db Postgres
}

type TxConfig struct {
	IsoLevel   pgx.TxIsoLevel
	Retry      uint
	ReadOnly   bool
	MaxBackoff time.Duration
}

func New(db Postgres) *TxManager {
	return &TxManager{db: db}
}

func (tm *TxManager) Do(ctx context.Context, h Handler, opts ...TxOption) error {
	cfg := TxConfig{
		IsoLevel: pgx.ReadCommitted,
		Retry:    5,
	}
	for _, o := range opts {
		o(&cfg)
	}

	for attempt := uint(0); attempt <= cfg.Retry; attempt++ {
		err := tm.execTx(ctx, cfg, h)
		if err == nil {
			return nil
		}

		var pgErr *pgconn.PgError
		if !(errors.As(err, &pgErr) && (pgErr.Code == "40001" || pgErr.Code == "40P01")) {
			return fmt.Errorf("%w", err)
		}

		if attempt < cfg.Retry {
			backoff := baseBackoff << attempt
			if cfg.MaxBackoff > 0 && backoff > cfg.MaxBackoff {
				backoff = cfg.MaxBackoff
			}
			sleep := backoff/2 + time.Duration(rand.Int63n(int64(backoff/2)))
			time.Sleep(sleep)
		}
	}

	return fmt.Errorf("retries exceeded")
}

func (tm *TxManager) execTx(ctx context.Context, cfg TxConfig, h Handler) (err error) {
	if _, ok := ExtractTX(ctx); ok {
		return h(ctx)
	}

	accessMode := pgx.ReadWrite
	if cfg.ReadOnly {
		accessMode = pgx.ReadOnly
	}

	tx, err := tm.db.BeginTx(ctx, pgx.TxOptions{
		IsoLevel:   cfg.IsoLevel,
		AccessMode: accessMode,
	})
	if err != nil {
		return fmt.Errorf("%w: begin tx", err)
	}

	ctx = InjectTX(ctx, tx)

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %v", r)
		}
		if err != nil {
			_ = tx.Rollback(ctx)
			return
		}
		err = tx.Commit(ctx)
	}()

	return h(ctx)
}

type txKey struct{}

func InjectTX(ctx context.Context, tx pgx.Tx) context.Context {
	return context.WithValue(ctx, txKey{}, tx)
}

func ExtractTX(ctx context.Context) (pgx.Tx, bool) {
	tx, ok := ctx.Value(txKey{}).(pgx.Tx)
	return tx, ok
}
