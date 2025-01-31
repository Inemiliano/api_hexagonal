package infraestructure

import (
    "api/src/Products/domain"
    "api/src/core"
    "context"
    "log"
    "time"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoRepository struct {
    collection *mongo.Collection
}

func NewMongos() *MongoRepository {
    client := core.GetMongoClient()
    collection := client.Database("api_hexa").Collection("Products")
    return &MongoRepository{collection: collection}
}

func (r *MongoRepository) Save(p *domain.Product) error {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    res, err := r.collection.InsertOne(ctx, p)
    if err != nil {
        log.Printf("Error al insertar producto: %v", err)
        return err
    }

    log.Printf("Producto insertado con éxito con ID: %v", res.InsertedID)
    return nil
}

func (r *MongoRepository) GetAll() ([]domain.Product, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    cursor, err := r.collection.Find(ctx, bson.M{})
    if err != nil {
        log.Printf("Error al obtener productos: %v", err)
        return nil, err
    }
    defer cursor.Close(ctx)

    var products []domain.Product
    if err = cursor.All(ctx, &products); err != nil {
        log.Printf("Error al leer los productos: %v", err)
        return nil, err
    }

    return products, nil
}

func (r *MongoRepository) Delete(id primitive.ObjectID) error {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    _, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
    if err != nil {
        log.Printf("Error al eliminar el producto: %v", err)
    }
    return err
}

func (repo *MongoRepository) Update(id primitive.ObjectID, p *domain.Product) error {
	// Crear contexto
	ctx := context.TODO()

	// Crear filtro para encontrar el documento a actualizar
	filter := bson.M{"_id": id}

	// Crear actualización
	update := bson.M{
		"$set": bson.M{
			"nombre": p.Nombre,
			"precio": p.Precio,
		},
	}

	// Actualizar el producto
	_, err := repo.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Println("Error al actualizar el producto:", err)
		return err
	}

	log.Println("Producto actualizado correctamente")
	return nil
}
