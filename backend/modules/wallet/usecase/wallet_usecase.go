package usecase

import (
	"context"
	"time"

	"grab-hack-for-good/domain"
)

type walletUsecase struct {
	walletRepo     domain.WalletRepository
	contextTimeout time.Duration
}

// NewWalletUsecase will create new an walletUsecase object representation of domain.WalletUsecase interface
func NewWalletUsecase(cr domain.WalletRepository, timeout time.Duration) domain.WalletUsecase {
	return &walletUsecase{
		walletRepo:     cr,
		contextTimeout: timeout,
	}
}

func (usecase *walletUsecase) First(c context.Context, walletFilter *domain.Wallet) (wallet domain.Wallet, err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()

	wallet, err = usecase.walletRepo.First(ctx, walletFilter)
	if err != nil {
		return domain.Wallet{}, err
	}

	return
}

func (usecase *walletUsecase) Fetch(c context.Context) (res []domain.Wallet, err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()

	res, err = usecase.walletRepo.Fetch(ctx)
	if err != nil {
		return make([]domain.Wallet, 0), err
	}

	return
}

func (usecase *walletUsecase) Store(c context.Context, walletRequest *domain.WalletStoreRequest) (err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()

	wallet := domain.Wallet{
		WalletOriginId: walletRequest.WalletOriginId,
		Name:           walletRequest.Name,
	}

	err = usecase.walletRepo.Store(ctx, &wallet)

	return
}

func (usecase *walletUsecase) Update(c context.Context, wallet *domain.Wallet, walletData *domain.WalletUpdateRequest) (err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()

	wallet.Name = walletData.Name
	wallet.UpdatedAt = time.Now().Unix()

	return usecase.walletRepo.Update(ctx, wallet)
}

func (usecase *walletUsecase) Delete(c context.Context, wallet *domain.Wallet) (err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()

	err = usecase.walletRepo.Delete(ctx, wallet)

	return
}
