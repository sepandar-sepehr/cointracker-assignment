package handlers

import (
	"encoding/json"
	"net/http"
)

func WalletSyncHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Wallets Synced.")
}
