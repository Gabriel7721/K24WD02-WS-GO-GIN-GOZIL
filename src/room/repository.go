package room

import "go.mongodb.org/mongo-driver/v2/mongo"

type Repository struct {
	Rooms *mongo.Collection
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{
		Rooms: db.Collection("rooms"),
	}
}
