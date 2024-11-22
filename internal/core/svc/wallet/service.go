package wallet

import (
	"context"
	"errors"
	"strconv"

	storagePkg "github.com/tarmalonchik/entain/internal/pkg/storage"
	"github.com/tarmalonchik/entain/internal/pkg/tools"
)

type storage interface {
	GetUser(ctx context.Context, id string) (user storagePkg.User, err error)
	UpdateUserBalance(ctx context.Context, updateBalanceRequest storagePkg.Transaction) error
}

type Service struct {
	storage storage
}

func NewService(storage storage) *Service {
	return &Service{
		storage: storage,
	}
}

func (s *Service) GetUserBalance(ctx context.Context, userID string) (GetUserBalanceResponse, error) {
	user, err := s.storage.GetUser(ctx, userID)
	if err != nil {
		return GetUserBalanceResponse{}, err
	}

	userIDUint64, err := strconv.ParseUint(user.ID, 10, 64)
	if err != nil {
		return GetUserBalanceResponse{}, errors.New("invalid data in database")
	}

	return GetUserBalanceResponse{
		UserId:  userIDUint64,
		Balance: tools.CentsPrettyPrinted(user.CurrentAmount),
	}, nil
}

func (s *Service) UpdateBalance(ctx context.Context, req UpdateBalanceRequest) error {
	err := s.storage.UpdateUserBalance(ctx, storagePkg.Transaction{
		UserID:     req.UserID,
		ExternalID: req.TransactionId,
		Amount:     req.Amount,
		SourceType: req.SourceType.String(),
	})
	return err
}
