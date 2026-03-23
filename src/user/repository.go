package user

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Repository struct {
	Collection *mongo.Collection
}

func NewRepository(db *mongo.Database) *Repository {
	collection := db.Collection("users")
	if err := CreateUserIndexes(collection); err != nil {
		panic(err)
	}
	return &Repository{
		Collection: collection,
	}
}

func (r *Repository) Create(user *User) {
	user.CreatedAt = time.Now()
	r.Collection.InsertOne(context.TODO(), user)
}

func (r *Repository) FindByEmail(email string) *User {
	var user User
	r.Collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	return &user
}

func CreateUserIndexes(collection *mongo.Collection) error {
	names, err := collection.Indexes().CreateMany(context.TODO(), []mongo.IndexModel{
		{
			Keys:    bson.D{{Key: "username", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys:    bson.D{{Key: "email", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	})
	if err != nil {
		return err
	}
	fmt.Println("Created Indexes:", names)
	return nil
}
