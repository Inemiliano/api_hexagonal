package controllers

import (
	"api/src/Users/application"
	"api/src/Users/infraestructure"
	"encoding/json"
	"net/http"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	
	repo := infraestructure.NewMongoUserRepository()
	useCase := application.NewGetUsers(repo)

	// Ejecutamos el caso de uso para obtener la lista de usuarios
	users, err := useCase.Execute()
	if err != nil {
		http.Error(w, "Error al obtener los usuarios", http.StatusInternalServerError)
		return
	}

	// Configuramos la respuesta con JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}
