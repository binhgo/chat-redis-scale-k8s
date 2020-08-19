package db

import (
  "time"

  "gitlab.ghn.vn/common-projects/go-sdk/sdk"
  "chatsystem/model"
)

var _dbTicket *sdk.DBModel2 = nil

func TicketDB() *sdk.DBModel2 {
  if _dbTicket == nil {
    _dbTicket = &sdk.DBModel2{DBName: GetDBName(), ColName: "ticket"}
    _dbTicket.TemplateObject = model.Ticket{}
    session, err := GetDBSession()
    if err != nil {
      time.Sleep(2 * time.Second)
      // retry
      session, err = GetDBSession()
    }

    err = _dbTicket.Init(session)
    if err != nil {
      panic(err)
    }
    // todo: index
  }
  return _dbTicket
}