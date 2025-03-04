package routes

import (
	"api/src/Products/application/service"
	"api/src/Products/infraestructure/controllers"
	"github.com/gin-gonic/gin"
)


func SetupEvents(r *gin.Engine, eventService *services.EventService) {
	
	r.GET("/events", controllers.LongPollingHandler(eventService))
}
