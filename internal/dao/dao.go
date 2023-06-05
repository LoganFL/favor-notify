package dao

import (
	"favor-notify/internal/core"
	"favor-notify/internal/mongo"
	"github.com/sirupsen/logrus"
	"sync"
)

var (
	cs     core.CoreInterface
	onceDs sync.Once
)

func DaoService() core.CoreInterface {
	onceDs.Do(func() {
		var v core.VersionInfo
		cs, v = mongo.NewCoreService()
		logrus.Infof("use %s as data service with version %s", v.Name(), v.Version())
	})
	return cs
}
