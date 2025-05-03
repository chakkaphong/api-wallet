package handlers

import "api-wallet/internal/api/wallets"

type WalletHandler struct {
	service wallets.Service
}

func NewWalletHandler(service wallets.Service) *WalletHandler {
	return &WalletHandler{
		service: service,
	}
}
