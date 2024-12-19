package models

import (
	"errors"
	"strconv"
	"time"
)

type Car struct {
	CarId     int64     `json:"car_id"`
	Name      string    `json:"name"`
	Year      string    `json:"year"`
	Brand     string    `json:"brand"`
	FuelType  string    `json:"fuel_type"`
	Engine    Engine    `json:"engine"`
	Price     int64     `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CarRequest struct {
	Name     string `json:"name"`
	Year     string `json:"year"`
	Brand    string `json:"brand"`
	FuelType string `json:"fuel_type"`
	EngineId int64  `json:"engine_id"`
	Price    int64  `json:"price"`
}

func ValidateCarRequest(carRequest CarRequest) error {
	if err := ValidateName(carRequest.Name); err != nil {
		return err
	}
	if err := ValidateYear(carRequest.Year); err != nil {
		return err
	}
	if err := ValidateBrand(carRequest.Brand); err != nil {
		return err
	}
	if err := ValidateFuelTypes(carRequest.FuelType); err != nil {
		return err
	}
	if err := ValidateEngine(carRequest.EngineId); err != nil {
		return err
	}
	if err := ValidatePrice(carRequest.Price); err != nil {
		return err
	}
	return nil
}

func ValidateName(name string) error {
	if name == "" {
		return errors.New("car Name is required")
	}
	return nil
}

func ValidateYear(dateTime string) error {
	year, err := strconv.Atoi(dateTime)
	if err != nil {
		return errors.New("error while converting year")
	}
	if year <= 1800 || year > time.Now().Year() {
		return errors.New("year must be between 1800 and current year")
	}
	return nil
}

func ValidateBrand(brand string) error {
	if brand == "" {
		return errors.New("brand is required")
	}
	return nil
}

func ValidateFuelTypes(fuelType string) error {
	validFuelTypes := []string{"Petrol", "Diesel", "Electric", "Hybrid"}

	for _, validFuelType := range validFuelTypes {
		if validFuelType == fuelType {
			return nil
		}
	}

	errMessage := "fuel type must be on of these: "

	for index, validvalidFuelType := range validFuelTypes {
		errMessage += validvalidFuelType
		if index != len(validFuelTypes)-1 {
			errMessage += ", "
		}
	}

	return errors.New(errMessage)
}

func ValidateEngine(engineId int64) error {
	if engineId == 0 {
		return errors.New("engine id is required")
	}
	return nil
}

func ValidatePrice(price int64) error {
	if price <= 0 {
		return errors.New("price must be greater thatn zero")
	}
	return nil
}
