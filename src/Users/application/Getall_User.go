package application

import "api/src/Users/domain"

type GetUsers struct {
	repo domain.IUsersRepository
}

func NewGetUsers(repo domain.IUsersRepository) *GetUsers {
	return &GetUsers{repo: repo}
}

func (gu *GetUsers) Execute() ([]domain.User, error) {
	return gu.repo.GetAll()
}