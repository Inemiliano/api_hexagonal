package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type IUsersRepository interface{
	Save(user *User) error
	GetAll() ([]User, error)
	Update(id primitive.ObjectID, user *User)error
	Delete(id primitive.ObjectID)error
}