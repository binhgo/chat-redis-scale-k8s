package client

import (
	"chatsystem/conf"
	"chatsystem/util"
	"fmt"
	"gitlab.ghn.vn/common-projects/go-sdk/sdk"
)

type BmsClient struct {
}

var token string

type ReqLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type RespLogin struct {
	Status string `json:"status"`
	Data 	[]struct {
		Username	string `json:"username"`
		Token 		string `json:"token"`
	}
}

type ReqGetUser struct {
	Status string `json:"status"`
	Data 	[]struct {
		Username    string        `json:"username"`
		Avatar      string        `json:"avatar"`
		Role        string        `json:"role"`
	}
}

func (c BmsClient) Login() (*RespLogin,error) {
	config := conf.GetBMSConf()
	client := newClient(config.BaseConf)
	input := ReqLogin{Username:"bao.onggia",Password:"123123"}
	header := make(map[string]string)
	header["Content-Type"] = "application/json; charset=utf-8"
	//header["Authorization"] = config.BasicAuth

	result, err := client.MakeHTTPRequest(sdk.HTTPMethods.Post, header, nil, input, "user/login")
	if err != nil {
		return nil, err
	}
	var rs *RespLogin
	err=util.FromJson(result.Content,&rs)
	if err != nil {
		return nil, err
	}
	fmt.Println(rs)
	// return value
	return rs,nil
}

func (c BmsClient) GetUser(token string) (*ReqGetUser, error) {
	//fmt.Println(token)
	config := conf.GetBMSConf()
	client := newClient(config.BaseConf)
	header := make(map[string]string)
	header["Content-Type"] = "application/json; charset=utf-8"
	header["Authorization"] = "Bearer " + token

	result, err := client.MakeHTTPRequest(sdk.HTTPMethods.Get, header, nil, nil, "user/self")
	if err != nil {
		return nil, err
	}
	var rs *ReqGetUser
	err=util.FromJson(result.Content,&rs)
	if err != nil {
		return nil, err
	}
	fmt.Println(rs)

	return rs,nil
}

