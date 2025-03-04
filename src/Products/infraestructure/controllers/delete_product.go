// controllers/delete_product_handler.go
package controllers

import (
	"api/src/Products/infraestructure"
	"github.com/gin-gonic/gin"
	"net/http"
)

// DeleteProductHandler maneja la eliminación de un producto basado en su nombre
func DeleteProductHandler(c *gin.Context) {
	nombre := c.Param("nombre")
	if nombre == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nombre del producto requerido"})
		return
	}

	repo := infraestructure.NewMongos()
	if repo == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al inicializar el repositorio"})
		return
	}

	if err := repo.Delete(nombre); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar el producto"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Producto eliminado correctamente"})
}
