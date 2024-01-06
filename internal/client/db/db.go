package db

import (
	"context"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

// Client клиент для работы с бд
type Client interface {
	DB() DB
	Close() error
}

// Transaction интерфейс для работы с транзакциями
type Transaction interface {
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
}

// Query  оберка для запросов - хранит имя запроса и сам запрос
type Query struct {
	Name     string
	QueryRaw string
}

// SQLExecer группируем NamedExecer и QueryExecer
type SQLExecer interface {
	NamedExecer
	QueryExecer
}

// NamedExecer  интрефейс для работы с именованными запросами (теги структур)
type NamedExecer interface {
	ScanOneContext(ctx context.Context, dest interface{}, q Query, args ...interface{}) error
	ScanAllContext(ctx context.Context, dest interface{}, q Query, args ...interface{}) error
}

// QueryExecer интерфейс для работы с запрсоами
type QueryExecer interface {
	ExecContext(ctx context.Context, q Query, args ...interface{}) (pgconn.CommandTag, error)
	QueryContext(ctx context.Context, q Query, args ...interface{}) (pgx.Rows, error)
	QueryRowContext(ctx context.Context, q Query, args ...interface{}) pgx.Row
}

// Pinger интерфейс проверки соединения с бд
type Pinger interface {
	Ping(ctx context.Context) error
}

// DB
type DB interface {
	SQLExecer
	Transaction
	Pinger
	Close()
}
