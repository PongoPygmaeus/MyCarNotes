package gateway

import (
	"errors"
	"github.com/PongoPygmaeus/MyCarNotes/internal/domain/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMockCarGateway(t *testing.T) {
	t.Run("should initialize empty car map", func(t *testing.T) {
		gateway := NewMockCarGateway()
		assert.NotNil(t, gateway)
		assert.NotNil(t, gateway.Cars)
		assert.Empty(t, gateway.Cars)
	})
}

func TestMockCarGateway_GetCarById(t *testing.T) {
	tests := []struct {
		name          string
		initialCars   map[int]*entity.Car
		id            int
		expectedCar   *entity.Car
		expectedError error
	}{
		{
			name:          "car not found",
			initialCars:   map[int]*entity.Car{},
			id:            1,
			expectedCar:   nil,
			expectedError: errors.New("car not found"),
		},
		{
			name: "car found",
			initialCars: map[int]*entity.Car{
				1: {Id: 1, Name: "Test Car"},
			},
			id:            1,
			expectedCar:   &entity.Car{Id: 1, Name: "Test Car"},
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gateway := &MockCarGateway{
				Cars: tt.initialCars,
			}

			car, err := gateway.GetCarById(tt.id)

			if tt.expectedError != nil {
				assert.EqualError(t, err, tt.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.expectedCar, car)
		})
	}
}

func TestMockCarGateway_SaveCar(t *testing.T) {
	tests := []struct {
		name         string
		initialCars  map[int]*entity.Car
		carToSave    *entity.Car
		expectedCars map[int]*entity.Car
	}{
		{
			name:        "save new car to empty storage",
			initialCars: map[int]*entity.Car{},
			carToSave:   &entity.Car{Id: 1, Name: "New Car"},
			expectedCars: map[int]*entity.Car{
				1: {Id: 1, Name: "New Car"},
			},
		},
		{
			name: "overwrite existing car",
			initialCars: map[int]*entity.Car{
				1: {Id: 1, Name: "Old Car"},
			},
			carToSave: &entity.Car{Id: 1, Name: "Updated Car"},
			expectedCars: map[int]*entity.Car{
				1: {Id: 1, Name: "Updated Car"},
			},
		},
		{
			name: "save multiple cars",
			initialCars: map[int]*entity.Car{
				1: {Id: 1, Name: "Car 1"},
			},
			carToSave: &entity.Car{Id: 2, Name: "Car 2"},
			expectedCars: map[int]*entity.Car{
				1: {Id: 1, Name: "Car 1"},
				2: {Id: 2, Name: "Car 2"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gateway := &MockCarGateway{
				Cars: tt.initialCars,
			}

			err := gateway.SaveCar(tt.carToSave)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedCars, gateway.Cars)
		})
	}
}

func TestMockCarGateway_Integration(t *testing.T) {
	t.Run("save and retrieve car", func(t *testing.T) {
		gateway := NewMockCarGateway()
		testCar := &entity.Car{Id: 1, Name: "Integration Test Car"}

		// Save the car
		err := gateway.SaveCar(testCar)
		assert.NoError(t, err)

		// Retrieve the car
		retrievedCar, err := gateway.GetCarById(1)
		assert.NoError(t, err)
		assert.Equal(t, testCar, retrievedCar)

		// Verify non-existent car
		_, err = gateway.GetCarById(2)
		assert.EqualError(t, err, "car not found")
	})
}
