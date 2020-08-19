package util

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"chatsystem/model"

	jsoniter "github.com/json-iterator/go"
)

// a Faster json parser
var json = jsoniter.ConfigCompatibleWithStandardLibrary

// to struct
func FromJson(data []byte, v interface{}) error {
	// call UnMarshal
	err := json.Unmarshal(data, v)
	return err
}

// to json
func ToJson(object interface{}) ([]byte, error) {
	// call Marshal
	b, err := json.Marshal(&object)
	return b, err
}

// func InterfaceToTicket(data interface{}) (*model.Ticket, error) {

// 	tickets, ok := data.([]model.Ticket)
// 	if !ok {
// 		return nil, errors.New("CastToTicket Fail")
// 	}

// 	if len(tickets) == 0 {
// 		return nil, errors.New("CastToTicket Fail")
// 	}

// 	return &tickets[0], nil
// }

func InterfaceToTicketArray(data interface{}) ([]model.Ticket, error) {

	tickets, ok := data.([]model.Ticket)

	if !ok {
		return nil, errors.New("CastToTicket Fail")
	}

	if len(tickets) == 0 {
		return nil, errors.New("CastToTicket Fail")
	}

	return tickets, nil
}

func GetRandom(n int) string {

	rand.Seed(time.Now().UnixNano())

	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

func GenChatroomName(username, orderCode string) string {
	str := fmt.Sprintf("%s-%s", orderCode, username)
	return str
}

func ToInt(str string, defaultValue int) int {

	var result int
	intValue, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		result = defaultValue
	} else {
		result = int(intValue)
	}
	return result
}

func ToBool(str string, defaultValue bool) bool {

	var result bool
	boolValue, err := strconv.ParseBool(str)
	if err != nil {
		result = defaultValue
	} else {
		result = boolValue
	}
	return result
}
