package model

import (
	"github.com/globalsign/mgo/bson"
	"time"
)

type User struct {
	ID          bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Username    string        `json:"username,omitempty" bson:"username,omitempty"`
	Password    string        `json:"password,omitempty" bson:"password,omitempty"`
	Email       string        `json:"email,omitempty" bson:"email,omitempty"`
	FullName    string       ` json:"fullname,omitempty" bson:"fullname,omitempty"`
	PhoneNumber string        `json:"phoneNumber,omitempty" bson:"phone_number,omitempty"`
	//Address     Address       `json:"address,omitempty" bson:"address,omitempty"`
	Avatar      string        `json:"avatar,omitempty" bson:"avatar,omitempty"`
	Role        string        `json:"role,omitempty" bson:"role,omitempty"`
	Status      string        `json:"status,omitempty" bson:"status,omitempty"`
	CreatedTime time.Time     `json:"createdTime,omitempty" bson:"created_time,omitempty"`
	Available   bool          `json:"available,omitempty" bson:"available,omitempty"`
}