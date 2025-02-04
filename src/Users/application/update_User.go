package application

import (
	"api/src/Users/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateUser struct {
	repo domain.IUsersRepository
}

func NewUpdateUser(repo domain.IUsersRepository) *UpdateUser {
	return &UpdateUser{repo: repo}
}

func (uu *UpdateUser) Execute(id primitive.ObjectID, user domain.User) error {
	return uu.repo.Update(id, &user)
}