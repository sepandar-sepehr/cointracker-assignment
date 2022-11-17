package handlers

import (
	"cointracker-assignment/blockchain"
	"cointracker-assignment/models"
	"encoding/json"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type WalletSyncHandler struct {
	db               *gorm.DB
	blockchainClient blockchain.Client
}

func NewWalletSyncHandler(db *gorm.DB, blockchainClient blockchain.Client) *WalletSyncHandler {
	return &WalletSyncHandler{
		db,
		blockchainClient,
	}
}

func (h *WalletSyncHandler) ServeRequest(w http.ResponseWriter, r *http.Request) {
	var wallets []models.Wallet
	result := h.db.Find(&wallets)
	if result.Error != nil {
		http.Error(w, "Could not fetch wallets", http.StatusInternalServerError)
		return
	}

	for _, wallet := range wallets {
		transactionsResponse, err := h.blockchainClient.GetTransactions(wallet.WalletID, wallet.NextOffset)
		if err != nil {
			http.Error(w, "Could not fetch wallets", http.StatusBadGateway)
			return
		}
		h.db.Model(&models.Wallet{}).
			Where("wallet_id = ?", wallet.WalletID).
			Updates(map[string]interface{}{
				"final_balance":          transactionsResponse.FinalBalance,
				"last_synced_time_in_ms": time.Now().UnixMilli(),
				"next_offset":            transactionsResponse.NextOffset,
			})
		for _, transaction := range transactionsResponse.Transactions {
			transaction.WalletID = wallet.WalletID
			h.db.Create(transaction)
		}
	}

	json.NewEncoder(w).Encode("Wallets Synced.")
}
