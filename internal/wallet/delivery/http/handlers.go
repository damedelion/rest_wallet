package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/damedelion/rest_wallet/internal/entities"
	"github.com/damedelion/rest_wallet/internal/utils"
	"github.com/damedelion/rest_wallet/internal/wallet"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type WalletDelivery struct {
	usecase wallet.WalletUsecase
}

func NewWalletDelivery(usecase wallet.WalletUsecase) wallet.WalletDelivery {
	return &WalletDelivery{usecase}
}

func (w *WalletDelivery) CreateWallet(response http.ResponseWriter, request *http.Request) {
	walletId, err := w.usecase.CreateWallet()
	if err != nil {
		utils.NewErrorResponse(response, http.StatusInternalServerError, fmt.Sprintf("can't create wallet, err: %v", err))
		return
	}

	fmt.Printf("successfully created wallet, id: %s\n", walletId.String())
	utils.NewMessageResponse(response, http.StatusOK, fmt.Sprintf("successfully created wallet, id: %s", walletId.String()))
}

func (w *WalletDelivery) GetWallet(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	idString := vars["id"]
	id, err := uuid.Parse(idString)
	if err != nil {
		utils.NewErrorResponse(response, http.StatusBadRequest, "bad request")
		return
	}

	getDTO := entities.GetWalletDTO{Id: id}

	balance, err := w.usecase.GetWallet(&getDTO)
	if err != nil {
		utils.NewErrorResponse(response, http.StatusNotFound, "not found")
		return
	}

	balanceResult := entities.WalletBalance{Balance: balance}

	response.WriteHeader(http.StatusOK)
	response.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(response).Encode(balanceResult); err != nil {
		fmt.Printf("failed to encode balance: %v\n", err)
		return
	}
}

func (w *WalletDelivery) ChangeBalance(response http.ResponseWriter, request *http.Request) {
	var changeDTO entities.ChangeBalanceDTO
	err := json.NewDecoder(request.Body).Decode(&changeDTO)
	if err != nil {
		utils.NewErrorResponse(response, http.StatusBadRequest, fmt.Sprintf("bad request, err: %v", err))
		return
	}

	err = w.usecase.ChangeBalance(&changeDTO)
	if err != nil {
		utils.NewErrorResponse(response, http.StatusNotFound, fmt.Sprintf("not found, err: %v", err))
		return
	}

	utils.NewMessageResponse(response, http.StatusOK, "successful change balance")
}
