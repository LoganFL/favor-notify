package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID         primitive.ObjectID `json:"id"               bson:"_id,omitempty"`
	CreatedOn  int64              `json:"created_on"       bson:"created_on"`
	ModifiedOn int64              `json:"modified_on"      bson:"modified_on"`
	DeletedOn  int64              `json:"deleted_on"       bson:"deleted_on"`
	IsDel      int                `json:"is_del"           bson:"is_del"`
	Nickname   string             `json:"nickname"         bson:"nickname"`
	Address    string             `json:"address"          bson:"address"`
	Avatar     string             `json:"avatar"           bson:"avatar"`
	Role       string             `json:"role"             bson:"role"`
	Token      string             `json:"token"            bson:"token"`
	LoginAt    int64              `json:"login_at"         bson:"login_at"`
}

func (m *User) Table() string {
	return "d_user"
}

func (m *User) Get(ctx context.Context, db *mongo.Database) (*User, error) {
	var (
		user User
		res  *mongo.SingleResult
	)
	if !m.ID.IsZero() {
		filter := bson.D{{"_id", m.ID}, {"is_del", 0}, {"token", bson.M{"$ne": bson.TypeNull}}}
		res = db.Collection(m.Table()).FindOne(ctx, filter)
	} else if m.Address != "" {
		filter := bson.D{{"address", m.Address}, {"is_del", 0}}
		res = db.Collection(m.Table()).FindOne(ctx, filter)
	}
	err := res.Err()
	if err != nil {
		return &user, err
	}
	err = res.Decode(&user)
	if err != nil {
		return &user, err
	}
	return &user, nil
}

func (m *User) List(db *mongo.Database, conditions *ConditionsT, daoId primitive.ObjectID) ([]*User, error) {
	var (
		users  []*User
		err    error
		cursor *mongo.Cursor
		query  bson.M
	)
	if len(*conditions) == 0 {
		if query != nil {
			query = findQuery([]bson.M{query})
		} else {
			query = bson.M{"is_del": 0}
		}
	}
	query = findQuery([]bson.M{query, {"token": bson.M{"$en": ""}}, {"dao.dao_id": daoId}})
	for _, v := range *conditions {
		if query != nil {
			query = findQuery([]bson.M{query, v})
		} else {
			query = findQuery([]bson.M{v})
		}
	}

	pipeline := mongo.Pipeline{
		{{"$lookup", bson.M{
			"from":         "dao_bookmark",
			"localField":   "address",
			"foreignField": "address",
			"as":           "dao",
		}}},
		{{"$match", query}},
		{{"$unwind", "$dao"}},
	}
	ctx := context.TODO()
	cursor, err = db.Collection(m.Table()).Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var user User
		if cursor.Decode(&user) != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}
