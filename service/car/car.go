package car

import (
	"context"
	"fmt"

	"github.com/winterscar922/carZone/models"
	"github.com/winterscar922/carZone/store"
)

type CarService struct {
	store store.CarStoreInterface
}

func NewService(store store.CarStoreInterface) *CarService {
	return &CarService{store: store}
}

func (s *CarService) GetCarById(ctx context.Context, id int) (models.Car, error) {
	fmt.Println("went inside one!!!")
	car, err := s.store.GetCarById(ctx, id)
	if err != nil {
		return models.Car{}, err
	}
	return car, nil
}
