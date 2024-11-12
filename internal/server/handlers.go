package server

import (
	"github.com/damedelion/rest_wallet/internal/wallet/delivery/http"
	"github.com/damedelion/rest_wallet/internal/wallet/repository"
	"github.com/damedelion/rest_wallet/internal/wallet/usecase"
)

func (s *Server) HandlerRegister() {
	walletRepo := repository.NewWalletRepository(s.db)
	walletUsecase := usecase.NewWalletUsecase(walletRepo)
	walletHandlers := http.NewWalletDelivery(walletUsecase)

	s.mux.HandleFunc("/api/v1/wallet/create", walletHandlers.CreateWallet).Methods("GET")
	s.mux.HandleFunc("/api/v1/wallets/{id}", walletHandlers.GetWallet).Methods("GET")
	s.mux.HandleFunc("/api/v1/wallet", walletHandlers.ChangeBalance).Methods("POST")
}
