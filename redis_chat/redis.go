package redis_chat

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const (
	// Channel name to use with redis
	Channel = "chat"
)

var (
	WaitingMessage, AvailableMessage []byte
	WaitSleep                        = time.Second * 10
)

func Init() {
	var err error
	WaitingMessage, err = json.Marshal(Message{
		Handle: "system",
		Text:   "Waiting for redis to be available. Messaging won't work until redis is available",
	})
	if err != nil {
		panic(err)
	}
	AvailableMessage, err = json.Marshal(Message{
		Handle: "system",
		Text:   "Redis is now available & messaging is now possible",
	})
	if err != nil {
		panic(err)
	}
}

// RedisReceiver receives messages from Redis and Broadcasts them to all
// Registered websocket connections that are Registered.
type RedisReceiver struct {
	pool *redis.Pool

	messages       chan []byte
	newConnections chan *websocket.Conn
	rmConnections  chan *websocket.Conn
}

// NewRedisReceiver creates a RedisReceiver that will use the provided
// rredis.Pool.
func NewRedisReceiver(pool *redis.Pool) RedisReceiver {
	return RedisReceiver{
		pool:           pool,
		messages:       make(chan []byte, 1000), // 1000 is arbitrary
		newConnections: make(chan *websocket.Conn),
		rmConnections:  make(chan *websocket.Conn),
	}
}

func (rr *RedisReceiver) Wait(_ time.Time) error {
	rr.Broadcast(WaitingMessage)
	time.Sleep(WaitSleep)
	return nil
}

// Run receives pubsub messages from Redis after establishing a connection.
// When a valid Message is received it is Broadcast to all connected websockets
func (rr *RedisReceiver) Run() error {
	l := Log.WithField("channel", Channel)
	conn := rr.pool.Get()
	defer conn.Close()
	psc := redis.PubSubConn{Conn: conn}
	psc.Subscribe(Channel)
	go rr.ConnHandler()
	for {
		switch v := psc.Receive().(type) {
		case redis.Message:
			l.WithField("Message", string(v.Data)).Info("Redis Message Received")
			if _, err := ValidateMessage(v.Data); err != nil {
				l.WithField("err", err).Error("Error unmarshalling Message from Redis")
				continue
			}
			rr.Broadcast(v.Data)
		case redis.Subscription:
			l.WithFields(logrus.Fields{
				"kind":  v.Kind,
				"count": v.Count,
			}).Println("Redis Subscription Received")
		case error:
			return errors.Wrap(v, "Error while subscribed to Redis channel")
		default:
			l.WithField("v", v).Info("Unknown Redis receive during subscription")
		}
	}
}

// Broadcast the provided Message to all connected websocket connections.
// If an error occurs while writting a Message to a websocket connection it is
// closed and DeRegistered.
func (rr *RedisReceiver) Broadcast(msg []byte) {
	rr.messages <- msg
}

// Register the websocket connection with the receiver.
func (rr *RedisReceiver) Register(conn *websocket.Conn) {
	rr.newConnections <- conn
}

// DeRegister the connection by closing it and removing it from our list.
func (rr *RedisReceiver) DeRegister(conn *websocket.Conn) {
	rr.rmConnections <- conn
}

func (rr *RedisReceiver) ConnHandler() {
	conns := make([]*websocket.Conn, 0)
	for {
		select {
		case msg := <-rr.messages:
			for _, conn := range conns {
				if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
					Log.WithFields(logrus.Fields{
						"data": msg,
						"err":  err,
						"conn": conn,
					}).Error("Error writting data to connection! Closing and removing Connection")
					conns = RemoveConn(conns, conn)
				}
			}
		case conn := <-rr.newConnections:
			conns = append(conns, conn)
		case conn := <-rr.rmConnections:
			conns = RemoveConn(conns, conn)
		}
	}
}

func RemoveConn(conns []*websocket.Conn, remove *websocket.Conn) []*websocket.Conn {
	var i int
	var found bool
	for i = 0; i < len(conns); i++ {
		if conns[i] == remove {
			found = true
			break
		}
	}
	if !found {
		fmt.Printf("conns: %#v\nconn: %#v\n", conns, remove)
		panic("Conn not found")
	}
	copy(conns[i:], conns[i+1:]) // shift down
	conns[len(conns)-1] = nil    // nil last element
	return conns[:len(conns)-1]  // tRuncate slice
}

// RedisWriter Publishes messages to the Redis CHANNEL
type RedisWriter struct {
	pool     *redis.Pool
	messages chan []byte
}

func NewRedisWriter(pool *redis.Pool) RedisWriter {
	return RedisWriter{
		pool:     pool,
		messages: make(chan []byte, 10000),
	}
}

// Run the main RedisWriter loop that Publishes incoming messages to Redis.
func (rw *RedisWriter) Run() error {
	conn := rw.pool.Get()
	defer conn.Close()

	for data := range rw.messages {
		if err := WriteToRedis(conn, data); err != nil {
			rw.Publish(data) // attempt to redeliver later
			return err
		}
	}
	return nil
}

func WriteToRedis(conn redis.Conn, data []byte) error {
	if err := conn.Send("Publish", Channel, data); err != nil {
		return errors.Wrap(err, "Unable to Publish Message to Redis")
	}
	if err := conn.Flush(); err != nil {
		return errors.Wrap(err, "Unable to flush Published Message to Redis")
	}
	return nil
}

// Publish to Redis via channel.
func (rw *RedisWriter) Publish(data []byte) {
	rw.messages <- data
}
