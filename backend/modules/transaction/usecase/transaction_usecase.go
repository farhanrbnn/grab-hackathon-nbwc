package usecase

import (
	"context"
	"time"

	"grab-hack-for-good/domain"
)

type transactionUsecase struct {
	transactionRepo domain.TransactionRepository
	contextTimeout  time.Duration
}

// NewTransactionUsecase will create new an transactionUsecase object representation of domain.TransactionUsecase interface
func NewTransactionUsecase(cr domain.TransactionRepository, timeout time.Duration) domain.TransactionUsecase {
	return &transactionUsecase{
		transactionRepo: cr,
		contextTimeout:  timeout,
	}
}

func (usecase *transactionUsecase) First(c context.Context, transactionFilter *domain.Transaction) (transaction domain.Transaction, err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()

	transaction, err = usecase.transactionRepo.First(ctx, transactionFilter)
	if err != nil {
		return domain.Transaction{}, err
	}

	return
}

func (usecase *transactionUsecase) FetchByUserId(c context.Context, transactionFilter *domain.Transaction) (res []domain.Transaction, err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()

	res, err = usecase.transactionRepo.FetchByUserId(ctx, transactionFilter)
	if err != nil {
		return make([]domain.Transaction, 0), err
	}

	return
}

func (usecase *transactionUsecase) Fetch(c context.Context) (res []domain.Transaction, err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()

	res, err = usecase.transactionRepo.Fetch(ctx)
	if err != nil {
		return make([]domain.Transaction, 0), err
	}

	return
}

func (usecase *transactionUsecase) Store(c context.Context, transactionStoreRequest *domain.TransactionStoreRequest) (transaction *domain.Transaction, err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()

	transaction = &domain.Transaction{
		OrderId:      transactionStoreRequest.OrderId,
		UserId:       transactionStoreRequest.UserId,
		UserWalletId: transactionStoreRequest.UserWalletId,
		Status:       transactionStoreRequest.Status,
		Amount:       transactionStoreRequest.Amount,
	}

	err = usecase.transactionRepo.Store(ctx, transaction)

	return
}

func (usecase *transactionUsecase) Update(c context.Context, transaction *domain.Transaction, transactionUpdateRequest *domain.TransactionUpdateRequest) (err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()

	transaction.Status = transactionUpdateRequest.Status
	// transaction.Amount = transactionUpdateRequest.Amount

	transaction.UpdatedAt = time.Now().Unix()

	return usecase.transactionRepo.Update(ctx, transaction)
}

func (usecase *transactionUsecase) Delete(c context.Context, transaction *domain.Transaction) (err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()

	err = usecase.transactionRepo.Delete(ctx, transaction)

	return
}
