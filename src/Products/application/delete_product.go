package application

import (
	"api/src/Products/domain"
	"log"
)

type DeleteProduct struct {
	repo domain.IProductRepository
}

func NewDeleteProduct(repo domain.IProductRepository) *DeleteProduct {
	return &DeleteProduct{repo: repo}
}

func (dp *DeleteProduct) Execute(name string) error {
	log.Println("Eliminando producto con Nombre:", name)

	
	return dp.repo.Delete(name)
}

