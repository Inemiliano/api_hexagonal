package controllers

import (
    "api/src/Users/application"
    "api/src/Users/infraestructure"
     "go.mongodb.org/mongo-driver/bson/primitive"
    "log"
    "net/http"
)

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
    log.Println("Método recibido:", r.Method)

    if r.Method != http.MethodDelete {
        http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
        return
    }

    // Obtener el ID del usuario desde la URL
    idStr := r.URL.Query().Get("id")
    if idStr == "" {
        http.Error(w, "ID del usuario es requerido", http.StatusBadRequest)
        return
    }

    // Convertir el string ID a primitive.ObjectID
    id, err := primitive.ObjectIDFromHex(idStr)
    if err != nil {
        http.Error(w, "ID no válido", http.StatusBadRequest)
        return
    }

    // Crear una instancia del repositorio
    repo := infraestructure.NewMongoUserRepository()

    // Llamar al caso de uso para eliminar el usuario
    useCase := application.NewDeleteUser(repo)
    if err := useCase.Execute(id); err != nil {
        log.Printf("Error al eliminar el usuario: %v", err)
        http.Error(w, "Error al eliminar el usuario", http.StatusInternalServerError)
        return
    }

    log.Println("Usuario eliminado correctamente")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Usuario eliminado correctamente"))
}