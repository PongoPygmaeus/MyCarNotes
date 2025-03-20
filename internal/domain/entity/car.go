package entity

import (
	"errors"
	"strconv"
	"time"
)

type Car struct {
	Name         string
	Manufacturer string
	Model        string
	Year         string
	ModelYear    string
	FuelType     Fuel
}

type Fuel int

const (
	E100 Fuel = iota
	E60
	E30
)

func NewCar(name, manufacturer, model, year, modelYear string, fuelType Fuel) (*Car, error) {
	if err := validateName(name); err != nil {
		return nil, err
	}

	if err := validateManufacturer(manufacturer); err != nil {
		return nil, err
	}

	if err := validateModel(model); err != nil {
		return nil, err
	}

	if err := validateYear(year); err != nil {
		return nil, err
	}

	if err := validateModelYear(modelYear); err != nil {
		return nil, err
	}

	return &Car{
		Name:         name,
		Manufacturer: manufacturer,
		Model:        model,
		Year:         year,
		ModelYear:    modelYear,
		FuelType:     fuelType,
	}, nil
}

func validateModelYear(modelYear string) error {
	if modelYear == "" {
		return errors.New("modelYear cannot be empty")
	}
	if len(modelYear) < 3 || len(modelYear) > 130 {
		return errors.New("modelYear must be between 3 and 150 characters")
	}

	return nil

}

func validateYear(year string) error {
	if len(year) != 4 {
		return errors.New("invalid year")
	}
	yearInt, err := strconv.Atoi(year)
	if err != nil {
		return errors.New("invalid year")
	}
	currentYear := time.Now().Year()
	if yearInt > currentYear+1 {
		return errors.New("year if out of the valid range")
	}
	return nil

}

func validateModel(model string) error {
	if model == "" {
		return errors.New("model cannot be empty")
	}
	if len(model) < 3 || len(model) > 130 {
		return errors.New("model must be between 3 and 150 characters")
	}
	return nil
}

func validateManufacturer(manufacturer string) error {
	if manufacturer == "" {
		return errors.New("manufacturer cannot be empty")
	}
	if len(manufacturer) < 3 || len(manufacturer) > 130 {
		return errors.New("manufacturer must be between 3 and 150 characters")
	}
	return nil
}

func validateName(name string) error {
	if name == "" {
		return errors.New("name cannot be empty")
	}
	if len(name) < 3 || len(name) > 130 {
		return errors.New("name should be between 3 and 130 characters")
	}
	return nil
}
