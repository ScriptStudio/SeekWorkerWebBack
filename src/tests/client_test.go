package tests

import (
	"fmt"
	"testing"

	"github.com/ScriptStudio/SeekWorkerWebBack/src/models"
	"github.com/ScriptStudio/SeekWorkerWebBack/src/models/enums"
)

func TestClient_Validate(t *testing.T) {
	type fields struct {
		name         string
		phone        models.Telephone
		email        string
		addr         models.Address
		logo         string
		contacts     []models.ContactOnClient
		observations string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.

		{
			name: "ClientWithInvalidTelephone",
			fields: fields{
				name: "Renner",
				phone: models.Telephone{
					Ddd:    49,
					Number: 331155,
				},
				email: "matheus.saraiva@gmail.com",
				addr: models.Address{
					ZipCode:      "89815-500",
					State:        enums.SC,
					City:         "Chapecó",
					Neighborhood: "Desbravador",
					Street:       "Rua Bauduino",
					Complement:   "",
					Coordinates:  [...]float64{-29.3615343, -50.8514137},
				},
				logo: "",
				contacts: []models.ContactOnClient{
					{
						Name: "Joice",
						Phone: models.Telephone{
							Ddd:    49,
							Number: 999923554,
						},
						Whatsapp: models.Telephone{
							Ddd:    49,
							Number: 999923554,
						},
						Email: "ijoice2306@outlook.com",
						Skype: "ijoice2306@outlook.com",
					},
				},
				observations: "",
			},
			wantErr: true,
		},
		{
			name: "ClientWithInvalidAddress",
			fields: fields{
				name: "Renner",
				phone: models.Telephone{
					Ddd:    49,
					Number: 33115512,
				},
				email: "matheus.saraiva@gmail.com",
				addr: models.Address{
					ZipCode:      "89815-500",
					State:        enums.SC,
					City:         "",
					Neighborhood: "Desbravador",
					Street:       "Rua Bauduino",
					Complement:   "",
					Coordinates:  [...]float64{-29.3615343, -50.8514137},
				},
				logo: "",
				contacts: []models.ContactOnClient{
					{
						Name: "Joice",
						Phone: models.Telephone{
							Ddd:    49,
							Number: 999923554,
						},
						Whatsapp: models.Telephone{
							Ddd:    49,
							Number: 999923554,
						},
						Email: "ijoice2306@outlook.com",
						Skype: "ijoice2306@outlook.com",
					},
				},
				observations: "",
			},
			wantErr: true,
		},
		{
			name: "ClientWithInvalidContact",
			fields: fields{
				name: "Renner",
				phone: models.Telephone{
					Ddd:    49,
					Number: 33115511,
				},
				email: "matheus.saraiva@gmail.com",
				addr: models.Address{
					ZipCode:      "89815-500",
					State:        enums.SC,
					City:         "Chapecó",
					Neighborhood: "Desbravador",
					Street:       "Rua Bauduino",
					Complement:   "",
					Coordinates:  [...]float64{-29.3615343, -50.8514137},
				},
				logo: "",
				contacts: []models.ContactOnClient{
					{
						Name: "Joice999",
						Phone: models.Telephone{
							Ddd:    49,
							Number: 999923554,
						},
						Whatsapp: models.Telephone{
							Ddd:    49,
							Number: 999923554,
						},
						Email: "ijoice2306@outlook.com",
						Skype: "ijoice2306@outlook.com",
					},
				},
				observations: "",
			},
			wantErr: true,
		},
		{
			name: "ValidClient",
			fields: fields{
				name: "Renner",
				phone: models.Telephone{
					Ddd:    49,
					Number: 33115511,
				},
				email: "matheus.saraiva@gmail.com",
				addr: models.Address{
					ZipCode:      "89815-500",
					State:        enums.SC,
					City:         "Chapecó",
					Neighborhood: "Desbravador",
					Street:       "Rua Bauduino",
					Complement:   "",
					Coordinates:  [...]float64{-29.3615343, -50.8514137},
				},
				logo: "",
				contacts: []models.ContactOnClient{
					{
						Name: "Joice Saraiva",
						Phone: models.Telephone{
							Ddd:    49,
							Number: 999923554,
						},
						Whatsapp: models.Telephone{
							Ddd:    49,
							Number: 999923554,
						},
						Email: "ijoice2306@outlook.com",
						Skype: "ijoice2306@outlook.com",
					},
				},
				observations: "",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &models.Client{
				Name:         tt.fields.name,
				Phone:        tt.fields.phone,
				Email:        tt.fields.email,
				Addr:         tt.fields.addr,
				Logo:         tt.fields.logo,
				Contacts:     tt.fields.contacts,
				Observations: tt.fields.observations,
			}

			var err error = c.Validate()

			if err != nil {
				fmt.Printf("Test: %s, Erro: %s, Was expected?: %t\n", tt.name, err.Error(), tt.wantErr)
			}

			if err := c.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Client.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
