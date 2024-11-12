package entities

import (
	"time"

	"github.com/google/uuid"
)

type Wallet struct {
	Id        uuid.UUID `json:"id"`
	Balance   int       `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type WalletBalance struct {
	Balance int `json:"balance"`
}
