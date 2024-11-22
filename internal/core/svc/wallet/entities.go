package wallet

type GetUserBalanceResponse struct {
	UserId  uint64 `json:"userId"`
	Balance string `json:"balance"`
}

type UpdateBalanceRequest struct {
	Amount        int64
	TransactionId string
	SourceType    TransactionSourceType
	UserID        string
}
