package redis_chat

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

var WaitTimeout time.Duration
var Log logrus.FieldLogger
var RR *RedisReceiver
var RW *RedisWriter

// Message sent to us by the javascript client
type Message struct {
	Handle string `json:"handle"`
	Text   string `json:"text"`
}

// ValidateMessage so that we know it's valid JSON and contains a Handle and
// Text
func ValidateMessage(data []byte) (Message, error) {
	var msg Message

	if err := json.Unmarshal(data, &msg); err != nil {
		return msg, errors.Wrap(err, "Unmarshaling Message")
	}

	if msg.Handle == "" && msg.Text == "" {
		return msg, errors.New("Message has no Handle or Text")
	}

	return msg, nil
}

// HandleWebsocket connection.
func HandleWebsocket(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		m := "Unable to upgrade to websockets"
		Log.WithField("err", err).Println(m)
		http.Error(w, m, http.StatusBadRequest)
		return
	}

	RR.Register(ws)

	for {
		mt, data, err := ws.ReadMessage()
		l := Log.WithFields(logrus.Fields{"mt": mt, "data": data, "err": err})
		if err != nil {
			if websocket.IsCloseError(err, websocket.CloseGoingAway) || err == io.EOF {
				l.Info("Websocket closed!")
				break
			}
			l.Error("Error reading websocket Message")
		}
		switch mt {
		case websocket.TextMessage:
			msg, err := ValidateMessage(data)
			if err != nil {
				l.WithFields(logrus.Fields{"msg": msg, "err": err}).Error("Invalid Message")
				break
			}
			RW.Publish(data)
		default:
			l.Warning("Unknown Message!")
		}
	}

	RR.DeRegister(ws)

	ws.WriteMessage(websocket.CloseMessage, []byte{})
}
