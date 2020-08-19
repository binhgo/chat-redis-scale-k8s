package api

import (
	"chatsystem/action"
	"chatsystem/client"
	"chatsystem/model"
	"chatsystem/util"
	"fmt"

	"gitlab.ghn.vn/common-projects/go-sdk/sdk"
)

func GetAllTicket(req sdk.APIRequest, res sdk.APIResponder) error {
	//check access
	loginRes, errLogin := client.BmsClient{}.Login()
	if errLogin != nil {
		return errLogin
	}
	token := loginRes.Data[0]
	roleRes, errRole :=client.BmsClient{}.GetUser(token.Token)
	if errRole != nil {
		return errRole
	}
	fmt.Println(roleRes.Data[0].Role)
	//params
	//q := req.GetParam("q")
	off := util.ToInt(req.GetParam("offset"), 0)
	lim := util.ToInt(req.GetParam("limit"), 30)
	isReverse := util.ToBool(req.GetParam("reverse"), false)

	mess := action.GetTicket{}.GetMany(off, lim, isReverse)
	return res.Respond(mess)
}

func GetTicketByOrder(req sdk.APIRequest, res sdk.APIResponder) error {

	//param
	q := req.GetParam("q")

	var input *model.ReqGetTicket

	err := util.FromJson([]byte(q), &input)
	if err != nil {
		return res.Respond(util.NewError(err.Error()))
	}

	mess := action.GetTicket{}.GetOne(input)
	return res.Respond(mess)
}

func RateTicket(req sdk.APIRequest, res sdk.APIResponder) error {
	var input *model.RateTicket
	inputData := req.GetContentText()

	err := util.FromJson([]byte(inputData), &input)
	if err != nil {
		return res.Respond(util.NewError(err.Error()))
	}

	if input == nil || len(input.TicketID) <= 0 || input.Rate < 0 {
		mess := util.NewInvalid("")
		return res.Respond(mess)
	}

	mess := action.UpdateTicket{}.Rate(input)
	return res.Respond(mess)
}

func CloseTicket(req sdk.APIRequest, res sdk.APIResponder) error {
	var input *model.CloseTicket
	inputData := req.GetContentText()

	err := util.FromJson([]byte(inputData), &input)
	if err != nil {
		return res.Respond(util.NewError(err.Error()))
	}

	if input == nil || len(input.TicketID) <= 0 {
		mess := util.NewInvalid("")
		return res.Respond(mess)
	}

	mess := action.UpdateTicket{}.Close(input)
	return res.Respond(mess)

}

func CreateTicket(req sdk.APIRequest, res sdk.APIResponder) error {
	var input *model.ReqCreateTicket
	jData := req.GetContentText()
	
	err := util.FromJson([]byte(jData), &input)
	if err != nil {
		return res.Respond(util.NewError(err.Error()))
	}
	fmt.Println(input)
	//validate
	if input == nil {
		str := fmt.Sprintf("require [%+v]", model.ReqCreateTicket{})
		mess := util.NewInvalid(str)
		return res.Respond(mess)
	}

	if input.Title == "" || input.OrderCode == "" {
		mess := util.NewInvalid("require title|order code")
		return res.Respond(mess)
	}

	mess := action.CreateTicket{}.Create(input)
	return res.Respond(mess)
}
