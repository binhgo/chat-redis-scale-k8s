package action

import (
	"chatsystem/db"
	"chatsystem/model"
	"chatsystem/util"

	"github.com/globalsign/mgo/bson"
	"gitlab.ghn.vn/common-projects/go-sdk/sdk"
)

type GetMessage struct{}

func (a GetMessage) GetMany(input *model.ReqGetMessage) *sdk.APIResponse {
	filter := bson.M{}
	filter["chat_room_id"] = bson.ObjectIdHex(input.ChatRoomID)

	resp := db.MessageDB().QueryOne(filter)
	if resp.Status != sdk.APIStatus.Ok || resp.Status == sdk.APIStatus.NotFound {
		return resp
	}

	mess := util.NewOkWithData("ok", resp.Data)
	return mess
}
