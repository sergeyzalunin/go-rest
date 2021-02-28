package models_test

import (
	"testing"

	"github.com/sergeyzalunin/go-rest/internal/app/models"
	"github.com/stretchr/testify/assert"
)

type validateTestCase struct {
	name    string
	u       func() *models.User
	wantErr bool
}

func TestUser_BeforeCreate(t *testing.T) {
	u := models.TestUser(t)
	assert.NoError(t, u.BeforeCreate())
	assert.NotEmpty(t, u.EncryptedPassword)
}

func TestUser_Validate(t *testing.T) {
	tests := getValidateTestCases(t)

	for _, tt := range tests {
		tc := tt
		t.Run(tc.name, func(t *testing.T) {
			tu := tc.u()

			if tc.wantErr {
				assert.Error(t, tu.Validate())
			} else {
				assert.NoError(t, tu.Validate())
			}
		})
	}
}

func getValidateTestCases(t *testing.T) []validateTestCase {
	t.Helper()

	return []validateTestCase{
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
}
