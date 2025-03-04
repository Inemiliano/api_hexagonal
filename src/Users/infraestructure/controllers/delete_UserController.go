package controllers

import (
	"api/src/Users/application"
	"api/src/Users/infraestructure"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
)

// DeleteUserHandler maneja la eliminación de usuarios con Gin
func DeleteUserHandler(c *gin.Context) {
	log.Println("Método recibido: DELETE")

	// Obtener el ID del usuario desde la URL
	idStr := c.Query("id")
	if idStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID del usuario es requerido"})
		return
	}

	// Convertir el string ID a primitive.ObjectID
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID no válido"})
		return
	}

	// Crear una instancia del repositorio
	repo := infraestructure.NewMongoUserRepository()

	// Llamar al caso de uso para eliminar el usuario
	useCase := application.NewDeleteUser(repo)
	if err := useCase.Execute(id); err != nil {
		log.Printf("Error al eliminar el usuario: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar el usuario"})
		return
	}

	log.Println("Usuario eliminado correctamente")
	c.JSON(http.StatusOK, gin.H{"message": "Usuario eliminado correctamente"})
}
