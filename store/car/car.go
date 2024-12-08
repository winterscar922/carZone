package car

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/winterscar922/carZone/models"
	engineDataStore "github.com/winterscar922/carZone/store/engine"
)

type Store struct {
	db *sql.DB
}

func open(db *sql.DB) Store {
	return Store{db: db}
}

func (s Store) InsertCar(ctx context.Context, carReq models.CarRequest) (models.Car, error) {
	var car models.Car

	// verify if engine id is present or not in engine table
	var engine_id = carReq.Engine.EngineId

	if engine_id != uuid.Nil {
		engineStore := engineDataStore.Store{Db: s.db}
		exists, err := engineStore.CheckEngineById(ctx, engine_id)
		if err != nil {
			return models.Car{}, err
		}
		if !exists {
			return models.Car{}, fmt.Errorf("no engine found with id %d", engine_id)
		}
	}

	query := `insert into car (id, name, year, brand, fuel_type, engine_id, price, created_at, modified_at) 
	values ($1,$2,$3,$4,$5,$6,$7,$8,$9) 
	returning id, name, year, brand, fuel_type, engine_id, price, created_at, modified_at`

	err := s.db.QueryRowContext(ctx, query,
		uuid.New(),
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
		&car.ModifiedAt,
	)

	if err != nil {
		return models.Car{}, errors.New("error while inserting car")
	}
	return car, nil
}

func (s Store) UpdateCar(ctx context.Context, carReq models.CarRequest) (models.Car, error) {
	return models.Car{}, nil
}

func (s Store) DeleteCar(ctx context.Context, id int) error {
	return nil
}

func (s Store) GetCarById(ctx context.Context, id int) (models.Car, error) {
	var car models.Car
	query := `select * from car where id = $1`

	err := s.db.QueryRowContext(ctx, query, id).Scan(&car)

	if err != nil {
		return models.Car{}, errors.New(fmt.Sprintf("error while fetching car with id - %d", id))
	}
	return car, nil
}

func (s Store) GetCarByBrand(ctx context.Context, brand string) ([]models.Car, error) {
	return []models.Car{}, nil
}
