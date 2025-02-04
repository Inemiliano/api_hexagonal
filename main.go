package main

import (
	"api/src/core"
	"api/src/Users/infraestructure/routes" 

	Products "api/src/Products/infraestructure/routes" 
	
	
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	client := core.GetMongoClient()
	
	// Obtener la lista de bases de datos para asegurarse de que la conexión con MongoDB funciona
	databases, err := client.ListDatabaseNames(r.Context(), nil)
	if err != nil {
		log.Println("Error al obtener bases de datos:", err)
		http.Error(w, "Error al conectar con la base de datos", http.StatusInternalServerError)
		return
	}

	log.Printf("Bases de datos disponibles: %v", databases)
	w.Write([]byte("Conexión exitosa a MongoDB"))
}

func main() {

    Products.SetupRoutes()
    routes.SetupRoutesUsers()

	log.Println("Servidor escuchando en el puerto 8080...")

	log.Fatal(http.ListenAndServe(":8080", nil))
}

