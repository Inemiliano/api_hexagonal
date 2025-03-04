package controllers

import (
	"api/src/Products/application"
	"api/src/Products/domain"
	"api/src/Products/infraestructure"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// CreateProductHandler maneja la creación de productos usando Gin
func CreateProductHandler(c *gin.Context) {
	log.Println("Método recibido:", c.Request.Method)

	var product domain.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		log.Printf("Error al decodificar JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al procesar el JSON"})
		return
	}

	log.Printf("Producto recibido: %+v", product)

	// Validaciones
	if product.Nombre == "" || product.Precio <= 0 {
		log.Println("Error: Nombre y precio son obligatorios")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nombre y precio son obligatorios"})
		return
	}

	if len(product.Nombre) < 3 {
		log.Println("Error: Nombre debe tener al menos 3 caracteres")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nombre debe tener al menos 3 caracteres"})
		return
	}

	// Crear instancia del repositorio MongoDB
	repo := infraestructure.NewMongos()
	if repo == nil {
		log.Println("Error: No se pudo inicializar el repositorio MongoDB")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error interno del servidor"})
		return
	}

	log.Println("Repositorio MongoDB inicializado")

	// Ejecutar el caso de uso
	useCase := application.NewCreateProduct(repo)
	if err := useCase.Execute(product); err != nil {
		log.Printf("Error al ejecutar caso de uso: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar el producto en la base de datos"})
		return
	}

	log.Println("Producto insertado correctamente en MongoDB")

	// Responder con JSON
	c.JSON(http.StatusCreated, gin.H{
		"message": "Producto creado con éxito",
		"product": product,
	})
}
