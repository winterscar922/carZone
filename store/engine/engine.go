package engine

import (
	"context"
	"database/sql"
	"errors"
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
	query := `select id, car_range, cylinders_count, displacement, created_at, updated_at 
				from engine where id = $1`

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

func (s *Store) UpdateEngine(ctx context.Context, engineReq models.EngineRequest, id int64) error {
	query := `update engine
				set displacement=$1, cylinders_count=$2, car_range=$3, updated_at=$4
				where id = $5`

	tx, err := s.Db.BeginTx(ctx, nil)

	if err != nil {
		return errors.New(fmt.Sprintf("error while creating transaction for updating car with id - %d", id))
	}

	res, err := s.Db.ExecContext(ctx, query,
		engineReq.Displacement,
		engineReq.CylindersCount,
		engineReq.CarRange,
		time.Now(),
		id)

	if err != nil {
		tx.Rollback()
		return errors.New(fmt.Sprintf("error while updating engine with id - %d", id))
	}

	rowsAffected, _ := res.RowsAffected()

	if rowsAffected > 1 {
		tx.Rollback()
		return errors.New(fmt.Sprintf("multiple rows were effected while updating engine with id - %d, rollbacking changes", id))
	}

	if rowsAffected == 0 {
		return errors.New(fmt.Sprintf("engine with id - %d not found", id))
	}

	tx.Commit()
	return nil
}

func (s *Store) DeleteEngine(ctx context.Context, id int64) error {
	query := `delete from engine
				where id = $1`

	tx, err := s.Db.BeginTx(ctx, nil)

	if err != nil {
		return errors.New(fmt.Sprintf("error while creating transaction for deleting engine with id - %d", id))
	}

	res, err := s.Db.ExecContext(ctx, query, id)

	if err != nil {
		tx.Rollback()
		return errors.New(fmt.Sprintf("error while deleting engine with id - %d", id))
	}

	rowsAffected, _ := res.RowsAffected()

	if rowsAffected > 1 {
		tx.Rollback()
		return errors.New(fmt.Sprintf("multiple rows were effected while deleting engine with id - %d, rollbacking changes", id))
	}

	if rowsAffected == 0 {
		return errors.New(fmt.Sprintf("engine with id - %d not found", id))
	}

	tx.Commit()
	return nil
}

func (s *Store) GetAllEngines(ctx context.Context) ([]models.Engine, error) {
	var engines []models.Engine
	query := `select * from engine`

	rows, err := s.Db.QueryContext(ctx, query)

	if err != nil {
		return []models.Engine{}, errors.New("error while fetching engines")
	}

	for rows.Next() {
		var engine models.Engine

		rows.Scan(
			&engine.EngineId,
			&engine.CarRange,
			&engine.CylindersCount,
			&engine.Displacement,
			&engine.CreatedAt,
			&engine.UpdatedAt,
		)

		engines = append(engines, engine)
	}
	return engines, nil
}
