package controllers

import (
	"api/src/Products/application/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"context"
)

// LongPollingHandler maneja la espera de productos
func LongPollingHandler(es *services.EventService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Configurar el contexto para el long polling (con un timeout de 30 segundos)
		ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
		defer cancel()

		// Llamar al servicio para esperar productos
		products, ok := es.WaitForProducts(ctx)
		if !ok {
			c.JSON(http.StatusOK, gin.H{"message": "No hay productos disponibles o timeout alcanzado"})
			return
		}

		// Enviar los productos al cliente
		c.JSON(http.StatusOK, products)
	}
}
