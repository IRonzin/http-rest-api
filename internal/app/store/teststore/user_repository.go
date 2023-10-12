package teststore

import (
	"github.com/IRonzin/http-rest-api/internal/app/model"
	"github.com/IRonzin/http-rest-api/internal/app/store"
)

type UserRepository struct {
	store *Store
	users map[int]*model.User
}

// Create..
func (r *UserRepository) Create(u *model.User) error {

	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	u.Id = len(r.users) + 1
	r.users[u.Id] = u

	return nil
}

// Find by email...
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	for _, user := range r.users {
		if user.Email == email {
			return user, nil
		}
	}

	return nil, store.ErrRecordNotFound
}

// Find...
func (r *UserRepository) Find(id int) (*model.User, error) {

	if val, ok := r.users[id]; !ok {
		return nil, store.ErrRecordNotFound
	} else {
		return val, nil
	}
}
