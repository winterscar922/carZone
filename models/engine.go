package models

import (
	"errors"
	"time"
)

type Engine struct {
	EngineId       int       `json:"engine_id"`
	Displacement   int64     `json:"displacement"`
	CylindersCount int64     `json:"cylinders_count"`
	CarRange       int64     `json:"car_range"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type EngineRequest struct {
	Displacement   int64 `json:"displacement"`
	CylindersCount int64 `json:"cylinders_count"`
	CarRange       int64 `json:"car_range"`
}

func ValidateEngineRequest(engine Engine) error {
	if engine.Displacement <= 0 {
		return errors.New("displacement must be greater thatn zero")
	}
	if engine.CylindersCount <= 0 {
		return errors.New("cylinder count must be greater thatn zero")
	}
	if engine.CarRange <= 0 {
		return errors.New("car range must be greater thatn zero")
	}
	return nil
}
