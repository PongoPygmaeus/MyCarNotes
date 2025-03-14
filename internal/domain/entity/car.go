package entity

type Car struct {
	Name string
	Manufacturer string
	Model string
	Year string
	ModelYear string
	FuelType Fuel
}

type Fuel int

const (
	E100 Fuel = iota
	E60
	E30
)