package controllers

import (
	"api/src/Users/application"
	"api/src/Users/infraestructure"
	"github.com/gin-gonic/gin"
	"net/http"
	"encoding/json"  
	"time"
)

// Canal para notificar cambios en los usuarios
var userUpdateChannel = make(chan []byte, 1)

// Función para notificar cambios en los usuarios
func NotifyUserChange(data interface{}) {
	// Convierte los datos a JSON
	jsonData, _ := json.Marshal(data)
	select {
	case userUpdateChannel <- jsonData:
	default:
		// Si el canal está lleno, no se envía
	}
}

// GetUsersHandler maneja la obtención de usuarios con Long Polling
func GetUsersHandler(c *gin.Context) {
	repo := infraestructure.NewMongoUserRepository()
	useCase := application.NewGetUsers(repo)

	// Canal para esperar por cambios
	changes := make(chan bool)

	// Goroutine que simula la espera de cambios en la base de datos
	go func() {
		for i := 0; i < 30; i++ { // Espera hasta 30 segundos
			users, err := useCase.Execute()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo obtener los usuarios"})
				return
			}

			if len(users) > 0 { // Si se detectan usuarios, notificar el cambio
				changes <- true
				return
			}

			time.Sleep(1 * time.Second) // Espera 1 segundo antes de volver a verificar
		}
		changes <- false // No hubo cambios después de 30 segundos
	}()

	// Select que espera un cambio o el tiempo de espera
	select {
	case updated := <-changes:
		if updated {
			// Si hubo un cambio, devolver la lista de usuarios
			users, err := useCase.Execute()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo obtener los usuarios"})
				return
			}
			c.JSON(http.StatusOK, users)
		} else {
			// Si no hubo cambios en 30 segundos, devolver 204 No Content
			c.JSON(http.StatusNoContent, nil)
		}
	case <-time.After(30 * time.Second):
		// Timeout después de 30 segundos sin cambios
		c.JSON(http.StatusNoContent, nil)
	}
}
