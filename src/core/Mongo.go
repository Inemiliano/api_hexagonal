package core

import (
	"context"
	"log"
	"os" // Necesario para usar os.Getenv()
	"sync"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongoClient *mongo.Client
	mongoOnce   sync.Once
)

// Carga de variables de entorno y conexión a MongoDB
func init() {
	// Cargar variables de entorno desde el archivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error al cargar el archivo .env")
	} else {
		log.Println("Archivo .env cargado exitosamente")
	}
}

// Obtener la instancia de cliente MongoDB con inicialización perezosa (solo una vez)
func GetMongoClient() *mongo.Client {
	mongoOnce.Do(func() {
		// Obtener URI de MongoDB desde el archivo .env
		mongoURI := os.Getenv("MONGO_URI")
		if mongoURI == "" {
			log.Fatal("MONGO_URI no está definido en el archivo .env")
		}

		log.Println("Conectando a MongoDB con URI:", mongoURI)

		// Configurar las opciones de conexión con MongoDB
		clientOptions := options.Client().ApplyURI(mongoURI)

		// Crear un contexto con timeout de 10 segundos para la conexión
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// Intentar conectarse al cliente MongoDB
		client, err := mongo.Connect(ctx, clientOptions)
		if err != nil {
			log.Fatal("Error al conectar con MongoDB:", err)
		}

		// Verificar que se pueda hacer un ping a MongoDB para confirmar la conexión
		err = client.Ping(ctx, nil)
		if err != nil {
			log.Fatal("No se pudo hacer ping a MongoDB:", err)
		}

		// Imprimir log si la conexión fue exitosa
		log.Println("Conexión a MongoDB establecida correctamente")
		mongoClient = client
	})

	// Retornar el cliente MongoDB
	return mongoClient
}
