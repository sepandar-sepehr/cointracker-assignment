package models

type Transaction struct {
	WalletID string `gorm:"index"`
	Hash     string `json:"hash" gorm:"primaryKey"`
	Time     int    `json:"time"`
}

type TransactionsResponse struct {
	FinalBalance int           `json:"final_balance"`
	Transactions []Transaction `json:"txs"`
	NextOffset   string
}
