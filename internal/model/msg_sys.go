package model

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MsgSys struct {
	ID        primitive.ObjectID `json:"id"   bson:"_id,omitempty"`
	From      primitive.ObjectID `json:"from" bson:"from"`
	Title     string             `json:"title"   bson:"title"`
	Content   string             `json:"content" bson:"content"`
	Links     string             `json:"links" bson:"links"`
	CreatedAt int64              `json:"createdAt" bson:"createdAt"`
}

func (m *MsgSys) Table() string {
	return "msg_sys"
}

func (m *MsgSys) Create(ctx context.Context, db *mongo.Database) (*MsgSys, error) {
	now := time.Now().Unix()
	m.CreatedAt = now
	res, err := db.Collection(m.Table()).InsertOne(ctx, &m)
	if err != nil {
		return nil, err
	}
	m.ID = res.InsertedID.(primitive.ObjectID)
	return m, nil
}
