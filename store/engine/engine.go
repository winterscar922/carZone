package engine

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/winterscar922/carZone/models"
)

type Store struct {
	Db *sql.DB
}

func Open(db *sql.DB) *Store {
	return &Store{Db: db}
}

func (s *Store) CreateEngine(ctx context.Context, engineReq models.EngineRequest) (models.Engine, error) {
	var newEngine models.Engine

	query := `insert into engine (displacement, car_range, cylinders_count, created_at, updated_at)
				values ($1,$2,$3,$4,$5)
				returning id, displacement, car_range, cylinders_count, created_at, updated_at`

	err := s.Db.QueryRowContext(ctx, query, engineReq.Displacement, engineReq.CarRange,
		engineReq.CylindersCount, time.Now(), time.Now()).Scan(
		&newEngine.EngineId,
		&newEngine.Displacement,
		&newEngine.CarRange,
		&newEngine.CylindersCount,
		&newEngine.CreatedAt,
		&newEngine.UpdatedAt,
	)

	if err != nil {
		return models.Engine{}, fmt.Errorf("error inserting car")
	}

	return newEngine, nil
}

func (s *Store) GetEngineById(ctx context.Context, id int64) (models.Engine, error) {
	var engine models.Engine
	query := `select * from engine where id = $1`
	err := s.Db.QueryRowContext(ctx, query, id).Scan(
		&engine.EngineId,
		&engine.CarRange,
		&engine.CylindersCount,
		&engine.Displacement,
		&engine.CreatedAt,
		&engine.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return models.Engine{}, fmt.Errorf("no engine found with id %d", id)
		}
		return models.Engine{}, fmt.Errorf("error fetching engine with id %d: %w", id, err)
	}

	return engine, nil
}
