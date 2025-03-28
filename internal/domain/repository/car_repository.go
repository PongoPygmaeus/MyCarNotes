package repository

import "github.com/PongoPygmaeus/MyCarNotes/internal/domain/entity"

type CarRepository interface {
	GetCarById(id int) (*entity.Car, error)
	SaveCar(car *entity.Car) error
}
