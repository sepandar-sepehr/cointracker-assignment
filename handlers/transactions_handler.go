package handlers

import (
	"encoding/json"
	"gorm.io/gorm"
	"net/http"
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
	json.NewEncoder(w).Encode("Sample Transaction.")
}
