package tests

import (
	"fmt"
	"testing"

	"github.com/ScriptStudio/SeekWorkerWebBack/src/models"
	"github.com/ScriptStudio/SeekWorkerWebBack/src/models/enums"
)

func TestAddress_Validate(t *testing.T) {
	type fields struct {
		ZipCode      string
		State        enums.State
		City         string
		Neighborhood string
		Street       string
		Complement   string
		Coordinate   [2]float64
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "AddressWithInvalidZipCode",
			fields: fields{
				ZipCode:      "89815-55",
				State:        enums.SC,
				City:         "Chapecó",
				Neighborhood: "Desbravador",
				Street:       "Rua Balduino",
				Complement:   "",
				Coordinate:   [...]float64{-29.3615343, -50.8514137},
			},
			wantErr: true,
		},
		{
			name: "AddressWithInvalidState",
			fields: fields{
				ZipCode:      "89815-555",
				State:        30,
				City:         "Chapecó",
				Neighborhood: "Desbravador",
				Street:       "Rua Balduino",
				Complement:   "",
				Coordinate:   [...]float64{-29.3615343, -50.8514137},
			},
			wantErr: true,
		},
		{
			name: "AddressWithInvalidCity",
			fields: fields{
				ZipCode:      "89815-555",
				State:        enums.SC,
				City:         "",
				Neighborhood: "Desbravador",
				Street:       "Rua Balduino",
				Complement:   "",
				Coordinate:   [...]float64{-29.3615343, -50.8514137},
			},
			wantErr: true,
		},
		{
			name: "AddressWithInvalidNeighborhood",
			fields: fields{
				ZipCode:      "89815-555",
				State:        enums.SC,
				City:         "Chapecó",
				Neighborhood: "",
				Street:       "Rua Balduino",
				Complement:   "",
				Coordinate:   [...]float64{-29.3615343, -50.8514137},
			},
			wantErr: true,
		},
		{
			name: "AddressWithInvalidStreet",
			fields: fields{
				ZipCode:      "89815-555",
				State:        enums.SC,
				City:         "Chapecó",
				Neighborhood: "Desbravador",
				Street:       "",
				Complement:   "",
				Coordinate:   [...]float64{-29.3615343, -50.8514137},
			},
			wantErr: true,
		},
		{
			name: "AddressWithInvalidLatitude",
			fields: fields{
				ZipCode:      "89815-555",
				State:        enums.SC,
				City:         "Chapecó",
				Neighborhood: "Desbravador",
				Street:       "Rua Balduino Batistti",
				Complement:   "",
				Coordinate:   [...]float64{-91.3615343, -50.8514137},
			},
			wantErr: true,
		},
		{
			name: "AddressWithInvalidLongitude",
			fields: fields{
				ZipCode:      "89815-555",
				State:        enums.SC,
				City:         "Chapecó",
				Neighborhood: "Desbravador",
				Street:       "Rua Balduino Batistti",
				Complement:   "",
				Coordinate:   [...]float64{-29.3615343, 180.8514137},
			},
			wantErr: true,
		},
		{
			name: "ValidAddress",
			fields: fields{
				ZipCode:      "89815-555",
				State:        enums.SC,
				City:         "Chapecó",
				Neighborhood: "Desbravador",
				Street:       "Rua Balduino Batistti",
				Complement:   "",
				Coordinate:   [...]float64{-29.3615343, -50.8514137},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			addr := &models.Address{
				ZipCode:      tt.fields.ZipCode,
				State:        tt.fields.State,
				City:         tt.fields.City,
				Neighborhood: tt.fields.Neighborhood,
				Street:       tt.fields.Street,
				Complement:   tt.fields.Complement,
				Coordinates:  tt.fields.Coordinate,
			}

			var err error = addr.Validate()

			if err != nil {
				fmt.Printf("Test: %s, Erro: %s, Was expected?: %t\n", tt.name, err.Error(), tt.wantErr)
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("Address.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
