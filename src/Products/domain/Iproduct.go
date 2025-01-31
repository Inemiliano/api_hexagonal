package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type IProductRepository interface {
    Save(P *Product) error
    GetAll() ([]Product, error)
    Delete(id primitive.ObjectID) error  
    Update(ID primitive.ObjectID, p *Product) error  
}