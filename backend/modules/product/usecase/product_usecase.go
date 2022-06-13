package usecase

import (
	"context"
	"time"

	"grab-hack-for-good/domain"
)

type productUsecase struct {
	productRepo    domain.ProductRepository
	contextTimeout time.Duration
}

// NewProductUsecase will create new an productUsecase object representation of domain.ProductUsecase interface
func NewProductUsecase(cr domain.ProductRepository, timeout time.Duration) domain.ProductUsecase {
	return &productUsecase{
		productRepo:    cr,
		contextTimeout: timeout,
	}
}

func (usecase *productUsecase) First(c context.Context, productFilter *domain.Product) (product domain.Product, err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()

	product, err = usecase.productRepo.First(ctx, productFilter)
	if err != nil {
		return domain.Product{}, err
	}

	return
}

func (usecase *productUsecase) FetchByMerchantId(c context.Context, productFilter *domain.Product) (res []domain.Product, err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()

	res, err = usecase.productRepo.FetchByMerchantId(ctx, productFilter)
	if err != nil {
		return make([]domain.Product, 0), err
	}

	return
}

func (usecase *productUsecase) Fetch(c context.Context) (res []domain.Product, err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()

	res, err = usecase.productRepo.Fetch(ctx)
	if err != nil {
		return make([]domain.Product, 0), err
	}

	return
}

func (usecase *productUsecase) Store(c context.Context, productStoreRequest *domain.ProductStoreRequest) (err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()

	product := domain.Product{
		MerchantId:  productStoreRequest.MerchantId,
		Name:        productStoreRequest.Name,
		Description: productStoreRequest.Description,
		Images:      productStoreRequest.Images,
		Price:       productStoreRequest.Price,
		Dimensions:  productStoreRequest.Dimensions,
	}

	err = usecase.productRepo.Store(ctx, &product)

	return
}

func (usecase *productUsecase) Update(c context.Context, product *domain.Product, productUpdateRequest *domain.ProductUpdateRequest) (err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()

	product.Name = productUpdateRequest.Name
	product.Description = productUpdateRequest.Description
	product.Images = productUpdateRequest.Images
	product.Price = productUpdateRequest.Price
	product.Dimensions = productUpdateRequest.Dimensions

	product.UpdatedAt = time.Now().Unix()

	return usecase.productRepo.Update(ctx, product)
}

func (usecase *productUsecase) Delete(c context.Context, product *domain.Product) (err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()

	err = usecase.productRepo.Delete(ctx, product)

	return
}
