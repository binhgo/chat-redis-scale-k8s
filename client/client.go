package client

import (
	"chatsystem/conf"
	"chatsystem/db"
	"chatsystem/util"
	"gitlab.ghn.vn/common-projects/go-sdk/sdk"
	"time"
)

func setRestClientLog(c *sdk.RestClient) error {
	dbLogSS, err := db.GetDBLogSession()
	if err != nil {
		return err
	}
	c.SetDBLog(db.GetDBLogName(), dbLogSS)
	return nil
}

func newClient(config conf.BaseConf) *sdk.RestClient {
	client := sdk.NewRESTClient(config.URL, config.LogName, time.Duration(config.Timeout)*time.Second, config.MaxRetry, time.Duration(config.WaitTimeBetweenRetry)*time.Second)
	// ghi log
	setRestClientLog(client)
	return client
}

type Q struct {
	myMap map[string]string
}

func (q *Q) Append(key, value string) *Q {
	if q.myMap == nil {
		q.myMap = make(map[string]string)
	}

	q.myMap[key] = value
	return q
}

func (q *Q) ToJson() ([]byte, error) {
	b, err := util.ToJson(q.myMap)
	return b, err
}
