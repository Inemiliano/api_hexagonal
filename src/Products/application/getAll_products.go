// application/get_product.go
package application

import (
	"api/src/Products/domain"
)

type GetProduct struct {
	repo domain.IProductRepository
}

func NewGetProduct(repo domain.IProductRepository) *GetProduct {
	return &GetProduct{repo: repo}
}

func (gp *GetProduct) Execute() ([]domain.Product, error) {
	return gp.repo.GetAll()
}
