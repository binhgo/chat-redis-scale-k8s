package model

type Message struct {
	NameUser   string `json:"nameUser,omitempty" bson:"name_user,omitempty"`
	Mess       string `json:"mess,omitempty" bson:"mess,omitempty"`
	ChatRoomID string `json:"chatRoomID,omitempty" bson:"chat_room_id,omitempty"`
}
