package usecase

import (
	"context"
	"time"

	"grab-hack-for-good/domain"
)

type walletOriginUsecase struct {
	walletOriginRepo domain.WalletOriginRepository
	contextTimeout   time.Duration
}

// NewWalletOriginUsecase will create new an walletOriginUsecase object representation of domain.WalletOriginUsecase interface
func NewWalletOriginUsecase(cr domain.WalletOriginRepository, timeout time.Duration) domain.WalletOriginUsecase {
	return &walletOriginUsecase{
		walletOriginRepo: cr,
		contextTimeout:   timeout,
	}
}

func (usecase *walletOriginUsecase) First(c context.Context, walletOriginFilter *domain.WalletOrigin) (walletOrigin domain.WalletOrigin, err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()

	walletOrigin, err = usecase.walletOriginRepo.First(ctx, walletOriginFilter)
	if err != nil {
		return domain.WalletOrigin{}, err
	}

	return
}

func (usecase *walletOriginUsecase) Fetch(c context.Context) (res []domain.WalletOrigin, err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()

	res, err = usecase.walletOriginRepo.Fetch(ctx)
	if err != nil {
		return make([]domain.WalletOrigin, 0), err
	}

	return
}

func (usecase *walletOriginUsecase) Store(c context.Context, walletOriginRequest *domain.WalletOriginRequest) (err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()

	walletOrigin := domain.WalletOrigin{
		Name:    walletOriginRequest.Name,
		Company: walletOriginRequest.Company,
	}

	err = usecase.walletOriginRepo.Store(ctx, &walletOrigin)

	return
}

func (usecase *walletOriginUsecase) Update(c context.Context, walletOrigin *domain.WalletOrigin, walletOriginData *domain.WalletOriginRequest) (err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()

	walletOrigin.Name = walletOriginData.Name
	walletOrigin.Company = walletOriginData.Company
	walletOrigin.UpdatedAt = time.Now().Unix()

	return usecase.walletOriginRepo.Update(ctx, walletOrigin)
}

func (usecase *walletOriginUsecase) Delete(c context.Context, walletOrigin *domain.WalletOrigin) (err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()

	err = usecase.walletOriginRepo.Delete(ctx, walletOrigin)

	return
}
