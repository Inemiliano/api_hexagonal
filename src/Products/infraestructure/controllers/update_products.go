package controllers

import (
	"api/src/Products/application"
	"api/src/Products/infraestructure"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// UpdateProductHandler maneja la actualización de productos usando Gin
func UpdateProductHandler(c *gin.Context) {
	log.Println("Recibiendo solicitud PUT para actualizar producto")

	nombre := c.Param("nombre")
	if nombre == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nombre del producto requerido"})
		return
	}
	// Estructura para recibir los datos del JSON
	var requestData struct {
		ID     int    `json:"id"`
		Nombre string `json:"nombre"`
		Precio int16  `json:"precio"`
	}

	// Decodificar JSON
	if err := c.ShouldBindJSON(&requestData); err != nil {
		log.Printf("Error al procesar el JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al procesar el JSON"})
		return
	}

	// Crear instancia del repositorio
	repo := infraestructure.NewMongos()
	if repo == nil {
		log.Println("Error al inicializar el repositorio")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error interno del servidor"})
		return
	}

	// Llamar al caso de uso para actualizar
	useCase := application.NewUpdateProduct(repo)
	err := useCase.Execute(nombre,requestData.Nombre, requestData.Precio)
	if err != nil {
		log.Printf("Error al actualizar el producto: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar el producto"})
		return
	}

	// Responder con éxito
	c.JSON(http.StatusOK, gin.H{"message": "Producto actualizado correctamente"})
}
