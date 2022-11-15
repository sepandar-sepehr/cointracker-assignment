package handlers

import (
	"encoding/json"
	"gorm.io/gorm"
	"net/http"
)

type WalletSyncHandler struct {
	db *gorm.DB
}

func NewWalletSyncHandler(db *gorm.DB) *WalletSyncHandler {
	return &WalletSyncHandler{
		db,
	}
}

func (h *WalletSyncHandler) ServeRequest(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Wallets Synced.")
}
