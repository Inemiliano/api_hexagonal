package application

import (
	"api/src/Users/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DeleteUser struct {
	repo domain.IUsersRepository
}

func NewDeleteUser(repo domain.IUsersRepository) *DeleteUser {
	return &DeleteUser{repo: repo}
}

func (du *DeleteUser) Execute(id primitive.ObjectID) error {
	return du.repo.Delete(id)
}