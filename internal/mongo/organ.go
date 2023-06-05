package mongo

import (
	"context"
	"favor-notify/internal/core"
	"favor-notify/internal/model"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	_ core.OrganInterface = (*organManageService)(nil)
)

type organManageService struct {
	db *mongo.Database
}

func (o organManageService) GetOrganByKey(key string) (*model.Organ, error) {
	organ := &model.Organ{Key: key}
	return organ.Get(context.TODO(), o.db)
}

func newOrganManageService(db *mongo.Database) core.OrganInterface {
	return &organManageService{
		db: db,
	}
}
