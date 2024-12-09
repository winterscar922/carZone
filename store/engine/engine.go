package engine

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/winterscar922/carZone/models"
)

type Store struct {
	Db *sql.DB
}

func Open(db *sql.DB) Store {
	return Store{Db: db}
}

func (s Store) GetEngineById(ctx context.Context, id int) (models.Engine, error) {
	var engine models.Engine
	query := `select * from engine where engine_id = $1`
	err := s.Db.QueryRowContext(ctx, query, id).Scan(&engine)

	if err != nil {
		if err == sql.ErrNoRows {
			return models.Engine{}, fmt.Errorf("no engine found with id %d", id)
		}
		return models.Engine{}, fmt.Errorf("error fetching engine with id %d: %w", id, err)
	}

	return engine, nil
}

func (s Store) CheckEngineById(ctx context.Context, id int) (bool, error) {
	var exists bool
	query := `select 1 from engine where engine_id = $1 limit 1`
	err := s.Db.QueryRowContext(ctx, query, id).Scan(&exists)

	if err != nil {
		return false, fmt.Errorf("error fetching engine with id %d: %w", id, err)
	}

	return exists, nil
}
