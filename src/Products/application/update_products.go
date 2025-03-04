package application

import (
	"api/src/Products/domain"
	
)

type UpdateProduct struct {
	repo domain.IProductRepository
}

func NewUpdateProduct(repo domain.IProductRepository) *UpdateProduct {
	return &UpdateProduct{repo: repo}
}

func (up *UpdateProduct) Execute(n string,nombre string, precio int16) error {
	
	product := &domain.Product{
		Nombre: nombre,
		Precio: precio,
	}

	// Llamar al método `Update` del repositorio
	return up.repo.Update(n, product)
}





