package action

import (
	"chatsystem/db"
	"chatsystem/model"
	"chatsystem/util"

	"github.com/globalsign/mgo/bson"
	"gitlab.ghn.vn/common-projects/go-sdk/sdk"
)

type GetRoom struct {}

func (a GetRoom) GetOne(input *model.ReqGetRoom) *sdk.APIResponse {
	filter := bson.M{}
	filter["ticket_id"] = input.TicketID

	resp := db.ChatRoomDB().QueryOne(filter)
	if resp.Status != sdk.APIStatus.Ok || resp.Status == sdk.APIStatus.NotFound {
		return resp
	}

	mess := util.NewOkWithData("ok", resp.Data)
	return mess
}