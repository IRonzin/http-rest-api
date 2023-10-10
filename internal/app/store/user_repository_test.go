package store_test

import (
	"testing"

	"github.com/IRonzin/http-rest-api/internal/app/model"
	"github.com/IRonzin/http-rest-api/internal/app/store"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseUrl)
	defer teardown("users")

	email := "user@example.org"
	pass := "encrypted"
	u, err := s.User().Create(&model.User{
		Email:             email,
		EncryptedPassword: pass,
	})

	assert.NoError(t, err)
	assert.NotNil(t, u)
	assert.NotEqual(t, 0, u.Id)
	assert.Equal(t, email, u.Email)
	assert.Equal(t, pass, u.EncryptedPassword)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, databaseUrl)
	defer teardown("users")

	email := "user@example.org"
	pass := "encrypted"
	_, err := s.User().FindByEmail(email)
	assert.Error(t, err)

	u, err := s.User().Create(&model.User{
		Email:             email,
		EncryptedPassword: pass,
	})

	us, er := s.User().FindByEmail(u.Email)
	assert.NoError(t, er)
	assert.NotNil(t, us)
	assert.NotEqual(t, 0, us.Id)
	assert.Equal(t, email, us.Email)
	assert.Equal(t, pass, us.EncryptedPassword)
}
