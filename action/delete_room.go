package action

import (
	"chatsystem/db"
	"chatsystem/model"

	"github.com/globalsign/mgo/bson"
	"gitlab.ghn.vn/common-projects/go-sdk/sdk"
)

type DeleteChatRoom struct{}

func (a DeleteChatRoom) Delete(input *model.ReqDeleteRoom) *sdk.APIResponse {
	filter := bson.M{}
	filter["ticket_id"] = bson.ObjectIdHex(input.TicketID)
	resp := db.TicketDB().QueryOne(filter)
	if resp.Status != sdk.APIStatus.Ok {
		return resp
	}

	resp = db.ChatRoomDB().Delete(filter)
	return resp
}
