package action

import (
	"chatsystem/db"
	"chatsystem/model"
	"chatsystem/util"
	"errors"
	"fmt"
	"sort"

	"github.com/globalsign/mgo/bson"
	"gitlab.ghn.vn/common-projects/go-sdk/sdk"
)

type GetTicket struct {
}

func (a GetTicket) GetOne(input *model.ReqGetTicket) *sdk.APIResponse {
	filter := bson.M{}
	filter["order_code"] = input.OrderCode
	filter["is_close"] = false
	// filter := model.Ticket{}
	// filter.OrderCode = new(string)
	// *filter.OrderCode = input.OrderCode

	resp := db.TicketDB().QueryOne(filter)
	if resp.Status != sdk.APIStatus.Ok || resp.Status == sdk.APIStatus.NotFound {
		return resp
	}

	mess := util.NewOkWithData("ok", resp.Data)

	return mess
}

func (a GetTicket) GetMany(off, lim int, isReverse bool) *sdk.APIResponse {
	//filter := bson.M{}

	tickets, err := a.queryTicket(off, lim, isReverse)
	if err != nil {
		mess := util.NewError(err.Error())
		return mess
	}
	fmt.Println(tickets)
	mess := util.NewOkWithData("ok", tickets)
	return mess
}

func (a GetTicket) queryTicket(off int, lim int, isReverse bool) ([]model.Ticket, error) {

	result := db.TicketDB().Query(bson.M{}, off, lim, false)

	if result.Status != sdk.APIStatus.Ok || result.Status == sdk.APIStatus.NotFound {
		return nil, errors.New(result.Message)
	}

	tickets, err := util.InterfaceToTicketArray(result.Data)
	if err != nil {
		return nil, err
	}

	if len(tickets) == 0 {
		return nil, errors.New("not found ticket")
	}

	if isReverse {
		sort.Slice(tickets, func(i, j int) bool {
			return tickets[i].ID.Hex() < tickets[j].ID.Hex()
		})
	}

	return tickets, nil
}
