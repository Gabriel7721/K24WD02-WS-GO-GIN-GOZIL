package friend

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Repository struct {
	FriendRequest *mongo.Collection
	Friends       *mongo.Collection
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{
		FriendRequest: db.Collection("friend_requests"),
		Friends:       db.Collection("friends"),
	}
}

func (r *Repository) SendRequest(fromID, toID bson.ObjectID) error {
	req := FriendRequest{
		FromUserID: fromID,
		ToUserID:   toID,
		Status:     "pending",
		CreatedAt:  time.Now(),
	}
	_, err := r.FriendRequest.InsertOne(context.TODO(), req)
	return err
}
func (r *Repository) AcceptRequest(requestID bson.ObjectID) error {
	_, err := r.FriendRequest.UpdateByID(context.TODO(), requestID, bson.M{"$set": bson.M{
		"status": "accept"}})
	if err != nil {
		return err
	}
	var req FriendRequest
	err = r.FriendRequest.FindOne(
		context.TODO(), bson.M{"_id": requestID}).Decode(&req)
	if err != nil {
		return err
	}
	friend := Friend{
		User1:     req.FromUserID,
		User2:     req.ToUserID,
		CreatedAt: time.Now(),
	}
	_, err = r.Friends.InsertOne(context.TODO(), friend)
	return err
}
