package application

import (
	"api/src/Products/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateProduct struct {
	repo domain.IProductRepository
}

func NewUpdateProduct(repo domain.IProductRepository) *UpdateProduct {
	return &UpdateProduct{repo: repo}
}

func (up *UpdateProduct) Execute(id primitive.ObjectID, nombre string, precio int16) error {
	
	product := &domain.Product{
		Nombre: nombre,
		Precio: precio,
	}

	
	return up.repo.Update(id, product)
}



