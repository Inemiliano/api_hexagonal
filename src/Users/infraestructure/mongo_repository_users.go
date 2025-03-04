package infraestructure

import (
	"api/src/Users/domain"
	"api/src/core"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoUserRepository struct {
	collection *mongo.Collection
}

func NewMongoUserRepository() *MongoUserRepository {
	client := core.GetMongoClient()
	collection := client.Database("api_hexa").Collection("Users")
	return &MongoUserRepository{collection: collection}
}

func (r *MongoUserRepository) Save(user *domain.User) error {
	_, err := r.collection.InsertOne(context.TODO(), user)
	return err
}

func (r *MongoUserRepository) GetAll() ([]domain.User, error) {
	var users []domain.User
	cursor, err := r.collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var user domain.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *MongoUserRepository) Update(id primitive.ObjectID, user *domain.User) error {
    filter := bson.M{"_id": id}
    update := bson.M{"$set": bson.M{
        "name":  user.Name,
        "email": user.Email,
    }}

    _, err := r.collection.UpdateOne(context.TODO(), filter, update)
    return err
}



func (r *MongoUserRepository) Delete(id primitive.ObjectID) error {
    _, err := r.collection.DeleteOne(context.TODO(), bson.M{"_id": id})
    return err
}