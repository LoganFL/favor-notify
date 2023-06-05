package mongo

import (
	"context"

	"favor-notify/internal/core"
	"favor-notify/internal/model"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	_ core.MsgSysInterface = (*msgSysMangeService)(nil)
)

type msgSysMangeService struct {
	db *mongo.Database
}

func newMsgSysMangeService(db *mongo.Database) core.MsgSysInterface {
	return &msgSysMangeService{
		db: db,
	}
}

func (m msgSysMangeService) CreateMsgSys(ms *model.MsgSys) (*model.MsgSys, error) {
	return ms.Create(context.TODO(), m.db)
}
