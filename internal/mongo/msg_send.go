package mongo

import (
	"context"

	"favor-notify/internal/core"
	"favor-notify/internal/model"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	_ core.MsgSendInterface = (*msgSendMangeService)(nil)
)

type msgSendMangeService struct {
	db *mongo.Database
}

func newMsgSendMangeService(db *mongo.Database) core.MsgSendInterface {
	return &msgSendMangeService{
		db: db,
	}
}

func (m msgSendMangeService) CreateMsgSend(ms *model.MsgSend) (*model.MsgSend, error) {
	return ms.Create(context.TODO(), m.db)
}
