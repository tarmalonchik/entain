// Package wallet
// nolint
//
//go:generate go-enum -f=$GOFILE --nocase --sqlnullint
package wallet

// TransactionStateType ENUM(win, lose)
type TransactionStateType string

// TransactionSourceType ENUM(game, server, payment)
type TransactionSourceType string
