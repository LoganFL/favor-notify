package mongo

import (
	"context"
	"favor-notify/internal/core"
	"favor-notify/internal/model"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	_ core.MsgInterface = (*msgManageService)(nil)
)

type msgManageService struct {
	db *mongo.Database
}

func newMsgManageService(db *mongo.Database) core.MsgInterface {
	return &msgManageService{
		db: db,
	}
}

func (m msgManageService) CreateMsg(msg *model.Msg) (*model.Msg, error) {
	return msg.Create(context.TODO(), m.db)
}
