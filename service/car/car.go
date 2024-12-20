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
	return s.store.GetCarById(ctx, id)
}

func (s *CarService) CreateCar(ctx context.Context, carReq models.CarRequest) (models.Car, error) {
	return s.store.CreateCar(ctx, carReq)
}

func (s *CarService) UpdateCar(ctx context.Context, carReq models.CarRequest, id int64) error {
	return s.store.UpdateCar(ctx, carReq, id)
}

func (s *CarService) DeleteCar(ctx context.Context, id int64) error {
	return s.store.DeleteCar(ctx, id)
}

func (s *CarService) GetAllCars(ctx context.Context) ([]models.Car, error) {
	return s.store.GetAllCars(ctx)
}
