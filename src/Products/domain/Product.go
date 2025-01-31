package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
    ID     primitive.ObjectID 
    Nombre string             
    Precio int16              
}
