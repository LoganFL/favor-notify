package core

import "favor-notify/internal/model"

type MsgInterface interface {
	CreateMsg(msg *model.Msg) (*model.Msg, error)
}
