package model

import (
	"go.mongodb.org/mongo-driver/bson"
)

type ConditionsT map[string]bson.M

func findQuery(query []bson.M) bson.M {
	query = append(query, bson.M{"is_del": 0})
	if query != nil {
		if len(query) > 0 {
			return bson.M{"$and": query}
		}
	}
	return bson.M{}
}
