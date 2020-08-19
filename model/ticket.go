package model

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

type Ticket struct {
	ID          *bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	CreatedTime *time.Time     `json:"createdTime,omitempty" bson:"created_time,omitempty"`
	Title       string         `json:"title,omitempty" bson:"title,omitempty"`
	Description string         `json:"description,omitempty" bson:"description,omitempty"`
	OrderCode   string         `json:"orderCode,omitempty" bson:"order_code,omitempty"`
	Status      string         `json:"status,omitempty" bson:"status,omitempty"` //DONE, PROCESSING, NEED_SOLVE
	Rate        int           `json:"rate,omitempty" bson:"rate,omitempty"`
	IsClose     *bool          `json:"isClose" bson:"is_close"`
	UserID      string         `json:"userId,omitempty" bson:"user_id,omitempty"`
}
