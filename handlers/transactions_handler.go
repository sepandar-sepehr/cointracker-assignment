package handlers

import (
	"encoding/json"
	"net/http"
)

func TransactionsHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Sample Transaction.")
}
