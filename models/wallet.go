package models

type Wallet struct {
	WalletID           string `gorm:"primaryKey"`
	LastSyncedTimeInMS int64
	NextOffset         *string
	FinalBalance       int
}
