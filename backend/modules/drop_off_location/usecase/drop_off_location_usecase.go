package usecase

import (
	"context"
	"time"

	"grab-hack-for-good/domain"
)

type dropOffLocationUsecase struct {
	dropOffLocationRepo domain.DropOffLocationRepository
	contextTimeout      time.Duration
}

// NewDropOffLocationUsecase will create new an dropOffLocationUsecase object representation of domain.DropOffLocationUsecase interface
func NewDropOffLocationUsecase(cr domain.DropOffLocationRepository, timeout time.Duration) domain.DropOffLocationUsecase {
	return &dropOffLocationUsecase{
		dropOffLocationRepo: cr,
		contextTimeout:      timeout,
	}
}

func (usecase *dropOffLocationUsecase) First(c context.Context, dropOffLocationFilter *domain.DropOffLocation) (dropOffLocation domain.DropOffLocation, err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()

	dropOffLocation, err = usecase.dropOffLocationRepo.First(ctx, dropOffLocationFilter)
	if err != nil {
		return domain.DropOffLocation{}, err
	}

	return
}

func (usecase *dropOffLocationUsecase) Fetch(c context.Context) (res []domain.DropOffLocation, err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()

	res, err = usecase.dropOffLocationRepo.Fetch(ctx)
	if err != nil {
		return make([]domain.DropOffLocation, 0), err
	}

	return
}

func (usecase *dropOffLocationUsecase) Store(c context.Context, dropOffLocationStoreRequest *domain.DropOffLocationStoreRequest) (err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()

	dropOffLocation := domain.DropOffLocation{
		Name:          dropOffLocationStoreRequest.Name,
		PIC:           dropOffLocationStoreRequest.PIC,
		Phone:         dropOffLocationStoreRequest.Phone,
		Email:         dropOffLocationStoreRequest.Email,
		Coordinate:    dropOffLocationStoreRequest.Coordinate,
		Address:       dropOffLocationStoreRequest.Address,
		MaxSupply:     dropOffLocationStoreRequest.MaxSupply,
		CurrentSupply: dropOffLocationStoreRequest.CurrentSupply,
	}

	err = usecase.dropOffLocationRepo.Store(ctx, &dropOffLocation)

	return
}

func (usecase *dropOffLocationUsecase) Update(c context.Context, dropOffLocation *domain.DropOffLocation, dropOffLocationUpdateRequest *domain.DropOffLocationUpdateRequest) (err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()

	dropOffLocation.Name = dropOffLocationUpdateRequest.Name
	dropOffLocation.PIC = dropOffLocationUpdateRequest.PIC
	dropOffLocation.Phone = dropOffLocationUpdateRequest.Phone
	dropOffLocation.Email = dropOffLocationUpdateRequest.Email
	dropOffLocation.Coordinate = dropOffLocationUpdateRequest.Coordinate
	dropOffLocation.Address = dropOffLocationUpdateRequest.Address
	dropOffLocation.MaxSupply = dropOffLocationUpdateRequest.MaxSupply
	dropOffLocation.CurrentSupply = dropOffLocationUpdateRequest.CurrentSupply

	dropOffLocation.UpdatedAt = time.Now().Unix()

	return usecase.dropOffLocationRepo.Update(ctx, dropOffLocation)
}

func (usecase *dropOffLocationUsecase) Delete(c context.Context, dropOffLocation *domain.DropOffLocation) (err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()

	err = usecase.dropOffLocationRepo.Delete(ctx, dropOffLocation)

	return
}
