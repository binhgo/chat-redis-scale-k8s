package model

type RateTicket struct {
	TicketID string
	Rate     int
}

type CloseTicket struct {
	TicketID string
	IsClose  bool
}

type ReqCreateTicket struct {
	Title       string
	Description string
	OrderCode   string
	UserID      string
}

type ReqGetTicket struct {
	OrderCode string
}

type ReqGetRoom struct {
	TicketID string
}

type ReqGetMessage struct {
	ChatRoomID string
}

type ReqCreateChatRoom struct {
	TicketID string
	UserID   []string
	Name     string
}

type ReqDeleteRoom struct {
	TicketID string
}
