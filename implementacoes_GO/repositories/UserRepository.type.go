package repositories

import (
	t "githubendereco/entitys"
)

type UserRepository struct {
	Users []*t.User
}

type IUserRepository interface {
	// Users []*t.User
	Save(data t.User) t.User
	// List() []*t.User
}
