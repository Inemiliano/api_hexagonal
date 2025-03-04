// routes/products.go
package routes

import (
	"api/src/Products/infraestructure/controllers"
	"github.com/gin-gonic/gin"

)

// SetupRoutes configura las rutas
func SetupRoutes(r *gin.Engine) {
	

	r.POST("/products", controllers.CreateProductHandler)
	r.GET("/getProducts", controllers.GetProductsHandler)
	r.DELETE("/deleteProduct/:nombre", controllers.DeleteProductHandler)
	r.PUT("/updateProduct/:nombre", controllers.UpdateProductHandler)

}

