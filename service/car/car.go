package car

import (
	"context"

	"github.com/winterscar922/carZone/models"
	"github.com/winterscar922/carZone/store"
)

type CarService struct {
	store store.CarStoreInterface
}

func NewService(store store.CarStoreInterface) *CarService {
	return &CarService{store: store}
}

func (s *CarService) GetCarById(ctx context.Context, id int64) (models.Car, error) {
	car, err := s.store.GetCarById(ctx, id)
	if err != nil {
		return models.Car{}, err
	}
	return car, nil
}

func (s *CarService) CreateCar(ctx context.Context, carReq models.CarRequest) (models.Car, error) {
	car, err := s.store.CreateCar(ctx, carReq)
	if err != nil {
		return models.Car{}, err
	}
	return car, nil
}
