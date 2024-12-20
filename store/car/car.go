package car

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/winterscar922/carZone/models"
	engineDataStore "github.com/winterscar922/carZone/store/engine"
)

type Store struct {
	Db *sql.DB
}

func Open(db *sql.DB) *Store {
	return &Store{Db: db}
}

func (s *Store) CreateCar(ctx context.Context, carReq models.CarRequest) (models.Car, error) {
	var car models.Car

	engineStore := engineDataStore.Store{Db: s.Db}
	engine, err := engineStore.GetEngineById(ctx, carReq.EngineId)

	if err != nil {
		return models.Car{}, err
	}

	query := `insert into car (name, year, brand, fuel_type, engine_id, price, created_at, updated_at) 
	values ($1, $2, $3, $4, $5, $6, $7, $8) 
	returning id, name, year, brand, fuel_type, engine_id, price, created_at, updated_at`

	err = s.Db.QueryRowContext(ctx, query,
		carReq.Name,
		carReq.Year,
		carReq.Brand,
		carReq.FuelType,
		carReq.EngineId,
		carReq.Price,
		time.Now(),
		time.Now()).Scan(
		&car.CarId,
		&car.Name,
		&car.Year,
		&car.Brand,
		&car.FuelType,
		&car.Engine.EngineId,
		&car.Price,
		&car.CreatedAt,
		&car.UpdatedAt,
	)

	car.Engine = engine

	if err != nil {
		return models.Car{}, errors.New("error while inserting car")
	}
	return car, nil
}

func (s *Store) UpdateCar(ctx context.Context, carReq models.CarRequest, id int64) error {
	tx, err := s.Db.BeginTx(ctx, nil)

	if err != nil {
		return errors.New(fmt.Sprintf("error while creating transaction for updating car with id - %d", id))
	}

	engineStore := engineDataStore.Store{Db: s.Db}
	_, err = engineStore.GetEngineById(ctx, carReq.EngineId)

	if err != nil {
		return err
	}

	query := `update car
				set name=$1, year=$2, brand=$3, fuel_type=$4, engine_id=$5, price=$6, updated_at=$7
				where id = $8`

	res, err := s.Db.ExecContext(ctx, query,
		carReq.Name,
		carReq.Year,
		carReq.Brand,
		carReq.FuelType,
		carReq.EngineId,
		carReq.Price,
		time.Now(),
		id)

	if err != nil {
		tx.Rollback()
		return errors.New(fmt.Sprintf("error while updating car with id - %d", id))
	}

	rowsAffected, _ := res.RowsAffected()

	if rowsAffected > 1 {
		tx.Rollback()
		return errors.New(fmt.Sprintf("multiple rows were effected while updating car with id - %d, rollbacking changes", id))
	}

	if rowsAffected == 0 {
		return errors.New(fmt.Sprintf("car with id - %d not found", id))
	}

	tx.Commit()
	return nil
}

func (s *Store) DeleteCar(ctx context.Context, id int64) error {
	query := `delete from car
				where id = $1`

	tx, err := s.Db.BeginTx(ctx, nil)

	if err != nil {
		return errors.New(fmt.Sprintf("error while creating transaction for deleting car with id - %d", id))
	}

	res, err := s.Db.ExecContext(ctx, query, id)

	if err != nil {
		tx.Rollback()
		return errors.New(fmt.Sprintf("error while deleting car with id - %d", id))
	}

	rowsAffected, _ := res.RowsAffected()

	if rowsAffected > 1 {
		tx.Rollback()
		return errors.New(fmt.Sprintf("multiple rows were effected while deleting car with id - %d, rollbacking changes", id))
	}

	if rowsAffected == 0 {
		return errors.New(fmt.Sprintf("car with id - %d not found", id))
	}

	tx.Commit()
	return nil
}

func (s *Store) GetCarById(ctx context.Context, id int64) (models.Car, error) {
	var car models.Car
	query := `select c.id, c.name, c.year, c.brand, c.fuel_type, c.price, c.created_at, c.updated_at, c.engine_id, e.displacement, e.cylinders_count, e.car_range, e.created_at, e.updated_at from car c 
	left join engine e on e.id = c.id
	where c.id = $1`

	err := s.Db.QueryRowContext(ctx, query, id).Scan(
		&car.CarId,
		&car.Name,
		&car.Year,
		&car.Brand,
		&car.FuelType,
		&car.Price,
		&car.CreatedAt,
		&car.UpdatedAt,
		&car.Engine.EngineId,
		&car.Engine.Displacement,
		&car.Engine.CylindersCount,
		&car.Engine.CarRange,
		&car.Engine.CreatedAt,
		&car.Engine.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return models.Car{}, errors.New(fmt.Sprintf("car with id - %d not found", id))
		}
		return models.Car{}, errors.New(fmt.Sprintf("error while fetching car with id - %d", id))
	}
	return car, nil
}

func (s *Store) GetCarByBrand(ctx context.Context, brand string) ([]models.Car, error) {
	return []models.Car{}, nil
}
