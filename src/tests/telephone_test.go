package tests

import (
	"fmt"
	"testing"

	"github.com/ScriptStudio/SeekWorkerWebBack/src/models"
)

func TestTelephone_Validate(t *testing.T) {
	type fields struct {
		ddd    uint8
		number uint32
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "AddressWithInvalidDdd",
			fields: fields{
				ddd:    10,
				number: 988070350,
			},
			wantErr: true,
		},
		{
			name: "AddressWithInvalidNumber",
			fields: fields{
				ddd:    11,
				number: 988070,
			},
			wantErr: true,
		},
		{
			name: "ValidTelephone",
			fields: fields{
				ddd:    11,
				number: 988070350,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tel := &models.Telephone{
				Ddd:    tt.fields.ddd,
				Number: tt.fields.number,
			}

			var err error = tel.Validate()

			if err != nil {
				fmt.Printf("Test: %s, Erro: %s, Was expected?: %t\n", tt.name, err.Error(), tt.wantErr)
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("Telephone.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTelephone_FormatedTelephone(t *testing.T) {
	type fields struct {
		ddd    uint8
		number uint32
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
		{
			name: "ValidTelephone",
			fields: fields{
				ddd:    11,
				number: 988070350,
			},
			want: "(11) 98807-0350",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tel := &models.Telephone{
				Ddd:    tt.fields.ddd,
				Number: tt.fields.number,
			}
			if got := tel.FormatedTelephone(); got != tt.want {
				t.Errorf("Telephone.FormatedTelephone() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTelephone_CleanedTelephone(t *testing.T) {
	type fields struct {
		ddd    uint8
		number uint32
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
		{
			name: "ValidTelephone",
			fields: fields{
				ddd:    11,
				number: 988070350,
			},
			want: "11988070350",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tel := &models.Telephone{
				Ddd:    tt.fields.ddd,
				Number: tt.fields.number,
			}
			if got := tel.CleanedTelephone(); got != tt.want {
				t.Errorf("Telephone.CleanedTelephone() = %v, want %v", got, tt.want)
			}
		})
	}
}
