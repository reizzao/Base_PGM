package repositories

import (
	t "githubendereco/entitys"
)

type UserRepository struct {
	Users []*t.User
}

func (r UserRepository) Save(data t.User) t.User {
	_ = append(r.Users, &data)
	return data
}

func (r UserRepository) List() []*t.User {
	return r.Users
}

//todo : tipar metodo e criar repo memory em implementations/
