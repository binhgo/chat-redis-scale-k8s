package action

import (
	"chatsystem/db"
	"chatsystem/model"

	"github.com/globalsign/mgo/bson"
	"gitlab.ghn.vn/common-projects/go-sdk/sdk"
)

type UpdateTicket struct {
}

func (a UpdateTicket) Rate(input *model.RateTicket) *sdk.APIResponse {
	filter := bson.M{}
	filter["_id"] = bson.ObjectIdHex(input.TicketID)
	resp := db.TicketDB().QueryOne(filter)
	if resp.Status != sdk.APIStatus.Ok {
		return resp
	}

	updater := &model.Ticket{}
	updater.Rate = input.Rate

	resp = db.TicketDB().UpdateOne(filter, updater)
	return resp
}

func (a UpdateTicket) Close(input *model.CloseTicket) *sdk.APIResponse {
	filter := bson.M{}
	filter["_id"] = bson.ObjectIdHex(input.TicketID)
	resp := db.TicketDB().QueryOne(filter)
	if resp.Status != sdk.APIStatus.Ok {
		return resp
	}

	updater := &model.Ticket{}
	updater.IsClose = &input.IsClose

	resp = db.TicketDB().UpdateOne(filter, updater)
	return resp
}
