package application

import (
	"api/src/Products/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

type DeleteProduct struct {
	repo domain.IProductRepository
}

func NewDeleteProduct(repo domain.IProductRepository) *DeleteProduct {
	return &DeleteProduct{repo: repo}
}

func (dp *DeleteProduct) Execute(id string) error {
	// Convertir el ID a ObjectID
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Error al convertir ID a ObjectID:", err)
		return err
	}

	
	return dp.repo.Delete(objID)
}


