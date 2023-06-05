package core

import "favor-notify/internal/model"

type OrganInterface interface {
	GetOrganByKey(key string) (*model.Organ, error)
}
