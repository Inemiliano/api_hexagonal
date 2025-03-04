package routes

import (
	"api/src/Users/infraestructure/controllers"
	"github.com/gin-gonic/gin"
)

var userUpdates = make(chan struct{}) 

func SetupRoutesUsers(r *gin.Engine) {
	r.POST("/users", controllers.CreateUserHandler)       
	r.GET("/getUsers", controllers.GetUsersHandler)        
	r.DELETE("/deleteUsers", controllers.DeleteUserHandler) 
	r.PUT("/updateUsers", controllers.UpdateUserHandler)   

}
