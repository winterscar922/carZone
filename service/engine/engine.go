package engine

import (
	"context"

	"github.com/winterscar922/carZone/models"
	"github.com/winterscar922/carZone/store"
)

type EngineService struct {
	store store.EngineStoreInterface
}

func NewService(store store.EngineStoreInterface) *EngineService {
	return &EngineService{store: store}
}

func (s *EngineService) GetEngineById(ctx context.Context, id int) (models.Engine, error) {
	engine, err := s.store.GetEngineById(ctx, id)
	if err != nil {
		return models.Engine{}, err
	}
	return engine, nil
}
