package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"chatsystem/api"
	"chatsystem/conf"
	"chatsystem/db"
	"chatsystem/redis_chat"
	"github.com/heroku/x/hredis/redigo"

	"github.com/globalsign/mgo/bson"
	"github.com/sirupsen/logrus"

	"gitlab.ghn.vn/common-projects/go-sdk/sdk"
)

// const
const AppName = "CHAT_SERVICE"
const BaseUrl = "/cms"
const Version = "v1"
const Port = 8080
const DBName = "Demo"

// var
var startTime time.Time
var WaitTimeout = time.Minute * 10
var Log = logrus.WithField("cmd", "chat-system")
var RR redis_chat.RedisReceiver
var RW redis_chat.RedisWriter

// GetURL get service name + version
func GetURL() string {
	return "/cms/v1"
}

func main() {

	// redis
	redisURL := "redis://localhost/"
	redisPool, err2 := redigo.NewRedisPoolFromURL(redisURL)
	if err2 != nil {
		Log.WithField("url", redisURL).Fatal("Unable to create Redis pool")
	}
	RR = redis_chat.NewRedisReceiver(redisPool)
	RW = redis_chat.NewRedisWriter(redisPool)
	// setup variable in redis_chat
	redis_chat.RR = &RR
	redis_chat.RW = &RW
	redis_chat.Log = Log
	redis_chat.WaitTimeout = WaitTimeout

	go func() {
		for {
			waited, err := redigo.WaitForAvailability(redisURL, redis_chat.WaitTimeout, RR.Wait)
			if !waited || err != nil {
				Log.WithFields(logrus.Fields{"waitTimeout": redis_chat.WaitTimeout, "err": err}).Fatal("Redis not available by timeout!")
			}
			RR.Broadcast(redis_chat.AvailableMessage)
			err = RR.Run()
			if err == nil {
				break
			}
			Log.Error(err)
		}
	}()

	go func() {
		for {
			waited, err := redigo.WaitForAvailability(redisURL, redis_chat.WaitTimeout, nil)
			if !waited || err != nil {
				Log.WithFields(logrus.Fields{"waitTimeout": redis_chat.WaitTimeout, "err": err}).Fatal("Redis not available by timeout!")
			}
			err = RW.Run()
			if err == nil {
				break
			}
			Log.Error(err)
		}
	}()
	// channel to check if all db are connected (db, db log, db queue)
	// chanDBConnected = make(chan int)

	// get start server time
	startTime = time.Now()

	var app = sdk.NewApp("chat")

	conf.SetApp(app)

	setUpDB(conf.GetApp())

	// setup API Server
	protocol := conf.HTTP // conf.GetEnv().Protocol
	// if len(protocol) == 0 {
	// 	protocol = conf.HTTP
	// }

	var server, err = app.SetupAPIServer(protocol)
	if err != nil {
		panic(err)
	}

	// api

	// info
	server.SetHandler(sdk.APIMethod.GET, GetURL()+"/info", info)
	// ticket
	server.SetHandler(sdk.APIMethod.GET, GetURL()+"/ticket/get-tickets", api.GetAllTicket)
	server.SetHandler(sdk.APIMethod.GET, GetURL()+"/ticket", api.GetTicketByOrder)
	server.SetHandler(sdk.APIMethod.PUT, GetURL()+"/ticket/rate", api.RateTicket)
	server.SetHandler(sdk.APIMethod.PUT, GetURL()+"/ticket/close", api.CloseTicket)
	server.SetHandler(sdk.APIMethod.POST, GetURL()+"/ticket/create", api.CreateTicket)
	// chat room
	server.SetHandler(sdk.APIMethod.GET, GetURL()+"/chatroom/view", api.GetRoom)
	server.SetHandler(sdk.APIMethod.GET, GetURL()+"/chatroom/history", api.GetMessage)
	server.SetHandler(sdk.APIMethod.POST, GetURL()+"/chatroom/create", api.CreateChatRoom)
	server.SetHandler(sdk.APIMethod.DELETE, GetURL()+"/chatroom/delete", api.DeleteChatRoom)
	// websocket

	http.HandleFunc(GetURL()+"/chatroom/ws", redis_chat.HandleWebsocket)
	go func() {
		err1 := http.ListenAndServe(":8081", nil)
		if err1 != nil {
			panic(err1)
		}
	}()

	// launch
	server.Expose(80)
	err = app.Launch()
	if err != nil {
		panic(err)
	}

}

// callback main db
func onDBConnected(session *sdk.DBSession) error {
	db.SetDBSession(session)
	fmt.Println("onDBConnected")
	query := bson.M{}
	// query["test"] = "ok"

	var data []string

	myDB := session.GetMGOSession().DB("Demostgv2")
	myQuery := myDB.C("test").Find(query)

	err := myQuery.Distinct("parents.0", &data)
	log.Println(err)
	fmt.Println("data")
	// chanDBConnected <- 1
	fmt.Println("onDBConnected")
	return nil
}

// callback log db
func onDBLogConnected(session *sdk.DBSession) error {
	db.SetDBLogSession(session)
	fmt.Println("log")
	// chanDBConnected <- 1
	return nil
}

func setUpDB(app *sdk.App) {
	// var env string
	var configMap map[string]string

	// env = conf.GetEnv().Env
	configMap = conf.GetEnv().ConfigMap

	configMap = make(map[string]string)
	configMap["dbHost"] = "35.247.133.237"
	configMap["dbUser"] = "admin"
	configMap["dbPassword"] = "9aWuXyqKfnMvXc8rT27ZLRwww12345678"

	dbName := "Demostgv2" // DBName + env + Version
	dbNameLog := dbName + "_log"
	// dbNameQueue := dbName + "_queue"

	db.SetDBName(dbName)
	db.SetDBLogName(dbNameLog)
	// db.SetDBQueueName(dbNameQueue)

	// if env == conf.EnumEnv.Uat {
	// 	env = conf.EnumEnv.Prd
	// }

	// setup main database
	var db = app.SetupDBClient(sdk.DBConfiguration{
		Address:  strings.Split(configMap["dbHost"], ","),
		Username: configMap["dbUser"],
		Password: configMap["dbPassword"],
	})

	// setup log database
	var dbLog = app.SetupDBClient(sdk.DBConfiguration{
		Address:  strings.Split(configMap["dbHost"], ","),
		Username: configMap["dbUser"],
		Password: configMap["dbPassword"],
	})

	// setup queue database
	// var dbQueue = app.SetupDBClient(sdk.DBConfiguration{
	// 	Address:  strings.Split(configMap["dbHost"], ","),
	// 	Username: configMap["dbUser"],
	// 	Password: configMap["dbPassword"],
	// })

	// on db connected
	db.OnConnected(onDBConnected)
	dbLog.OnConnected(onDBLogConnected)
	// dbQueue.OnConnected(onDBQueueConnected)

}

func info(req sdk.APIRequest, res sdk.APIResponder) error {
	return res.Respond(&sdk.APIResponse{
		Status: sdk.APIStatus.Ok,
		Message: fmt.Sprintf("%s. Time start: %v:%v:%v \n",
			AppName,
			startTime.Hour(),
			startTime.Minute(),
			startTime.Second(),
		),
	})
}
