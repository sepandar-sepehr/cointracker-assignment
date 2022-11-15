package models

import "gorm.io/gorm"

type Wallet struct {
	gorm.Model
	WalletID         string `gorm:"primaryKey"`
	LastSyncTimeInMS int
}
