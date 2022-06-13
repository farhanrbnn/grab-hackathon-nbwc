package usecase

import (
	"context"
	"time"

	"grab-hack-for-good/domain"
)

type merchantUsecase struct {
	merchantRepo   domain.MerchantRepository
	contextTimeout time.Duration
}

// NewMerchantUsecase will create new an merchantUsecase object representation of domain.MerchantUsecase interface
func NewMerchantUsecase(cr domain.MerchantRepository, timeout time.Duration) domain.MerchantUsecase {
	return &merchantUsecase{
		merchantRepo:   cr,
		contextTimeout: timeout,
	}
}

func (usecase *merchantUsecase) First(c context.Context, merchantFilter *domain.Merchant) (merchant domain.Merchant, err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()

	merchant, err = usecase.merchantRepo.First(ctx, merchantFilter)
	if err != nil {
		return domain.Merchant{}, err
	}

	return
}

func (usecase *merchantUsecase) Fetch(c context.Context) (res []domain.Merchant, err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()

	res, err = usecase.merchantRepo.Fetch(ctx)
	if err != nil {
		return make([]domain.Merchant, 0), err
	}

	return
}

func (usecase *merchantUsecase) Store(c context.Context, merchantStoreRequest *domain.MerchantStoreRequest) (err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()

	merchant := domain.Merchant{
		Name:        merchantStoreRequest.Name,
		Coordinate:  merchantStoreRequest.Coordinate,
		Rating:      merchantStoreRequest.Rating,
		Address:     merchantStoreRequest.Address,
		Thumbnail:   merchantStoreRequest.Thumbnail,
		IsAvailable: merchantStoreRequest.IsAvailable,
	}

	err = usecase.merchantRepo.Store(ctx, &merchant)

	return
}

func (usecase *merchantUsecase) Update(c context.Context, merchant *domain.Merchant, merchantUpdateRequest *domain.MerchantUpdateRequest) (err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()

	merchant.Name = merchantUpdateRequest.Name
	merchant.Phone = merchantUpdateRequest.Phone
	merchant.Email = merchantUpdateRequest.Email
	merchant.Coordinate = merchantUpdateRequest.Coordinate
	merchant.Rating = merchantUpdateRequest.Rating
	merchant.Address = merchantUpdateRequest.Address
	merchant.Thumbnail = merchantUpdateRequest.Thumbnail
	merchant.IsAvailable = merchantUpdateRequest.IsAvailable

	merchant.UpdatedAt = time.Now().Unix()

	return usecase.merchantRepo.Update(ctx, merchant)
}

func (usecase *merchantUsecase) Delete(c context.Context, merchant *domain.Merchant) (err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()

	err = usecase.merchantRepo.Delete(ctx, merchant)

	return
}
