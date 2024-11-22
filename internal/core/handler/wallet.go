package handler

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"github.com/tarmalonchik/entain/internal/core/svc/wallet"
	"github.com/tarmalonchik/entain/internal/pkg/response"
	"github.com/tarmalonchik/entain/internal/pkg/storage"
	"github.com/tarmalonchik/entain/internal/pkg/tools"
)

const sourceTypeHeaderKey = "Source-Type"

func NewWalletHandler(walletSvc *wallet.Service) *WalletHandler {
	return &WalletHandler{
		walletSvc: walletSvc,
	}
}

type WalletHandler struct {
	walletSvc *wallet.Service
}

func (s *WalletHandler) GetUserBalance(w http.ResponseWriter, r *http.Request) {
	userIDString, ok := mux.Vars(r)["userId"]
	if !ok {
		response.JSON400txt(w, "should specify userId")
		return
	}

	_, err := strconv.ParseUint(userIDString, 10, 64)
	if err != nil {
		response.JSON400txt(w, "invalid userId")
		return
	}

	userBalance, err := s.walletSvc.GetUserBalance(r.Context(), userIDString)
	if err != nil {
		logrus.New().Errorf("GetUserBalance svc: %v", err)
		if errors.Is(err, storage.ErrUserNotFound) {
			response.JSON404txt(w, "user not found")
			return
		}
		response.JSON500txt(w, "service error")
		return
	}

	response.RenderJSON(w, 200, userBalance)
}

type UpdateBalanceRequest struct {
	State         string `json:"state"`
	Amount        string `json:"amount"`
	TransactionId string `json:"transactionId"`
}

func (u *UpdateBalanceRequest) ConvertToWalletRequest(sourceType string) (out wallet.UpdateBalanceRequest, valid bool) {
	if u.TransactionId == "" || u.Amount == "" {
		return wallet.UpdateBalanceRequest{}, false
	}

	state, err := wallet.ParseTransactionStateType(u.State)
	if err != nil {
		return wallet.UpdateBalanceRequest{}, false
	}

	out.SourceType, err = wallet.ParseTransactionSourceType(sourceType)
	if err != nil {
		return wallet.UpdateBalanceRequest{}, false
	}

	amount, err := tools.ConvertNonNegativeFloatToCents(u.Amount)
	if err != nil {
		return wallet.UpdateBalanceRequest{}, false
	}

	switch state {
	case wallet.TransactionStateTypeWin:
		out.Amount = amount
	case wallet.TransactionStateTypeLose:
		out.Amount = -amount
	}

	out.TransactionId = u.TransactionId
	return out, true
}

func (s *WalletHandler) UpdateBalance(w http.ResponseWriter, r *http.Request) {
	var (
		req  = UpdateBalanceRequest{}
		err  error
		body []byte
	)

	userIDString, ok := mux.Vars(r)["userId"]
	if !ok {
		response.JSON400txt(w, "should specify userId")
		return
	}

	_, err = strconv.ParseUint(userIDString, 10, 64)
	if err != nil {
		response.JSON400txt(w, "invalid userId")
		return
	}

	if body, err = io.ReadAll(r.Body); err != nil {
		response.JSON400(w)
		return
	}
	defer func() { _ = r.Body.Close() }()

	if err = json.Unmarshal(body, &req); err != nil {
		logrus.New().Errorf("UpdateBalance json unmarshal: %v", err)
		response.JSON500(w)
		return
	}

	svcReq, valid := req.ConvertToWalletRequest(r.Header.Get(sourceTypeHeaderKey))
	if !valid {
		response.JSON400txt(w, "invalid request")
		return
	}
	svcReq.UserID = userIDString

	if err = s.walletSvc.UpdateBalance(r.Context(), svcReq); err != nil {
		if errors.Is(err, storage.ErrTransactionDuplicate) {
			response.JSON409txt(w, "transaction conflict")
			return
		}
		if errors.Is(err, storage.ErrUserNotFound) {
			response.JSON404txt(w, "user not found")
			return
		}
		if errors.Is(err, storage.ErrOverflow) {
			response.JSON404txt(w, "user have max balance")
			return
		}
		if errors.Is(err, storage.ErrNotEnoughBalance) {
			response.JSON404txt(w, "not enough balance")
			return
		}
		logrus.New().Errorf("UpdateBalance svc: %v", err)
		response.JSON500(w)
		return
	}
	response.JSON200(w)
}
