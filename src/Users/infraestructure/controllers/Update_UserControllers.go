package controllers

import (
	"api/src/Users/domain"
	"api/src/Users/infraestructure"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

// UpdateUserHandler maneja la actualización de usuarios con Gin
func UpdateUserHandler(c *gin.Context) {
	// Comprobamos que el método sea PUT
	if c.Request.Method != http.MethodPut {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Método no permitido"})
		return
	}

	// Definimos la estructura para recibir el JSON de la petición
	var requestData struct {
		ID     string `json:"id"`
		Nombre string `json:"nombre"`
		Email  string `json:"email"`
	}

	// Decodificamos el cuerpo de la petición JSON
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al procesar el JSON"})
		return
	}

	// Convertimos el ID de string a ObjectID
	objectID, err := primitive.ObjectIDFromHex(requestData.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	// Creamos la instancia del usuario con los datos actualizados
	user := domain.User{
		ID:     objectID,
		Name:   requestData.Nombre,
		Email:  requestData.Email,
	}

	// Creamos el repositorio e intentamos actualizar el usuario
	repo := infraestructure.NewMongoUserRepository()
	if err := repo.Update(user.ID, &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar el usuario"})
		return
	}

	// Respondemos con un mensaje de éxito
	c.JSON(http.StatusOK, gin.H{"message": "Usuario actualizado correctamente"})
}
