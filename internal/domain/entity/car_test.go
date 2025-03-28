package entity

import (
	"errors"
	"strconv"
	"testing"
	"time"
)

func TestNewCar(t *testing.T) {
	tests := []struct {
		name          string
		input         func() (*Car, error)
		expectedError error
	}{
		{
			name: "valid car",
			input: func() (*Car, error) {
				return NewCar("My Car", "Toyota", "Corolla", "2020", "2020", E100)
			},
			expectedError: nil,
		},
		{
			name: "invalid name",
			input: func() (*Car, error) {
				return NewCar("", "Toyota", "Corolla", "2020", "2020", E60)
			},
			expectedError: errors.New("name cannot be empty"),
		},
		{
			name: "invalid manufacturer",
			input: func() (*Car, error) {
				return NewCar("My Car", "", "Corolla", "2020", "2020", E30)
			},
			expectedError: errors.New("manufacturer cannot be empty"),
		},
		{
			name: "invalid model",
			input: func() (*Car, error) {
				return NewCar("My Car", "Toyota", "", "2020", "2020", E100)
			},
			expectedError: errors.New("model cannot be empty"),
		},
		{
			name: "invalid year",
			input: func() (*Car, error) {
				return NewCar("My Car", "Toyota", "Corolla", "202", "2020", E30)
			},
			expectedError: errors.New("invalid year"),
		},
		{
			name: "future year",
			input: func() (*Car, error) {
				futureYear := strconv.Itoa(time.Now().Year() + 2)
				return NewCar("My Car", "Toyota", "Corolla", futureYear, "2020", E100)
			},
			expectedError: errors.New("year if out of the valid range"),
		},
		{
			name: "invalid model year",
			input: func() (*Car, error) {
				return NewCar("My Car", "Toyota", "Corolla", "2020", "", E30)
			},
			expectedError: errors.New("modelYear cannot be empty"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.input()
			if (err != nil) != (tt.expectedError != nil) {
				t.Errorf("expected error %v, got %v", tt.expectedError, err)
			}
			if err != nil && tt.expectedError != nil && err.Error() != tt.expectedError.Error() {
				t.Errorf("expected error message %q, got %q", tt.expectedError.Error(), err.Error())
			}
		})
	}
}

func TestValidateName(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		expectedError error
	}{
		{"valid name", "Valid Car Name", nil},
		{"empty name", "", errors.New("name cannot be empty")},
		{"too short name", "a", errors.New("name should be between 3 and 130 characters")},
		{"too long name", string(make([]byte, 131)), errors.New("name should be between 3 and 130 characters")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateName(tt.input)
			if (err != nil) != (tt.expectedError != nil) {
				t.Errorf("expected error %v, got %v", tt.expectedError, err)
			}
			if err != nil && tt.expectedError != nil && err.Error() != tt.expectedError.Error() {
				t.Errorf("expected error message %q, got %q", tt.expectedError.Error(), err.Error())
			}
		})
	}
}

func TestValidateManufacturer(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		expectedError error
	}{
		{"valid manufacturer", "Toyota", nil},
		{"empty manufacturer", "", errors.New("manufacturer cannot be empty")},
		{"too short manufacturer", "a", errors.New("manufacturer must be between 3 and 150 characters")},
		{"too long manufacturer", string(make([]byte, 151)), errors.New("manufacturer must be between 3 and 150 characters")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateManufacturer(tt.input)
			if (err != nil) != (tt.expectedError != nil) {
				t.Errorf("expected error %v, got %v", tt.expectedError, err)
			}
			if err != nil && tt.expectedError != nil && err.Error() != tt.expectedError.Error() {
				t.Errorf("expected error message %q, got %q", tt.expectedError.Error(), err.Error())
			}
		})
	}
}

func TestValidateModel(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		expectedError error
	}{
		{"valid model", "Corolla", nil},
		{"empty model", "", errors.New("model cannot be empty")},
		{"too short model", "a", errors.New("model must be between 3 and 150 characters")},
		{"too long model", string(make([]byte, 151)), errors.New("model must be between 3 and 150 characters")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateModel(tt.input)
			if (err != nil) != (tt.expectedError != nil) {
				t.Errorf("expected error %v, got %v", tt.expectedError, err)
			}
			if err != nil && tt.expectedError != nil && err.Error() != tt.expectedError.Error() {
				t.Errorf("expected error message %q, got %q", tt.expectedError.Error(), err.Error())
			}
		})
	}
}

func TestValidateYear(t *testing.T) {
	currentYear := time.Now().Year()
	tests := []struct {
		name          string
		input         string
		expectedError error
	}{
		{"valid year", "2020", nil},
		{"valid next year", strconv.Itoa(currentYear + 1), nil},
		{"invalid length", "202", errors.New("invalid year")},
		{"not a number", "abcd", errors.New("invalid year")},
		{"future year", strconv.Itoa(currentYear + 2), errors.New("year if out of the valid range")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateYear(tt.input)
			if (err != nil) != (tt.expectedError != nil) {
				t.Errorf("expected error %v, got %v", tt.expectedError, err)
			}
			if err != nil && tt.expectedError != nil && err.Error() != tt.expectedError.Error() {
				t.Errorf("expected error message %q, got %q", tt.expectedError.Error(), err.Error())
			}
		})
	}
}

func TestValidateModelYear(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		expectedError error
	}{
		{"valid model year", "2020", nil},
		{"empty model year", "", errors.New("modelYear cannot be empty")},
		{"too short model year", "a", errors.New("modelYear must be between 3 and 150 characters")},
		{"too long model year", string(make([]byte, 151)), errors.New("modelYear must be between 3 and 150 characters")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateModelYear(tt.input)
			if (err != nil) != (tt.expectedError != nil) {
				t.Errorf("expected error %v, got %v", tt.expectedError, err)
			}
			if err != nil && tt.expectedError != nil && err.Error() != tt.expectedError.Error() {
				t.Errorf("expected error message %q, got %q", tt.expectedError.Error(), err.Error())
			}
		})
	}
}
