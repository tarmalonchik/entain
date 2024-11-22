package storage

import (
	"database/sql"
	"errors"
	"time"
)

var (
	ErrUserNotFound         = errors.New("user not found")
	ErrTransactionDuplicate = errors.New("transaction duplicate")
	ErrNotEnoughBalance     = errors.New("not enough balance")
	ErrOverflow             = errors.New("overflow")
)

type User struct {
	ID            string       `db:"id"`
	CurrentAmount int64        `db:"current_amount"`
	CreatedAt     time.Time    `db:"created_at"`
	UpdatedAt     sql.NullTime `db:"updated_at"`
}

type Transaction struct {
	ID         int64     `db:"id"`
	UserID     string    `db:"user_id"`
	ExternalID string    `db:"external_id"`
	Amount     int64     `db:"amount"`
	SourceType string    `db:"source_type"`
	CreatedAt  time.Time `db:"created_at"`
}
