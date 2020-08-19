package api

import (
	"chatsystem/action"
	"chatsystem/model"
	"chatsystem/util"
	"fmt"

	"gitlab.ghn.vn/common-projects/go-sdk/sdk"
)

//GetRoom to get room infomration
func GetRoom(req sdk.APIRequest, res sdk.APIResponder) error {
	q := req.GetParam("q")

	var input *model.ReqGetRoom
	err := util.FromJson([]byte(q), &input)
	if err != nil {
		return res.Respond(util.NewError(err.Error()))
	}

	mess := action.GetRoom{}.GetOne(input)
	return res.Respond(mess)
}

//GetMessage to get chatroom's message
func GetMessage(req sdk.APIRequest, res sdk.APIResponder) error {
	q := req.GetParam("q")

	var input *model.ReqGetMessage
	err := util.FromJson([]byte(q), &input)
	if err != nil {
		return res.Respond(util.NewError(err.Error()))
	}

	mess := action.GetMessage{}.GetMany(input)
	return res.Respond(mess)
}

func CreateChatRoom(req sdk.APIRequest, res sdk.APIResponder) error {
	var input *model.ReqCreateChatRoom
	jData := req.GetContentText()

	err := util.FromJson([]byte(jData), &input)
	if err != nil {
		return res.Respond(util.NewError(err.Error()))
	}

	//validate
	if input == nil {
		str := fmt.Sprintf("require [%+v]", model.ReqCreateTicket{})
		mess := util.NewInvalid(str)
		return res.Respond(mess)
	}

	mess := action.CreateChatRoom{}.Create(input)
	return res.Respond(mess)
}

func DeleteChatRoom(req sdk.APIRequest, res sdk.APIResponder) error {
	var input *model.ReqDeleteRoom
	inputData := req.GetContentText()

	err := util.FromJson([]byte(inputData), &input)
	if err != nil {
		return res.Respond(util.NewError(err.Error()))
	}

	if input == nil || len(input.TicketID) <= 0 {
		mess := util.NewInvalid("")
		return res.Respond(mess)
	}

	mess := action.DeleteChatRoom{}.Delete(input)
	return res.Respond(mess)
}