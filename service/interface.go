package service

import (
	"context"

	"github.com/winterscar922/carZone/models"
)

type CarServiceInterface interface {
	GetCarById(ctx context.Context, id int64) (models.Car, error)
	CreateCar(ctx context.Context, carReq models.CarRequest) (models.Car, error)
	UpdateCar(ctx context.Context, carReq models.CarRequest, id int64) error
	DeleteCar(ctx context.Context, id int64) error
	GetAllCars(ctx context.Context) ([]models.Car, error)
}

type EngineServiceInterface interface {
	GetEngineById(ctx context.Context, id int64) (models.Engine, error)
	CreateEngine(ctx context.Context, engineReq models.EngineRequest) (models.Engine, error)
	UpdateEngine(ctx context.Context, engineReq models.EngineRequest, id int64) error
	DeleteEngine(ctx context.Context, id int64) error
	GetAllEngines(ctx context.Context) ([]models.Engine, error)
}
