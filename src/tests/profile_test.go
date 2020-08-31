package tests

import (
	"fmt"
	"testing"

	"github.com/ScriptStudio/SeekWorkerWebBack/src/models"
	"github.com/ScriptStudio/SeekWorkerWebBack/src/models/enums"
)

func TestProfile_Validate(t *testing.T) {
	type fields struct {
		name        string
		permissions map[enums.FeatureSystem]bool
		description string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "ProfileWithEmptyPermissionsList",
			fields: fields{
				name:        "TestProfileEmptyList",
				permissions: make(map[enums.FeatureSystem]bool),
				description: "Any description",
			},
			wantErr: true,
		},
		{
			name: "ProfileWithOneFeatureSystem",
			fields: fields{
				name: "TestProfileWithOneFeatureInvaid",
				permissions: map[enums.FeatureSystem]bool{
					100:               false,
					enums.ClouseCalls: true,
				},
				description: "Any Description",
			},
			wantErr: true,
		},
		{
			name: "VilidProfile",
			fields: fields{
				name: "Operator",
				permissions: map[enums.FeatureSystem]bool{
					enums.ChangeProfile: false,
					enums.ClouseCalls:   true,
				},
				description: "Any Description",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &models.Profile{
				Name:        tt.fields.name,
				Permissions: tt.fields.permissions,
				Description: tt.fields.description,
			}

			var err error = p.Validate()

			if err != nil {
				fmt.Printf("Test: %s, Erro: %s, Was expected?: %t\n", tt.name, err.Error(), tt.wantErr)
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("Profile.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestProfile_SetOnePermission(t *testing.T) {
	type fields struct {
		name        string
		permissions map[enums.FeatureSystem]bool
		description string
	}
	type args struct {
		fs         enums.FeatureSystem
		permission bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "SetInvalidPermission",
			fields: fields{
				name: "Operator",
				permissions: map[enums.FeatureSystem]bool{
					enums.ChangeProfile: false,
					enums.ClouseCalls:   true,
				},
				description: "Any Description",
			},
			args:    args{200, false},
			wantErr: true,
		},
		{
			name: "ValidPermission",
			fields: fields{
				name: "Operator",
				permissions: map[enums.FeatureSystem]bool{
					enums.ChangeProfile: false,
					enums.ClouseCalls:   true,
				},
				description: "Any Description",
			},
			args:    args{enums.ChangeProfile, true},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &models.Profile{
				Name:        tt.fields.name,
				Permissions: tt.fields.permissions,
				Description: tt.fields.description,
			}

			var err error = p.SetOnePermission(tt.args.fs, tt.args.permission)

			if err != nil {
				fmt.Printf("Test: %s, Erro: %s, Was expected?: %t\n", tt.name, err.Error(), tt.wantErr)
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("Profile.SetOnePermission() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestProfile_GetOnePermission(t *testing.T) {
	type fields struct {
		name        string
		permissions map[enums.FeatureSystem]bool
		description string
	}
	type args struct {
		fs enums.FeatureSystem
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "GetNoExistsFeature",
			fields: fields{
				name: "Operator",
				permissions: map[enums.FeatureSystem]bool{
					enums.ChangeProfile: false,
					enums.ClouseCalls:   true,
				},
				description: "Any Description",
			},
			args:    args{200},
			want:    false,
			wantErr: true,
		},
		{
			name: "GetValidFeature",
			fields: fields{
				name: "Operator",
				permissions: map[enums.FeatureSystem]bool{
					enums.ChangeProfile: false,
					enums.ClouseCalls:   true,
				},
				description: "Any Description",
			},
			args:    args{enums.ClouseCalls},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &models.Profile{
				Name:        tt.fields.name,
				Permissions: tt.fields.permissions,
				Description: tt.fields.description,
			}
			got, err := p.GetOnePermission(tt.args.fs)

			if err != nil {
				fmt.Printf("Test: %s, Erro: %s, Was expected?: %t\n", tt.name, err.Error(), tt.wantErr)
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("Profile.GetOnePermission() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Profile.GetOnePermission() = %v, want %v", got, tt.want)
			}
		})
	}
}
