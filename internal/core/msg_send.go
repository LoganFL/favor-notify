package core

import "favor-notify/internal/model"

type MsgSendInterface interface {
	CreateMsgSend(ms *model.MsgSend) (*model.MsgSend, error)
}
