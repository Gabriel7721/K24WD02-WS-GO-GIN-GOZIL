package user

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Repository struct {
	Collection *mongo.Collection
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{
		Collection: db.Collection("users"),
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
