package controllers

import (
	"api/src/Users/domain"
	"api/src/Users/infraestructure"
	"encoding/json"
	"net/http"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPut {
        http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
        return
    }

    var requestData struct {
        ID    string `json:"id"`
        Nombre string `json:"nombre"`
        Email  string `json:"email"`
    }
    if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
        http.Error(w, "Error al procesar el JSON", http.StatusBadRequest)
        return
    }

    objectID, err := primitive.ObjectIDFromHex(requestData.ID)
    if err != nil {
        http.Error(w, "ID inválido", http.StatusBadRequest)
        return
    }

    user := domain.User{
        ID:     objectID,
        Name: requestData.Nombre,
        Email:  requestData.Email,
    }

    repo := infraestructure.NewMongoUserRepository()
    err = repo.Update(user.ID, &user)
    if err != nil {
        http.Error(w, "Error al actualizar el usuario", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Usuario actualizado correctamente"))
}