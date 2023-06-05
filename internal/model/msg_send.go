package model

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type FromTypeEnum uint8

const (
	DAO FromTypeEnum = iota
	ORANGE
	USER
)

type MsgSend struct {
	ID        primitive.ObjectID `json:"id"   bson:"_id,omitempty"`
	MsgID     primitive.ObjectID `json:"msgID" bson:"msg_id"`
	From      primitive.ObjectID `json:"from" bson:"from"`
	To        primitive.ObjectID `json:"to" bson:"to"`
	FromType  FromTypeEnum       `json:"fromType" bson:"fromType"`
	CreatedAt int64              `json:"createdAt" bson:"createdAt"`
}

func (m *MsgSend) Table() string {
	return "msg_send"
}

func (m *MsgSend) Create(ctx context.Context, db *mongo.Database) (*MsgSend, error) {
	now := time.Now().Unix()
	m.CreatedAt = now
	res, err := db.Collection(m.Table()).InsertOne(ctx, &m)
	if err != nil {
		return nil, err
	}
	m.ID = res.InsertedID.(primitive.ObjectID)
	return m, nil
}
