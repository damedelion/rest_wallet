package wallet

import (
	"net/http"
)

type WalletDelivery interface {
	CreateWallet(response http.ResponseWriter, request *http.Request)
	GetWallet(response http.ResponseWriter, request *http.Request)
	ChangeBalance(response http.ResponseWriter, request *http.Request)
}
