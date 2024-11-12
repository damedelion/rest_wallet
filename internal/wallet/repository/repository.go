package repository

import (
	"database/sql"
	"fmt"

	"github.com/damedelion/rest_wallet/internal/entities"
	"github.com/damedelion/rest_wallet/internal/wallet"
)

type WalletRepository struct {
	db *sql.DB
}

func NewWalletRepository(db *sql.DB) wallet.WalletRepository {
	return &WalletRepository{db}
}

func (w *WalletRepository) CreateWallet() (*entities.Wallet, error) {
	newWallet := &entities.Wallet{}

	row := w.db.QueryRow(createWalletQuery)

	if err := row.Scan(
		&newWallet.Id,
		&newWallet.Balance,
		&newWallet.CreatedAt,
		&newWallet.UpdatedAt,
	); err != nil {
		return nil, fmt.Errorf("failed to scan from createWallet query")
	}

	return newWallet, nil
}

func (w *WalletRepository) GetWallet(in *entities.GetWalletDTO) (*entities.Wallet, error) {
	newWallet := &entities.Wallet{}

	row := w.db.QueryRow(getWalletQuery, in.Id)

	if err := row.Scan(
		&newWallet.Id,
		&newWallet.Balance,
		&newWallet.CreatedAt,
		&newWallet.UpdatedAt,
	); err != nil {
		return nil, fmt.Errorf("failed to scan from getWallet query")
	}

	return newWallet, nil
}

func (w *WalletRepository) ChangeBalance(in *entities.ChangeBalanceDTO) error {
	newWallet := &entities.Wallet{}

	row := w.db.QueryRow(changeBalanceQuery, in.Amount, in.Id)

	if err := row.Scan(
		&newWallet.Id,
		&newWallet.Balance,
		&newWallet.CreatedAt,
		&newWallet.UpdatedAt,
	); err != nil {
		return fmt.Errorf("failed to scan from changeBalance query")
	}

	return nil
}
