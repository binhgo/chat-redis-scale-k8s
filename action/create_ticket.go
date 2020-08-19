package action

import (
	"chatsystem/db"
	"chatsystem/model"
	"chatsystem/util"

	"github.com/globalsign/mgo/bson"
	"gitlab.ghn.vn/common-projects/go-sdk/sdk"
)

type CreateTicket struct{}

func (a CreateTicket) Create(input *model.ReqCreateTicket) *sdk.APIResponse {
	filter := bson.M{}
	filter["order_code"] = input.OrderCode

	check := db.TicketDB().QueryOne(filter)
	if check.Status == sdk.APIStatus.Ok {
		return &sdk.APIResponse{Status: sdk.APIStatus.Error, Message: "One order code can't has more than one ticket"}
	}

	ticket := &model.Ticket{}
	ticket.ID = new(bson.ObjectId)
	*ticket.ID = bson.NewObjectId()
	ticket.Title = input.Title
	ticket.Description = input.Description
	ticket.OrderCode = input.OrderCode
	ticket.UserID = input.UserID
	ticket.IsClose = new(bool)
	*ticket.IsClose = false
	ticket.Status = "NEED_SOLVE"
	ticket.Rate = -1

	resp := db.TicketDB().Create(ticket)
	if resp.Status != sdk.APIStatus.Ok {
		return resp
	}

	mess := util.NewOkWithData("ok", resp.Data)
	return mess
}


