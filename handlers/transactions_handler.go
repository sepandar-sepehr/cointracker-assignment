package handlers

import (
	"cointracker-assignment/models"
	"encoding/json"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

type TransactionHandler struct {
	db *gorm.DB
}

func NewTransactionHandler(db *gorm.DB) *TransactionHandler {
	return &TransactionHandler{
		db,
	}
}

func (h *TransactionHandler) ServeRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	walletId := strings.TrimPrefix(r.URL.Path, TransactionsPath)

	var transactions []models.Transaction
	result := h.db.Where(&models.Transaction{WalletID: walletId}).Find(&transactions)
	if result.Error != nil {
		http.Error(w, "Could not fetch transactions", http.StatusInternalServerError)
		return
	}

	var wallet models.Wallet
	result = h.db.First(&wallet, "wallet_id = ?", walletId)
	if result.Error != nil {
		http.Error(w, "Could not fetch transactions", http.StatusInternalServerError)
		return
	}

	response := convertResponse(walletId, wallet, transactions)

	json.NewEncoder(w).Encode(response)
}

func convertResponse(walletId string, wallet models.Wallet, transactions []models.Transaction) GetTransactionsApiResponse {
	responseTransactions := make([]Transaction, len(transactions))
	for i, transaction := range transactions {
		responseTransactions[i] = Transaction{
			Hash: transaction.Hash,
			Time: transaction.Time,
		}
	}

	return GetTransactionsApiResponse{
		WalletID:     walletId,
		FinalBalance: wallet.FinalBalance,
		Transactions: responseTransactions,
	}
}
