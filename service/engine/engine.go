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

func (s *EngineService) GetEngineById(ctx context.Context, id int64) (models.Engine, error) {
	return s.store.GetEngineById(ctx, id)
}

func (s *EngineService) CreateEngine(ctx context.Context, engineReq models.EngineRequest) (models.Engine, error) {
	return s.store.CreateEngine(ctx, engineReq)
}

func (s *EngineService) UpdateEngine(ctx context.Context, engineReq models.EngineRequest, id int64) error {
	return s.store.UpdateEngine(ctx, engineReq, id)
}

func (s *EngineService) DeleteEngine(ctx context.Context, id int64) error {
	return s.store.DeleteEngine(ctx, id)
}

func (s *EngineService) GetAllEngines(ctx context.Context) ([]models.Engine, error) {
	return s.store.GetAllEngines(ctx)
}
