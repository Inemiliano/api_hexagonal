package controllers

import (
    "api/src/Users/application"
    "api/src/Users/domain"
    "api/src/Users/infraestructure"
    "encoding/json"
    "log"
    "net/http"

    
)
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
    log.Println("Método recibido:", r.Method)

    if r.Method != http.MethodPost {
        http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
        return
    }

    var user domain.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        log.Printf("Error al decodificar JSON: %v", err)
        http.Error(w, "Error al procesar el JSON", http.StatusBadRequest)
        return
    }

    log.Printf("Usuario recibido: %+v", user)

    if user.Name == "" || user.Email == "" || user.Password == "" {
        log.Println("Error: Nombre, email y contraseña son obligatorios")
        http.Error(w, "Error: Nombre, email y contraseña son obligatorios", http.StatusBadRequest)
        return
    }

    if len(user.Password) < 6 {
        log.Println("Error: La contraseña debe tener al menos 6 caracteres")
        http.Error(w, "Error: La contraseña debe tener al menos 6 caracteres", http.StatusBadRequest)
        return
    }

    repo := infraestructure.NewMongoUserRepository()
    if repo == nil {
        log.Println("Error: No se pudo inicializar el repositorio MongoDB")
        http.Error(w, "Error interno del servidor", http.StatusInternalServerError)
        return
    }

    log.Println("Repositorio MongoDB inicializado")

    useCase := application.NewCreateUser(repo)
    if err := useCase.Execute(user); err != nil {
        log.Printf("Error al ejecutar caso de uso: %v", err)
        http.Error(w, "Error al guardar el usuario en la base de datos", http.StatusInternalServerError)
        return
    }

    log.Println("Usuario insertado correctamente en MongoDB")
    w.WriteHeader(http.StatusCreated)
    w.Write([]byte("Usuario creado con éxito"))
}
