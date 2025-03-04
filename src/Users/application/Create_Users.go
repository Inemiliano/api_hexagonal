package application

import "api/src/Users/domain"

type CreatUser struct {
	repo domain.IUsersRepository
}

func NewCreateUser(repo domain.IUsersRepository) *CreatUser {
	return &CreatUser{repo: repo}
}

func (cu *CreatUser) Execute(user domain.User) error {
	return cu.repo.Save(&user)
}
