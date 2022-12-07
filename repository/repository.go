package repository

import (
	"h-pay/ds"
	"h-pay/service"
)

var (
	SVC *service.Service
)

type Repository struct {
	DS     *ds.DataSource
	Bank   *bankRepository
	Admin  *adminRepository
	User   *userRepository
	Wallet *walletRepository
}

func NewRepository(ds *ds.DataSource, svc *service.Service) *Repository {
	bankRepo := newBankRepository(ds)
	adminRepo := newAdminRepository(ds)
	userRepo := newUserRepository(ds)
	walletRepo := newWalletRepository(ds)
	return &Repository{
		DS:     ds,
		Bank:   bankRepo,
		Admin:  adminRepo,
		User:   userRepo,
		Wallet: walletRepo,
	}
}
