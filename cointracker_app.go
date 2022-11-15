package main

import (
	"cointracker-assignment/blockchain"
	"cointracker-assignment/handlers"
	"cointracker-assignment/models"
	"context"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
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

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.Wallet{}, &models.Transaction{})

	http.HandleFunc("/wallet", renderPage)

	walletSyncHandler := handlers.NewWalletSyncHandler(db, blockchain.NewBcComClient())
	http.HandleFunc(handlers.WalletSyncPath, walletSyncHandler.ServeRequest)

	walletHandler := handlers.NewWalletHandler(db)
	http.HandleFunc(handlers.AddressPath, walletHandler.ServeRequest)

	transactionHandler := handlers.NewTransactionHandler(db)
	http.HandleFunc(handlers.TransactionsPath, transactionHandler.ServeRequest)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
