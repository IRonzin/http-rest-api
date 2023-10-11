package teststore

import (
	"github.com/IRonzin/http-rest-api/internal/app/model"
	"github.com/IRonzin/http-rest-api/internal/app/store"
)

type UserRepository struct {
	store *Store
	users map[string]*model.User
}

// Create..
func (r *UserRepository) Create(u *model.User) error {

	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	r.users[u.Email] = u
	u.Id = len(r.users)

	return nil
}

// Find...
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {

	if val, ok := r.users[email]; !ok {
		return nil, store.ErrRecordNotFound
	} else {
		return val, nil
	}
}
