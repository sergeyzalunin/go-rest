package models_test

import (
	"testing"

	"github.com/sergeyzalunin/go-rest/internal/app/models"
	"github.com/stretchr/testify/assert"
)

func TestUser_BeforeCreate(t *testing.T) {
	t.Parallel()

	u := models.TestUser(t)
	assert.NoError(t, u.BeforeCreate())
	assert.NotEmpty(t, u.EncryptedPassword)
}

func TestUser_Validate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		u       func() *models.User
		wantErr bool
	}{
		{
			name: "valid",
			u: func() *models.User {
				return models.TestUser(t)
			},
			wantErr: false,
		},
		{
			name: "empty email",
			u: func() *models.User {
				tu := models.TestUser(t)
				tu.Email = ""

				return tu
			},
			wantErr: true,
		},
		{
			name: "empty password",
			u: func() *models.User {
				tu := models.TestUser(t)
				tu.Password = ""

				return tu
			},
			wantErr: true,
		},
		{
			name: "short password",
			u: func() *models.User {
				tu := models.TestUser(t)
				tu.Password = "12345"

				return tu
			},
			wantErr: true,
		},
		{
			name: "with encrypt password",
			u: func() *models.User {
				tu := models.TestUser(t)
				tu.Password = ""
				tu.EncryptedPassword = "skfjdkfjkdjfdkjfkdjf"

				return tu
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tu := tt.u()

			if tt.wantErr {
				assert.Error(t, tu.Validate())
			} else {
				assert.NoError(t, tu.Validate())
			}
		})
	}
}
