package controllers

import (
	"api/src/Users/application"
	"api/src/Users/domain"
	"api/src/Users/infraestructure"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	
)


func CreateUserHandler(c *gin.Context) {
	log.Println("Método recibido: POST")

	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Printf("Error al decodificar JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al procesar el JSON"})
		return
	}

	log.Printf("Usuario recibido: %+v", user)

	if user.Name == "" || user.Email == "" || user.Password == "" {
		log.Println("Error: Nombre, email y contraseña son obligatorios")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nombre, email y contraseña son obligatorios"})
		return
	}

	if len(user.Password) < 6 {
		log.Println("Error: La contraseña debe tener al menos 6 caracteres")
		c.JSON(http.StatusBadRequest, gin.H{"error": "La contraseña debe tener al menos 6 caracteres"})
		return
	}

	repo := infraestructure.NewMongoUserRepository()
	if repo == nil {
		log.Println("Error: No se pudo inicializar el repositorio MongoDB")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error interno del servidor"})
		return
	}

	log.Println("Repositorio MongoDB inicializado")

	useCase := application.NewCreateUser(repo)
	if err := useCase.Execute(user); err != nil {
		log.Printf("Error al ejecutar caso de uso: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar el usuario en la base de datos"})
		return
	}

	log.Println("Usuario insertado correctamente en MongoDB")
}