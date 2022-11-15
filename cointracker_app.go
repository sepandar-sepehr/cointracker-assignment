package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

const (
	WalletPagePath   = "/wallet"
	WalletSyncPath   = "/wallet/sync"
	AddressPath      = "/wallet/bitcoin/address/"
	TransactionsPath = "/wallet/bitcoin/transactions/"
)

func walletHandler(w http.ResponseWriter, r *http.Request) {
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

func walletSyncHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Wallets Synced.")
}

func transactionsHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Sample Transaction.")
}

func handler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("web/wallet.html")
	err := t.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc(WalletPagePath, handler)
	http.HandleFunc(WalletSyncPath, walletSyncHandler)
	http.HandleFunc(AddressPath, walletHandler)
	http.HandleFunc(TransactionsPath, transactionsHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
