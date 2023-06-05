package model

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Msg struct {
	ID        primitive.ObjectID `json:"id"   bson:"_id,omitempty"`
	Title     string             `json:"title"   bson:"title"`
	Content   string             `json:"content" bson:"content"`
	Links     string             `json:"links" bson:"links"`
	CreatedAt int64              `json:"createdAt" bson:"createdAt"`
}

func (m *Msg) Table() string {
	return "msg"
}

func (m *Msg) Create(ctx context.Context, db *mongo.Database) (*Msg, error) {
	now := time.Now().Unix()
	m.CreatedAt = now

	res, err := db.Collection(m.Table()).InsertOne(ctx, &m)
	if err != nil {
		return nil, err
	}
	m.ID = res.InsertedID.(primitive.ObjectID)
	return m, nil
}
