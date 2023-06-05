package mongo

import (
	"context"
	"favor-notify/internal/core"
	"favor-notify/internal/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	_ core.UserInterface = (*userManageService)(nil)
)

type userManageService struct {
	db *mongo.Database
}

func newUserManageService(db *mongo.Database) core.UserInterface {
	return &userManageService{
		db: db,
	}
}

func (u userManageService) GetUserById(id primitive.ObjectID) (*model.User, error) {
	user := &model.User{
		ID: id,
	}
	return user.Get(context.TODO(), u.db)
}

func (u userManageService) GetUsersByDaoId(daoId primitive.ObjectID) ([]*model.User, error) {
	user := &model.User{}
	return user.List(u.db, daoId)
}
