package controllers

import (
	"api/src/Products/application"
	"api/src/Products/infraestructure"
	"encoding/json"
	"net/http"
)

func GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Metodo no permitido, ", http.StatusMethodNotAllowed)
		return
	}

	repo := infraestructure.NewMongos()
	useCase := application.NewGetProduct(repo)

	products, err := useCase.Execute()
	if err != nil {
		http.Error(w, "Error al obtener los productos ", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}
