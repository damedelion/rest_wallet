package wallet

import "github.com/damedelion/rest_wallet/internal/entities"

type WalletRepository interface {
	CreateWallet() (*entities.Wallet, error)
	GetWallet(in *entities.GetWalletDTO) (*entities.Wallet, error)
	ChangeBalance(in *entities.ChangeBalanceDTO) error
}
