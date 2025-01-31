package controllers

import (
	"api/src/Products/application"
	"api/src/Products/infraestructure"
	"api/src/Products/domain" 
	"encoding/json"
	"log"
	"net/http"

	
)

func UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	// Decodificar el JSON del cuerpo de la solicitud
	var product domain.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Error al procesar el JSON", http.StatusBadRequest)
		return
	}

	// Verificar que el ID está presente en la solicitud
	if product.ID.IsZero() {
		http.Error(w, "ID del producto es requerido en el JSON", http.StatusBadRequest)
		return
	}

	// Crear una instancia del repositorio
	repo := infraestructure.NewMongos()

	// Llamar al caso de uso para actualizar el producto
	useCase := application.NewUpdateProduct(repo)
	err := useCase.Execute(product.ID, product.Nombre, product.Precio)
	if err != nil {
		log.Println("Error al actualizar el producto:", err)
		http.Error(w, "Error al actualizar el producto", http.StatusInternalServerError)
		return
	}

	// Responder con éxito
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Producto actualizado correctamente"))
}
