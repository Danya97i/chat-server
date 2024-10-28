package db

import (
	"context"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

type Handler func(ctx context.Context) error

type TxManager interface {
	ReadCommited(ctx context.Context, f Handler) error
}

type Transactor interface {
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
}

type Query struct {
	Name     string
	RawQuery string
}

type Client interface {
	DB() DB
	Close() error
}

type SQLExecer interface {
	QueryExecer
	NamedExecer
}

type QueryExecer interface {
	ExecContext(ctx context.Context, query Query, args ...interface{}) (pgconn.CommandTag, error)
	QueryRowContext(ctx context.Context, Query Query, args ...interface{}) pgx.Row
}

type NamedExecer interface {
	ScanOneContext(ctx context.Context, dest interface{}, q Query, args ...interface{}) error
	ScanAllContext(ctx context.Context, dest interface{}, q Query, args ...interface{}) error
}

type DB interface {
	SQLExecer
	NamedExecer
	Transactor
	Close()
}
