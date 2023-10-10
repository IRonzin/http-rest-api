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
	pass := "password"
	testUser := model.TestUser(t)
	testUser.Email = email
	testUser.Password = pass
	u, err := s.User().Create(testUser)

	assert.NoError(t, err)
	assert.NotNil(t, u)
	assert.NotEqual(t, 0, u.Id)
	assert.Equal(t, email, u.Email)
	assert.NotEmpty(t, u.EncryptedPassword)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, databaseUrl)
	defer teardown("users")

	email := "user@example.org"
	pass := "password"
	_, err := s.User().FindByEmail(email)
	assert.Error(t, err)

	testUser := model.TestUser(t)
	testUser.Email = email
	testUser.Password = pass

	u, err := s.User().Create(testUser)

	us, er := s.User().FindByEmail(u.Email)
	assert.NoError(t, er)
	assert.NotNil(t, us)
	assert.NotEqual(t, 0, us.Id)
	assert.Equal(t, email, us.Email)
	assert.NotEmpty(t, us.EncryptedPassword)
}
