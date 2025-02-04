package routes

import (
	"api/src/Products/infraestructure/controllers"
	"net/http"
)

func SetupRoutes() {
	http.HandleFunc("/products", controllers.CreateProductHandler)
	http.HandleFunc("/getProducts", controllers.GetProductsHandler)
	http.HandleFunc("/deleteProduct", controllers.DeleteProductHandler)
	http.HandleFunc("/updateProduct", controllers.UpdateProductHandler)

	http.HandleFunc("/longP", controllers.LongPollingHandler)
}
