package handlers

type Transaction struct {
	Hash string `json:"hash"`
	Time int    `json:"time"`
}

type GetTransactionsApiResponse struct {
	WalletID     string        `json:"wallet_id"`
	FinalBalance int           `json:"final_balance"`
	Transactions []Transaction `json:"transactions"`
}
