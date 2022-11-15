package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func WalletHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodPost:
		walletId := strings.TrimPrefix(r.URL.Path, AddressPath)
		json.NewEncoder(w).Encode(fmt.Sprintf("Wallet ID %s stored!", walletId))
	case http.MethodDelete:
		walletId := strings.TrimPrefix(r.URL.Path, AddressPath)
		json.NewEncoder(w).Encode(fmt.Sprintf("Wallet ID %s not found!", walletId))
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
