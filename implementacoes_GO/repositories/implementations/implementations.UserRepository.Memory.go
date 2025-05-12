package implementations

import (
	t "githubendereco/entitys"
	// rp "githubendereco/repositories"
)

func UserRepositoryMemory() rp.UserRepository {
	res := rp.UserRepository{
		Users: []*t.User{},
	}
	return res
}

func Save(data t.User) t.User {
	// _ = append(rp.Users, &data)
	return data
}

func List() []*t.User {
	// return rp.Users
}

//todo : tipar metodo e criar repo memory em implementations/
