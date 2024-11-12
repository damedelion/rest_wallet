package wallet

import (
	"github.com/damedelion/rest_wallet/internal/entities"
	"github.com/google/uuid"
)

type WalletUsecase interface {
	CreateWallet() (uuid.UUID, error)
	GetWallet(*entities.GetWalletDTO) (int, error)
	ChangeBalance(*entities.ChangeBalanceDTO) error
}
