package storage

import (
	"context"
	"fmt"
	"math"
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/tarmalonchik/entain/internal/pkg/postgresql"
)

type Model struct {
	db *postgresql.SQLXClient
}

func NewModel(db *postgresql.SQLXClient) *Model {
	return &Model{db: db}
}

func (m *Model) GetUser(ctx context.Context, id string) (user User, err error) {
	const query = `select * from "user" where id = $1`

	if err = m.db.GetContext(ctx, &user, query, id); err != nil {
		if postgresql.IsNotFound(err) {
			return User{}, ErrUserNotFound
		}
		return User{}, fmt.Errorf("getting user: %w", err)
	}
	return user, nil
}

func (m *Model) UpdateUserBalance(ctx context.Context, updateBalanceRequest Transaction) error {
	return m.handlePGTransaction(func(tx *sqlx.Tx) error {
		selectUserQuery := `select * from "user" where id = $1 for update`
		var user User

		if err := tx.GetContext(ctx, &user, selectUserQuery, updateBalanceRequest.UserID); err != nil {
			if postgresql.IsNotFound(err) {
				return ErrUserNotFound
			}
			return err
		}

		newAmount, err := Add64(user.CurrentAmount, updateBalanceRequest.Amount)
		if err != nil {
			return err
		}

		if newAmount < 0 {
			return ErrNotEnoughBalance
		}

		createTransactionQuery := `
			insert into transaction (user_id, external_id, source_type, amount, created_at) values 
			($1,$2,$3,$4,$5) returning *`
		var transaction Transaction

		err = tx.GetContext(
			ctx,
			&transaction,
			createTransactionQuery,
			updateBalanceRequest.UserID,
			updateBalanceRequest.ExternalID,
			updateBalanceRequest.SourceType,
			updateBalanceRequest.Amount,
			time.Now().UTC(),
		)
		if err != nil {
			if postgresql.ViolatesUniqueConstraint(err) {
				return ErrTransactionDuplicate
			}
			return err
		}

		updateUserBalanceQuery := `update "user" set current_amount = $1, updated_at = $3 where id = $2`
		_, err = tx.ExecContext(ctx, updateUserBalanceQuery, newAmount, transaction.UserID, time.Now().UTC())
		if err != nil {
			return err
		}
		return nil
	})
}

func Add64(balance, addition int64) (int64, error) {
	if balance == 0 || addition == 0 {
		return balance + addition, nil
	}
	if addition > 0 {
		if balance > math.MaxInt64-addition {
			return 0, ErrOverflow
		}
	} else {
		if balance < math.MaxInt64-addition {
			return 0, ErrOverflow
		}
	}
	return balance + addition, nil
}
