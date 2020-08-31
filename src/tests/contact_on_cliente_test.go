package tests

import (
	"fmt"
	"testing"

	"github.com/ScriptStudio/SeekWorkerWebBack/src/models"
)

func TestContactOnClient_Validate(t *testing.T) {
	type fields struct {
		name     string
		phone    models.Telephone
		whatsapp models.Telephone
		email    string
		skype    string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "ContactOnClientWithInvalidName",
			fields: fields{
				name: "Joice999",
				phone: models.Telephone{
					Ddd:    49,
					Number: 988070350,
				},
				whatsapp: models.Telephone{
					Ddd:    49,
					Number: 988070350,
				},
				email: "matheus.saraiva@gmail.com",
				skype: "matheus.saraiva@gmail.com",
			},
			wantErr: true,
		},
		{
			name: "ContactOnClientWithInvalidTelephone",
			fields: fields{
				name: "Matheus Saraiva",
				phone: models.Telephone{
					Ddd:    100,
					Number: 98807,
				},
				whatsapp: models.Telephone{
					Ddd:    49,
					Number: 988070350,
				},
				email: "matheus.saraiva@gmail.com",
				skype: "matheus.saraiva@gmail.com",
			},
			wantErr: true,
		},
		{
			name: "ContactOnClientWithInvalidWhatsapp",
			fields: fields{
				name: "Matheus Saraiva",
				phone: models.Telephone{
					Ddd:    49,
					Number: 988070350,
				},
				whatsapp: models.Telephone{
					Ddd:    49,
					Number: 9880703,
				},
				email: "matheus.saraiva@gmail.com",
				skype: "matheus.saraiva@gmail.com",
			},
			wantErr: true,
		},
		{
			name: "ValidContactOnClient",
			fields: fields{
				name: "Matheus Saraiva",
				phone: models.Telephone{
					Ddd:    49,
					Number: 988070350,
				},
				whatsapp: models.Telephone{
					Ddd:    49,
					Number: 988070350,
				},
				email: "matheus.saraiva@gmail.com",
				skype: "matheus.saraiva@gmail.com",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			coc := &models.ContactOnClient{
				Name:     tt.fields.name,
				Phone:    tt.fields.phone,
				Whatsapp: tt.fields.whatsapp,
				Email:    tt.fields.email,
				Skype:    tt.fields.skype,
			}
			var err error = coc.Validate()

			if err != nil {
				fmt.Printf("Test: %s, Erro: %s, Was expected?: %t\n", tt.name, err.Error(), tt.wantErr)
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("ContactOnClient.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
