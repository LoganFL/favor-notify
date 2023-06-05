package core

import "favor-notify/internal/model"

type MsgSysInterface interface {
	CreateMsgSys(ms *model.MsgSys) (*model.MsgSys, error)
}
