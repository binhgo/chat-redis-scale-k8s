package db

import (
	"errors"

	"gitlab.ghn.vn/common-projects/go-sdk/sdk"
)

// db session
var dbSession *sdk.DBSession
var dbLogSession *sdk.DBSession
var dbQueueSession *sdk.DBSession

// db name
var dbName string
var dbLogName string
var dbQueueName string

func SetDBSession(session *sdk.DBSession) {
	dbSession = session
}

func GetDBSession() (*sdk.DBSession, error) {
	if dbSession != nil {
		return dbSession, nil
	}

	return nil, errors.New("dbSession = nil")
}

func SetDBLogSession(session *sdk.DBSession) {
	dbLogSession = session
}

func GetDBLogSession() (*sdk.DBSession, error) {
	if dbLogSession != nil {
		return dbLogSession, nil
	}

	return nil, errors.New("dbSession = nil")
}

func SetDBQueueSession(session *sdk.DBSession) {
	dbQueueSession = session
}

func GetDBQueueSession() (*sdk.DBSession, error) {
	if dbQueueSession != nil {
		return dbQueueSession, nil
	}

	return nil, errors.New("dbSession = nil")
}

func SetDBName(name string) {
	dbName = name
}

func GetDBName() string {
	return dbName
}

func SetDBLogName(name string) {
	dbLogName = name
}

func GetDBLogName() string {
	return dbLogName
}

func SetDBQueueName(name string) {
	dbQueueName = name
}

func GetDBQueueName() string {
	return dbQueueName
}
