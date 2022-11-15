package blockchain

import "cointracker-assignment/models"

type Client interface {
	GetTransactions(walletID string, nextOffset *string) (*models.TransactionsResponse, error)
}
