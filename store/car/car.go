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

func (s *Store) InsertCar(ctx context.Context, carReq models.CarRequest) (models.Car, error) {
	var car models.Car

	// verify if engine id is present or not in engine table
	var engine_id = carReq.Engine.EngineId

	if engine_id != 0 {
		engineStore := engineDataStore.Store{Db: s.Db}
		exists, err := engineStore.CheckEngineById(ctx, engine_id)
		if err != nil {
			return models.Car{}, err
		}
		if !exists {
			return models.Car{}, fmt.Errorf("no engine found with id %d", engine_id)
		}
	}

	query := `insert into car (name, year, brand, fuel_type, engine_id, price, created_at, modified_at) 
	values ($1,$2,$3,$4,$5,$6,$7,$8) 
	returning id, name, year, brand, fuel_type, engine_id, price, created_at, modified_at`

	err := s.Db.QueryRowContext(ctx, query,
		carReq.Name,
		carReq.Year,
		carReq.Brand,
		carReq.FuelType,
		engine_id,
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

	if err != nil {
		return models.Car{}, errors.New("error while inserting car")
	}
	return car, nil
}

func (s *Store) UpdateCar(ctx context.Context, carReq models.CarRequest) (models.Car, error) {
	return models.Car{}, nil
}

func (s *Store) DeleteCar(ctx context.Context, id int) error {
	return nil
}

func (s *Store) GetCarById(ctx context.Context, id int) (models.Car, error) {
	var car models.Car
	query := `select c.id, c.name, c.year, c.brand, c.fuel_type, c.price, c.created_at, c.updated_at, c.engine_id, e.displacement, e.no_of_cylinders, e.car_range, e.created_at, e.updated_at from car c 
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
		return models.Car{}, errors.New(fmt.Sprintf("error while fetching car with id - %d", id))
	}
	return car, nil
}

func (s *Store) GetCarByBrand(ctx context.Context, brand string) ([]models.Car, error) {
	return []models.Car{}, nil
}
