// controllers/create_product_handler.go
package controllers

import (
    "api/src/Products/application"
    "api/src/Products/infraestructure"
    "encoding/json"
    "net/http"
)

func DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodDelete {
        http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
        return
    }

    var requestData struct {
        ID string `json:"id"`
    }
    if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
        http.Error(w, "Error al procesar el JSON", http.StatusBadRequest)
        return
    }

    // Inicializamos el repositorio que implementa IProductRepository
    repo := infraestructure.NewMongos()
    if repo == nil {
        http.Error(w, "Error al inicializar el repositorio", http.StatusInternalServerError)
        return
    }

    // Creamos el caso de uso para eliminar el producto
    useCase := application.NewDeleteProduct(repo)

    // Ejecutamos el caso de uso para eliminar el producto
    if err := useCase.Execute(requestData.ID); err != nil {
        http.Error(w, "Error al eliminar el producto", http.StatusInternalServerError)
        return
    }

    // Respondemos con éxito si no hubo error
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Producto eliminado correctamente"))
}
