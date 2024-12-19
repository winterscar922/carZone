package service

import (
	"context"

	"github.com/winterscar922/carZone/models"
)

type CarServiceInterface interface {
	GetCarById(ctx context.Context, id int64) (models.Car, error)
	CreateCar(ctx context.Context, carReq models.CarRequest) (models.Car, error)
}

type EngineServiceInterface interface {
	GetEngineById(ctx context.Context, id int64) (models.Engine, error)
	CreateEngine(ctx context.Context, engineReq models.EngineRequest) (models.Engine, error)
}
