package services

import (
	"context"
	"favor-notify/internal/model"
	"favor-notify/pkg/errcode"
	"fmt"
	"github.com/gogf/gf/util/gconv"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type ReqPushNotify struct {
	IsSave    bool               `json:"isSave"`
	From      string             `json:"from"`
	FromType  model.FromTypeEnum `json:"fromType"`
	To        string             `json:"to"`
	Title     string             `json:"title"`
	Content   string             `json:"content"`
	Links     string             `json:"links"`
	Region    string             `json:"region"`
	NetworkId int                `json:"networkId"`
}

type ReqPushNotifySys struct {
	IsSave    bool   `json:"isSave"`
	From      string `json:"from"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Links     string `json:"links"`
	Region    string `json:"region"`
	NetworkId int    `json:"networkId"`
}

func PushNotify(req ReqPushNotify) *errcode.Error {
	userId, err := primitive.ObjectIDFromHex(req.To)
	if err != nil {
		return errcode.IDFailed
	}
	var from primitive.ObjectID
	if req.FromType == model.ORANGE {
		organ, err := ci.GetOrganByKey(req.From)
		if err != nil {
			return errcode.OrganFailed
		}
		from = organ.ID
	} else {
		from, err = primitive.ObjectIDFromHex(req.From)
		if err != nil {
			return errcode.IDFailed
		}
	}
	user, err := ci.GetUserById(userId)
	if err != nil {
		return errcode.UserFailed
	}
	if user == nil {
		return nil
	}
	data := make(map[string]string, 6)
	data["origin"] = req.From
	data["to"] = req.To
	data["links"] = req.Links
	data["region"] = req.Region
	data["networkId"] = gconv.String(req.NetworkId)

	_, err = notifyFirebase.Send(context.TODO(), user.Token, req.Title, req.Content, data)
	if err != nil {
		fmt.Println(err)
		logrus.Errorf("firebase errs: %v", err)
		return errcode.FirebaseSendFailed
	}
	if req.IsSave {
		msgId, errs := saveMsg(req)
		if errs != nil {
			return errs
		}
		return saveMsgSend(msgId, from, userId, req.FromType)
	}
	return nil
}

func PushNotifyDao(req ReqPushNotify) *errcode.Error {
	var msgId, from primitive.ObjectID

	daoId, err := primitive.ObjectIDFromHex(req.To)
	if err != nil {
		return errcode.IDFailed
	}
	if req.FromType == model.ORANGE {
		organ, err := ci.GetOrganByKey(req.From)
		if err != nil {
			return errcode.OrganFailed
		}
		from = organ.ID
	} else {
		from, err = primitive.ObjectIDFromHex(req.From)
		if err != nil {
			return errcode.IDFailed
		}
	}
	users, err := ci.GetUsersByDaoId(daoId)
	if err != nil {
		return errcode.DaoFailed
	}
	if len(users) == 0 {
		return nil
	}
	if req.IsSave {
		id, errs := saveMsg(req)
		if errs != nil {
			return errs
		}
		msgId = id
	}
	data := make(map[string]string, 6)
	data["origin"] = req.From
	data["to"] = req.To
	data["links"] = req.Links
	data["region"] = req.Region
	data["networkId"] = gconv.String(req.NetworkId)
	for _, user := range users {
		_, err = notifyFirebase.Send(context.TODO(), user.Token, req.Title, req.Content, data)
		if err != nil {
			logrus.Errorf("firebase errs: %v", err)
			return errcode.FirebaseSendFailed
		}
		if req.IsSave {
			errs := saveMsgSend(msgId, from, user.ID, req.FromType)
			if errs != nil {
				return errs
			}
		}
	}
	return nil
}

func PushNotifySys(req ReqPushNotifySys) *errcode.Error {

	organ, err := ci.GetOrganByKey(req.From)
	if err != nil {
		return errcode.OrganFailed
	}
	from := organ.ID
	data := make(map[string]string, 6)
	data["origin"] = req.From
	data["links"] = req.Links
	data["region"] = req.Region
	data["networkId"] = gconv.String(req.NetworkId)
	_, err = notifyFirebase.SendTopic(context.TODO(), req.From, req.Title, req.Content, data)
	if err != nil {
		logrus.Errorf("firebase errs: %v", err)
		return errcode.FirebaseSendFailed
	}
	if req.IsSave {
		sys := &model.MsgSys{
			From:      from,
			Title:     req.Title,
			Content:   req.Content,
			Links:     req.Links,
			CreatedAt: time.Now().Unix(),
		}
		_, err = ci.CreateMsgSys(sys)
		if err != nil {
			return errcode.MsgSaveFailed
		}
	}
	return nil
}

func saveMsg(req ReqPushNotify) (primitive.ObjectID, *errcode.Error) {
	t := time.Now().Unix()
	msg := &model.Msg{
		Title:     req.Title,
		Content:   req.Content,
		Links:     req.Links,
		CreatedAt: t,
	}
	sm, err := ci.CreateMsg(msg)
	if err != nil {
		return primitive.NewObjectID(), errcode.MsgSaveFailed
	}
	return sm.ID, nil
}

func saveMsgSend(msgId, from, to primitive.ObjectID, fromType model.FromTypeEnum) *errcode.Error {
	msgSend := &model.MsgSend{
		MsgID:     msgId,
		From:      from,
		To:        to,
		FromType:  fromType,
		CreatedAt: time.Now().Unix(),
	}
	_, err := ci.CreateMsgSend(msgSend)
	if err != nil {
		return errcode.MsgSendSaveFailed
	}
	return nil
}
