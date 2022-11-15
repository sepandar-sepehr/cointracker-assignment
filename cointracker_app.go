package main

import (
	"cointracker-assignment/handlers"
	"context"
	"html/template"
	"log"
	"net/http"
	"time"
)

func renderPage(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("web/wallet.html")
	err := t.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	// Setting timeout 30 seconds for all APIs
	ctx := context.Background()
	ctx, cancelCtx := context.WithTimeout(ctx, 30*time.Second)
	defer cancelCtx()

	http.HandleFunc("/wallet", renderPage)
	http.HandleFunc(handlers.WalletSyncPath, handlers.WalletSyncHandler)
	http.HandleFunc(handlers.AddressPath, handlers.WalletHandler)
	http.HandleFunc(handlers.TransactionsPath, handlers.TransactionsHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
