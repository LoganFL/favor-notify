package core

import (
	"favor-notify/internal/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserInterface interface {
	GetUserById(id primitive.ObjectID) (*model.User, error)

	GetUsersByDaoId(daoId primitive.ObjectID) ([]*model.User, error)
}
