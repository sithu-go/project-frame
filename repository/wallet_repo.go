package repository

import (
	"h-pay/ds"

	"gorm.io/gorm"
)

type walletRepository struct {
	DB *gorm.DB
}

func newWalletRepository(ds *ds.DataSource) *walletRepository {
	return &walletRepository{
		DB: ds.DB,
	}
}
