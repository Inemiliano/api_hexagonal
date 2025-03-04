package infraestructure

import (
	"api/src/Products/domain"
	"api/src/core"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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

	// Obtener un ID autoincrementable
	newID, err := r.getNextSequence("productid")
	if err != nil {
		log.Printf(" Error al obtener el ID autoincrementable: %v", err)
		return err
	}

	// Asignar el nuevo ID al producto
	p.ID = newID

	// Insertar el producto en la colección de MongoDB
	res, err := r.collection.InsertOne(ctx, p)
	if err != nil {
		log.Printf("Error al insertar el producto: %v", err)
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
    for cursor.Next(ctx) {
        var product domain.Product

        
        if err := cursor.Decode(&product); err != nil {
            log.Printf("Error al decodificar un producto: %v", err)
            continue
        }

        products = append(products, product)
    }

    
    if err := cursor.Err(); err != nil {
        log.Printf("Error al recorrer los productos: %v", err)
        return nil, err
    }

    return products, nil
}


func (r *MongoRepository) Delete(name string) error {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    _, err := r.collection.DeleteOne(ctx, bson.M{"nombre": name})
    if err != nil {
        log.Printf("Error al eliminar el producto: %v", err)
    }
    return err
}

func (r *MongoRepository) Update(id string, p *domain.Product) error {
	// Crear contexto
	ctx := context.TODO()

	// Filtro para encontrar el producto por su ID
	filter := bson.M{"nombre": id} // Cambiar _id a id, porque ahora es un int

	// Datos de actualización
	update := bson.M{
		"$set": bson.M{
			"nombre": p.Nombre,
			"precio": p.Precio,
		},
	}

	// Ejecutar actualización en MongoDB
	_, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Println("Error al actualizar el producto:", err)
		return err
	}

	log.Println("Producto actualizado correctamente")
	return nil
}