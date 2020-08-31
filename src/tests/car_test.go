package tests

import (
	"fmt"
	"testing"

	"github.com/ScriptStudio/SeekWorkerWebBack/src/models"
)

func TestCar_Validate(t *testing.T) {
	type fields struct {
		CarPlate     string
		TankCapacity float32
		Consumption  float32
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "CarWithInvalidCarPlate",
			fields: fields{
				CarPlate:     "QPSSS78",
				TankCapacity: 45,
				Consumption:  11,
			},
			wantErr: true,
		},
		{
			name: "CarWithInvalidTankCapacity",
			fields: fields{
				CarPlate:     "QPS8103",
				TankCapacity: -1,
				Consumption:  11,
			},
			wantErr: true,
		},
		{
			name: "CarWithInvalidConsumption",
			fields: fields{
				CarPlate:     "QPS8103",
				TankCapacity: 45,
				Consumption:  0,
			},
			wantErr: true,
		},
		{
			name: "ValidCarWithLegacyCarPlate",
			fields: fields{
				CarPlate:     "QPS8103",
				TankCapacity: 45,
				Consumption:  11,
			},
			wantErr: false,
		},
		{
			name: "ValidCarWithLegacyCarPlate",
			fields: fields{
				CarPlate:     "QPS4A44",
				TankCapacity: 45,
				Consumption:  11,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &models.Car{
				CarPlate:     tt.fields.CarPlate,
				TankCapacity: tt.fields.TankCapacity,
				Consumption:  tt.fields.Consumption,
			}

			var err error = c.Validate()

			if err != nil {
				fmt.Printf("Test: %s, Erro: %s, Was expected?: %t\n", tt.name, err.Error(), tt.wantErr)
			}

			if err := c.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Car.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCar_Autonomy(t *testing.T) {
	type fields struct {
		CarPlate     string
		TankCapacity float32
		Consumption  float32
	}
	tests := []struct {
		name   string
		fields fields
		want   float32
	}{
		// TODO: Add test cases.
		{
			name: "ValidCarWithLegacyCarPlate",
			fields: fields{
				CarPlate:     "QPS4A44",
				TankCapacity: 45,
				Consumption:  11,
			},
			want: 495,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &models.Car{
				CarPlate:     tt.fields.CarPlate,
				TankCapacity: tt.fields.TankCapacity,
				Consumption:  tt.fields.Consumption,
			}
			if got := c.Autonomy(); got != tt.want {
				t.Errorf("Car.Autonomy() = %v, want %v", got, tt.want)
			}
		})
	}
}
