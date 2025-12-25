package postgres

import (
	"context"

	"github.com/jaam8/pkg/txman"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PoolWithTxman struct {
	conn *pgxpool.Pool
}

func NewPoolWithTxman(conn *pgxpool.Pool) *PoolWithTxman {
	return &PoolWithTxman{conn: conn}
}

func (p *PoolWithTxman) Exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error) {
	if tx, ok := txman.ExtractTX(ctx); ok {
		return tx.Exec(ctx, sql, args...)
	}
	return p.conn.Exec(ctx, sql, args...)
}

func (p *PoolWithTxman) Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
	if tx, ok := txman.ExtractTX(ctx); ok {
		return tx.Query(ctx, sql, args...)
	}
	return p.conn.Query(ctx, sql, args...)
}

func (p *PoolWithTxman) QueryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	if tx, ok := txman.ExtractTX(ctx); ok {
		return tx.QueryRow(ctx, sql, args...)
	}
	return p.conn.QueryRow(ctx, sql, args...)
}

func (p *PoolWithTxman) SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults {
	if tx, ok := txman.ExtractTX(ctx); ok {
		return tx.SendBatch(ctx, b)
	}
	return p.conn.SendBatch(ctx, b)
}

func (p *PoolWithTxman) Begin(ctx context.Context) (pgx.Tx, error) {
	return p.conn.Begin(ctx)
}

func (p *PoolWithTxman) BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error) {
	return p.conn.BeginTx(ctx, txOptions)
}
