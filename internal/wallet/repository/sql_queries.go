package repository

const (
	createWalletQuery = `INSERT INTO wallet
DEFAULT VALUES
RETURNING id, balance, created_at, updated_at`

	getWalletQuery = `SELECT id, balance, created_at, updated_at FROM wallet WHERE id = $1`

	changeBalanceQuery = `UPDATE wallet
SET balance = balance + $1
WHERE id = $2
RETURNING id, balance, created_at, updated_at`
)
