package handlers

import (
	"cointracker-assignment/models"
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

type WalletHandler struct {
	db *gorm.DB
}

func NewWalletHandler(db *gorm.DB) *WalletHandler {
	return &WalletHandler{
		db,
	}
}

func (h *WalletHandler) ServeRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodPost:
		walletId := strings.TrimPrefix(r.URL.Path, AddressPath)
		result := h.db.Create(&models.Wallet{
			WalletID:         walletId,
			LastSyncTimeInMS: 0,
		})
		if result.Error != nil {
			http.Error(w, "Wallet ID could not be saved", http.StatusInternalServerError)
			return
		}

		var wallet models.Wallet
		result = h.db.First(&wallet, "wallet_id = ?", walletId)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			http.Error(w, "Failed to query for stored wallet", http.StatusInternalServerError)
		} else {
			json.NewEncoder(w).Encode(fmt.Sprintf("Wallet ID %s stored!", wallet.WalletID))
		}

	case http.MethodDelete:
		walletId := strings.TrimPrefix(r.URL.Path, AddressPath)
		var wallet models.Wallet
		result := h.db.First(&wallet, "wallet_id = ?", walletId)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			http.Error(w, "Wallet ID could not be found", http.StatusBadRequest)
		} else {
			h.db.Where("wallet_id = ?", walletId).Delete(&wallet)
			json.NewEncoder(w).Encode(fmt.Sprintf("Wallet ID %s deleted!", walletId))
		}
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
