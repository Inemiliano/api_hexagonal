package routes

import (
	"api/src/Users/infraestructure/controllers"
	"net/http"
)

func SetupRoutesUsers(){
	http.HandleFunc("/users", controllers.CreateUserHandler)
	http.HandleFunc("/getUsers", controllers.GetUsersHandler)
	http.HandleFunc("/deleteUsers", controllers.DeleteUserHandler)
	http.HandleFunc("/updateUsers", controllers.UpdateUserHandler)

}