package sqlstore_test

import (
	"testing"

	"github.com/IRonzin/http-rest-api/internal/app/model"
	"github.com/IRonzin/http-rest-api/internal/app/store"
	"github.com/IRonzin/http-rest-api/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseUrl)
	defer teardown("users")
	s := sqlstore.New(db)

	email := "user@example.org"
	pass := "password"
	testUser := model.TestUser(t)
	testUser.Email = email
	testUser.Password = pass
	err := s.User().Create(testUser)

	assert.NoError(t, err)
	assert.NotNil(t, testUser)
	assert.NotEqual(t, 0, testUser.Id)
	assert.Equal(t, email, testUser.Email)
	assert.NotEmpty(t, testUser.EncryptedPassword)

	testUser2 := model.TestUser(t)
	testUser2.Password = "short"
	err = s.User().Create(testUser2)
	assert.Error(t, err)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseUrl)
	defer teardown("users")
	s := sqlstore.New(db)

	email := "user@example.org"
	pass := "password"
	_, err := s.User().FindByEmail(email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	testUser := model.TestUser(t)
	testUser.Email = email
	testUser.Password = pass

	err = s.User().Create(testUser)

	us, er := s.User().FindByEmail(testUser.Email)
	assert.NoError(t, er)
	assert.NotNil(t, us)
	assert.NotEqual(t, 0, us.Id)
	assert.Equal(t, email, us.Email)
	assert.NotEmpty(t, us.EncryptedPassword)
}

func TestUserRepository_Find(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseUrl)
	defer teardown("users")
	s := sqlstore.New(db)

	email := "user@example.org"
	pass := "password"
	_, err := s.User().Find(0)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	testUser := model.TestUser(t)
	testUser.Email = email
	testUser.Password = pass

	err = s.User().Create(testUser)

	us, er := s.User().Find(testUser.Id)
	assert.NoError(t, er)
	assert.NotNil(t, us)
	assert.NotEqual(t, 0, us.Id)
	assert.Equal(t, email, us.Email)
	assert.NotEmpty(t, us.EncryptedPassword)
}
