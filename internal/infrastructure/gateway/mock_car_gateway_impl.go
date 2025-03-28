package gateway

import (
	"errors"
	"github.com/PongoPygmaeus/MyCarNotes/internal/domain/entity"
)

type MockCarGateway struct {
	Cars map[int]*entity.Car
}

func NewMockCarGateway() *MockCarGateway {
	return &MockCarGateway{
		Cars: make(map[int]*entity.Car),
	}
}

func (m *MockCarGateway) GetCarById(id int) (*entity.Car, error) {
	user, ok := m.Cars[id]
	if !ok {
		return nil, errors.New("car not found")
	}
	return user, nil
}

func (m *MockCarGateway) SaveCar(car *entity.Car) error {
	m.Cars[car.Id] = car
	return nil
}
