package db

import (
	"time"

	"chatsystem/model"

	"gitlab.ghn.vn/common-projects/go-sdk/sdk"
)

var _dbChatRoom *sdk.DBModel2 = nil

func ChatRoomDB() *sdk.DBModel2 {
	if _dbChatRoom == nil {
		_dbChatRoom = &sdk.DBModel2{DBName: GetDBName(), ColName: "chatroom"}
		_dbChatRoom.TemplateObject = model.ChatRoom{}
		session, err := GetDBSession()
		if err != nil {
			time.Sleep(2 * time.Second)
			// retry
			session, err = GetDBSession()
		}

		err = _dbChatRoom.Init(session)
		if err != nil {
			panic(err)
		}
		// todo: index
	}
	return _dbChatRoom
}
