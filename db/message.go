package db

import (
	"time"

	"chatsystem/model"

	"gitlab.ghn.vn/common-projects/go-sdk/sdk"
)

var _dbMessage *sdk.DBModel2 = nil

func MessageDB() *sdk.DBModel2 {
	if _dbMessage == nil {
		_dbMessage = &sdk.DBModel2{DBName: GetDBName(), ColName: "message"}
		_dbMessage.TemplateObject = model.Message{}
		session, err := GetDBSession()
		if err != nil {
			time.Sleep(2 * time.Second)
			// retry
			session, err = GetDBSession()
		}

		err = _dbMessage.Init(session)
		if err != nil {
			panic(err)
		}
		// todo: index
	}
	return _dbMessage
}
