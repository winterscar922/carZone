package store

import (
	"context"

	"github.com/winterscar922/carZone/models"
)

type CarStoreInterface interface {
	GetCarById(ctx context.Context, id int) (models.Car, error)
}

type EngineStoreInterface interface {
	GetEngineById(ctx context.Context, id int) (models.Engine, error)
}
