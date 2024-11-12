package usecase

import (
	"fmt"
	"strings"

	"github.com/damedelion/rest_wallet/internal/entities"
	"github.com/damedelion/rest_wallet/internal/wallet"
	"github.com/google/uuid"
)

type WalletUsecase struct {
	repo wallet.WalletRepository
}

func NewWalletUsecase(repo wallet.WalletRepository) wallet.WalletUsecase {
	return &WalletUsecase{repo}
}

func (w *WalletUsecase) CreateWallet() (uuid.UUID, error) {
	wallet, err := w.repo.CreateWallet()
	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to create wallet, err: %v", err)
	}
	return wallet.Id, nil
}

func (w *WalletUsecase) GetWallet(in *entities.GetWalletDTO) (int, error) {
	wallet, err := w.repo.GetWallet(in)
	if err != nil {
		return 0, err
	}
	return wallet.Balance, nil
}

func (w *WalletUsecase) ChangeBalance(in *entities.ChangeBalanceDTO) error {
	getWalletIn := entities.GetWalletDTO{Id: in.Id}
	wallet, err := w.repo.GetWallet(&getWalletIn)
	if err != nil {
		return err
	}

	switch strings.ToLower(in.OperationType) {
	case "deposit":

	case "withdraw":
		in.Amount = -in.Amount
	default:
		return fmt.Errorf("invalid operation type")
	}

	if wallet.Balance+in.Amount < 0 {
		return fmt.Errorf("don't have enough money on balance")
	}

	return w.repo.ChangeBalance(in)
}
