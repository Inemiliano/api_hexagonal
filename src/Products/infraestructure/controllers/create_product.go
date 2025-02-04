// controllers/create_product_handler.go
package controllers

import (
    "api/src/Products/application"
    "api/src/Products/domain"
    "api/src/Products/infraestructure"
    "encoding/json"
    "log"
    "net/http"
)

func CreateProductHandler(w http.ResponseWriter, r *http.Request) {
    log.Println("Método recibido:", r.Method)

    if r.Method != http.MethodPost {
        http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
        return
    }

    var product domain.Product
    if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
        log.Printf("Error al decodificar JSON: %v", err)
        http.Error(w, "Error al procesar el JSON", http.StatusBadRequest)
        return
    }

    log.Printf("Producto recibido: %+v", product)

    if product.Nombre == "" || product.Precio <= 0 {
        log.Println("Error: Nombre y precio son obligatorios")
        http.Error(w, "Error: Nombre y precio son obligatorios", http.StatusBadRequest)
        return
    }

    if len(product.Nombre) < 3 {
        log.Println("Error: Nombre debe tener al menos 3 caracteres")
        http.Error(w, "Error: Nombre debe tener al menos 3 caracteres", http.StatusBadRequest)
        return
    }

    // Crear una instancia del repositorio MongoDB que implementa IProductRepository
    repo := infraestructure.NewMongos()
    if repo == nil {
        log.Println("Error: No se pudo inicializar el repositorio MongoDB")
        http.Error(w, "Error interno del servidor", http.StatusInternalServerError)
        return
    }

    log.Println("Repositorio MongoDB inicializado")

   
    useCase := application.NewCreateProduct(repo)

    if err := useCase.Execute(product); err != nil {
        log.Printf("Error al ejecutar caso de uso: %v", err)
        http.Error(w, "Error al guardar el producto en la base de datos", http.StatusInternalServerError)
        return
    }

    log.Println("Producto insertado correctamente en MongoDB")
    w.WriteHeader(http.StatusCreated)
    w.Write([]byte("Producto creado con éxito"))
}
