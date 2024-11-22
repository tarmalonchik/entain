package bootstrap

import (
	"github.com/tarmalonchik/entain/internal/core/handler"
)

// HandlerContainer - contains handlers
type HandlerContainer struct {
	walletHandler *handler.WalletHandler
}

// GetHandlers - handlers provider
func GetHandlers(services *ServiceContainer) *HandlerContainer {
	hCont := &HandlerContainer{}

	hCont.walletHandler = handler.NewWalletHandler(services.transactionSvc)
	return hCont
}
