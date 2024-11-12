package entities

import "github.com/google/uuid"

type GetWalletDTO struct {
	Id uuid.UUID `json:"id"`
}

type ChangeBalanceDTO struct {
	Id            uuid.UUID `json:"id"`
	OperationType string    `json:"operation_type"`
	Amount        int       `json:"amount"`
}
