package services

import (
	"favor-notify/internal/conf"
	"favor-notify/internal/core"
	"favor-notify/internal/dao"
	"favor-notify/pkg/firebase"
)

var (
	ci             core.CoreInterface
	notifyFirebase *firebase.Client
)

func Initialize() {
	ci = dao.DaoService()
	client, err := firebase.New(conf.FirebaseSetting.Config)
	if err != nil {
		panic(err)
	}
	notifyFirebase = client
}
