package action

import (
	"chatsystem/db"
	"chatsystem/model"
	"chatsystem/util"

	"github.com/globalsign/mgo/bson"
	"gitlab.ghn.vn/common-projects/go-sdk/sdk"
)

type CreateChatRoom struct{}

func (a CreateChatRoom) Create(input *model.ReqCreateChatRoom) *sdk.APIResponse {
	room := model.ChatRoom{}
	room.ID = new(bson.ObjectId)
	*room.ID = bson.NewObjectId()
	room.Name = input.Name
	room.TicketID = new(bson.ObjectId)
	*room.TicketID = bson.ObjectIdHex(input.TicketID)
	room.UserID = input.UserID

	check := db.ChatRoomDB().QueryOne(room)
	if check.Status == sdk.APIStatus.Ok {
		return &sdk.APIResponse{Status: sdk.APIStatus.Error, Message: "ChatRoom already exist."}
	}

	resp := db.ChatRoomDB().Create(room)
	if resp.Status != sdk.APIStatus.Ok {
		return resp
	}
	mess := util.NewOkWithData("ok", resp.Data)
	return mess
}
