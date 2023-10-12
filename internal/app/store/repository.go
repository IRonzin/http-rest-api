package store

import "github.com/IRonzin/http-rest-api/internal/app/model"

// User repo interface
type UserRepository interface {
	Create(*model.User) error
	FindByEmail(string) (*model.User, error)
	Find(int) (*model.User, error)
}
