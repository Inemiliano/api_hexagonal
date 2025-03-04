package main

import (
	"api/src/core"
	"api/src/Products/application/service"
	Products "api/src/Products/infraestructure/routes"
	Users "api/src/Users/infraestructure/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func handler(c *gin.Context) {
	client := core.GetMongoClient()
	databases, err := client.ListDatabaseNames(c, nil)
	if err != nil {
		log.Println("Error al obtener bases de datos:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al conectar con la base de datos"})
		return
	}

	log.Printf("Bases de datos disponibles: %v", databases)
	c.JSON(http.StatusOK, gin.H{"message": "Conexión exitosa a MongoDB"})
}

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	
	eventService := services.NewEventService()

	
	Products.SetupRoutes(r)
	Users.SetupRoutesUsers(r) 


	Products.SetupEvents(r, eventService)
	

	r.GET("/testMongo", handler)

	go func() {
		for {
			time.Sleep(10 * time.Second)
		}
	}()

	log.Println("Servidor escuchando en el puerto 8080...")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
