package model

import (
	"github.com/globalsign/mgo/bson"
)

type ChatRoom struct {
	ID       *bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Name     string         `json:"name,omitempty" bson:"name,omitempty"`
	TicketID *bson.ObjectId `json:"ticketId,omitempty" bson:"ticket_id,omitempty"`
	UserID   []string       `json:"userID,omitempty" bson:"user_id,omitempty"`
}
