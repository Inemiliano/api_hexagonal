// application/create_product.go
package application

import "api/src/Products/domain"

type CreateProduct struct {
    repo domain.IProductRepository
}

func NewCreateProduct(repo domain.IProductRepository) *CreateProduct {
    return &CreateProduct{repo: repo}
}

func (c *CreateProduct) Execute(p domain.Product) error {
  
    err := c.repo.Save(&p)
    if err != nil {
        return err
    }
    return nil
}
