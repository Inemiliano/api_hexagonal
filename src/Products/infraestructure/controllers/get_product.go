package controllers

import (
	"api/src/Products/application"
	"api/src/Products/infraestructure"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

// GetProductsHandler maneja la obtención de productos usando Long Polling con ciclo de peticiones y timeout de 10 segundos
func GetProductsHandler(c *gin.Context) {
	log.Println("Recibiendo solicitud GET para obtener productos con Long Polling y ciclo de peticiones")

	// Inicializar repositorio
	repo := infraestructure.NewMongos()
	if repo == nil {
		log.Println("Error al inicializar el repositorio")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error interno del servidor"})
		return
	}

	useCase := application.NewGetProduct(repo)

	// Canal para recibir cambios de productos
	updates := make(chan []interface{}, 1)

	// Timeout de 10 segundos para la conexión
	timeout := time.After(10 * time.Second)

	// Goroutine para obtener productos con un ciclo de peticiones durante el tiempo de espera
	go func() {
		startTime := time.Now()

		// Mientras no hayan pasado 10 segundos
		for time.Since(startTime) < 10*time.Second {
			products, err := useCase.Execute()
			if err != nil {
				log.Printf("Error al obtener los productos: %v", err)
				close(updates) // Cerramos el canal si hay un error
				return
			}

			// Convertimos []domain.Product a []interface{}
			var result []interface{}
			for _, p := range products {
				result = append(result, p)
			}

			// Si hay productos nuevos, los enviamos al canal y terminamos
			if len(result) > 0 {
				updates <- result
				return
			}

			// Esperamos 1 segundo antes de intentar de nuevo
			time.Sleep(1 * time.Second)
		}

		// Si después de 10 segundos no se encontraron productos, enviamos el resultado vacío
		updates <- nil
	}()

	// Esperamos hasta recibir productos o el timeout
	select {
	case products := <-updates:
		// Si se reciben productos, los enviamos
		if products != nil {
			c.JSON(http.StatusOK, products)
		} else {
			// Si no se encontraron productos después del ciclo de 10 segundos
			c.JSON(http.StatusNoContent, gin.H{"message": "No hay productos nuevos"})
		}
	case <-timeout:
		// Si el tiempo de espera de 10 segundos se acaba, respondemos con un timeout
		log.Println("Long polling timeout alcanzado después de 10 segundos")
		c.JSON(http.StatusNoContent, gin.H{"message": "No hay productos nuevos"})
	}
}
