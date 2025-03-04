package infraestructure
import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// getNextSequence obtiene el siguiente ID autoincrementable desde la colección "counters"
func (r *MongoRepository) getNextSequence(sequenceName string) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Filtro para buscar la secuencia
	filter := bson.M{"_id": sequenceName}

	// Incrementa en 1 el valor actual de "seq"
	update := bson.M{"$inc": bson.M{"seq": 1}}

	// Opciones para que cree el documento si no existe y retorne el nuevo valor
	options := options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After)

	// Estructura para capturar el nuevo valor del contador
	var result struct {
		Seq int `bson:"seq"`
	}

	// Realizar la actualización y obtener el nuevo ID
	err := r.collection.Database().Collection("counters").FindOneAndUpdate(ctx, filter, update, options).Decode(&result)
	if err != nil {
		log.Printf("Error al actualizar la secuencia: %v", err)
		return 0, err
	}

	return result.Seq, nil
}
