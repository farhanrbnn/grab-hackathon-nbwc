package usecase

import (
	"context"
	"time"

	"grab-hack-for-good/domain"
)

type userWalletUsecase struct {
	userWalletRepo domain.UserWalletRepository
	contextTimeout time.Duration
}

// NewUserWalletUsecase will create new an userWalletUsecase object representation of domain.UserWalletUsecase interface
func NewUserWalletUsecase(cr domain.UserWalletRepository, timeout time.Duration) domain.UserWalletUsecase {
	return &userWalletUsecase{
		userWalletRepo: cr,
		contextTimeout: timeout,
	}
}

func (usecase *userWalletUsecase) First(c context.Context, userWalletFilter *domain.UserWallet) (userWallet domain.UserWallet, err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()

	userWallet, err = usecase.userWalletRepo.First(ctx, userWalletFilter)
	if err != nil {
		return domain.UserWallet{}, err
	}

	return
}

func (usecase *userWalletUsecase) FetchByUserId(c context.Context, userWalletFilter *domain.UserWallet) (res []domain.UserWallet, err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()

	res, err = usecase.userWalletRepo.FetchByUserId(ctx, userWalletFilter)
	if err != nil {
		return make([]domain.UserWallet, 0), err
	}

	return
}

func (usecase *userWalletUsecase) Fetch(c context.Context) (res []domain.UserWallet, err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()

	res, err = usecase.userWalletRepo.Fetch(ctx)
	if err != nil {
		return make([]domain.UserWallet, 0), err
	}

	return
}

func (usecase *userWalletUsecase) Store(c context.Context, userWalletStoreRequest *domain.UserWalletStoreRequest) (err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()

	userWallet := domain.UserWallet{
		UserId:          userWalletStoreRequest.UserId,
		WalletId:        userWalletStoreRequest.WalletId,
		EffectiveAmount: userWalletStoreRequest.EffectiveAmount,
		OnHoldAmount:    userWalletStoreRequest.OnHoldAmount,
	}

	err = usecase.userWalletRepo.Store(ctx, &userWallet)

	return
}

func (usecase *userWalletUsecase) Pay(c context.Context, userWallet *domain.UserWallet, userWalletTransactionRequest *domain.UserWalletTransactionRequest) (err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()

	userWallet.EffectiveAmount = userWallet.EffectiveAmount - userWalletTransactionRequest.Amount
	userWallet.OnHoldAmount = userWallet.OnHoldAmount + userWalletTransactionRequest.Amount
	userWallet.UpdatedAt = time.Now().Unix()

	return usecase.userWalletRepo.Update(ctx, userWallet)
}

func (usecase *userWalletUsecase) Refund(c context.Context, userWallet *domain.UserWallet, userWalletTransactionRequest *domain.UserWalletTransactionRequest) (err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()

	userWallet.EffectiveAmount = userWallet.EffectiveAmount + userWalletTransactionRequest.Amount
	userWallet.OnHoldAmount = userWallet.OnHoldAmount - userWalletTransactionRequest.Amount
	userWallet.UpdatedAt = time.Now().Unix()

	return usecase.userWalletRepo.Update(ctx, userWallet)
}

func (usecase *userWalletUsecase) Commit(c context.Context, userWallet *domain.UserWallet, userWalletTransactionRequest *domain.UserWalletTransactionRequest) (err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()

	userWallet.OnHoldAmount = userWallet.OnHoldAmount - userWalletTransactionRequest.Amount
	userWallet.UpdatedAt = time.Now().Unix()

	return usecase.userWalletRepo.Update(ctx, userWallet)
}

func (usecase *userWalletUsecase) Update(c context.Context, userWallet *domain.UserWallet, userWalletUpdateRequest *domain.UserWalletUpdateRequest) (err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()

	userWallet.EffectiveAmount = userWalletUpdateRequest.EffectiveAmount
	userWallet.OnHoldAmount = userWalletUpdateRequest.OnHoldAmount
	userWallet.UpdatedAt = time.Now().Unix()

	return usecase.userWalletRepo.Update(ctx, userWallet)
}

func (usecase *userWalletUsecase) Delete(c context.Context, userWallet *domain.UserWallet) (err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()

	err = usecase.userWalletRepo.Delete(ctx, userWallet)

	return
}
