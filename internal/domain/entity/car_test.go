package entity

import (
	"strconv"
	"testing"
	"time"
)

func TestNewCar(t *testing.T) {
	currentYear := time.Now().Year()
	type Fuel int

	const (
		E100 Fuel = iota
		E60
		E30
	)
	type args struct {
		name         string
		manufacturer string
		model        string
		year         string
		modelYear    string
		fuelType     Fuel
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
		errMsg  string
	}{
		{
			name: "valid car",
			args: args{
				name:         "Valid Name",
				manufacturer: "Manufacturer",
				model:        "Model",
				year:         strconv.Itoa(currentYear),
				modelYear:    "2023",
				fuelType:     0,
			},
			wantErr: false,
		},
		// Name validation tests
		{
			name:    "empty name",
			args:    args{name: "", manufacturer: "M", model: "M", year: "2020", modelYear: "2020"},
			wantErr: true,
			errMsg:  "name cannot be empty",
		},
		{
			name:    "name too short",
			args:    args{name: "Na", manufacturer: "M", model: "M", year: "2020", modelYear: "2020"},
			wantErr: true,
			errMsg:  "name should be between 3 and 130 characters",
		},
		{
			name:    "name too long",
			args:    args{name: createString(131), manufacturer: "M", model: "M", year: "2020", modelYear: "2020"},
			wantErr: true,
			errMsg:  "name should be between 3 and 130 characters",
		},
		// Manufacturer validation tests
		{
			name:    "empty manufacturer",
			args:    args{name: "Name", manufacturer: "", model: "M", year: "2020", modelYear: "2020"},
			wantErr: true,
			errMsg:  "manufacturer cannot be empty",
		},
		{
			name:    "manufacturer too short",
			args:    args{name: "Name", manufacturer: "Ma", model: "M", year: "2020", modelYear: "2020"},
			wantErr: true,
			errMsg:  "manufacturer must be between 3 and 150 characters",
		},
		// Model validation tests
		{
			name:    "empty model",
			args:    args{name: "Name", manufacturer: "Manu", model: "", year: "2020", modelYear: "2020"},
			wantErr: true,
			errMsg:  "model cannot be empty",
		},
		{
			name:    "model too short",
			args:    args{name: "Name", manufacturer: "Manu", model: "Mo", year: "2020", modelYear: "2020"},
			wantErr: true,
			errMsg:  "model must be between 3 and 150 characters",
		},
		// Year validation tests
		{
			name:    "invalid year length",
			args:    args{name: "Name", manufacturer: "Manu", model: "Model", year: "202", modelYear: "2020"},
			wantErr: true,
			errMsg:  "invalid year",
		},
		{
			name:    "non-numeric year",
			args:    args{name: "Name", manufacturer: "Manu", model: "Model", year: "20ab", modelYear: "2020"},
			wantErr: true,
			errMsg:  "invalid year",
		},
		{
			name:    "year too far in future",
			args:    args{name: "Name", manufacturer: "Manu", model: "Model", year: strconv.Itoa(currentYear + 2), modelYear: "2020"},
			wantErr: true,
			errMsg:  "year if out of the valid range",
		},
		// ModelYear validation tests
		{
			name:    "modelYear too short",
			args:    args{name: "Name", manufacturer: "Manu", model: "Model", year: "2020", modelYear: "20"},
			wantErr: true,
			errMsg:  "modelYear must be between 3 and 150 characters",
		},
		{
			name:    "modelYear too long",
			args:    args{name: "Name", manufacturer: "Manu", model: "Model", year: "2020", modelYear: createString(131)},
			wantErr: true,
			errMsg:  "modelYear must be between 3 and 150 characters",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewCar(
				tt.args.name,
				tt.args.manufacturer,
				tt.args.model,
				tt.args.year,
				tt.args.modelYear,
				0,
			)

			if (err != nil) != tt.wantErr {
				t.Errorf("NewCar() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr && err.Error() != tt.errMsg {
				t.Errorf("Expected error message '%s', got '%s'", tt.errMsg, err.Error())
			}
		})
	}
}

// Helper to create long strings for validation
func createString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = 'a'
	}
	return string(b)
}
