package mongo

import (
	"favor-notify/internal/conf"
	"favor-notify/internal/core"
	"github.com/Masterminds/semver/v3"
)

var (
	_ core.CoreInterface = (*coreService)(nil)
	_ core.VersionInfo   = (*coreService)(nil)
)

type coreService struct {
	core.MsgInterface
	core.MsgSysInterface
	core.MsgSendInterface
	core.UserInterface
	core.OrganInterface
}

func NewCoreService() (core.CoreInterface, core.VersionInfo) {
	db := conf.MustMongoDB()

	cs := &coreService{
		newMsgManageService(db),
		newMsgSysMangeService(db),
		newMsgSendMangeService(db),
		newUserManageService(db),
		newOrganManageService(db),
	}
	return cs, cs
}

func (c coreService) Name() string {
	return "Mongo"
}

func (c *coreService) Version() *semver.Version {
	return semver.MustParse("v0.1.0")
}
