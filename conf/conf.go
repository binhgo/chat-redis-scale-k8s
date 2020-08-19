package conf

import (
	"encoding/base64"

	jsoniter "github.com/json-iterator/go"
)

type EnvConfig struct {
	Env       string
	Protocol  string
	ConfigMap map[string]string
}

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

func GetEnv() EnvConfig {
	// var e string
	// var p string
	var c map[string]string
	// e = os.Getenv("env")
	// p = os.Getenv("protocol")
	// configStr := os.Getenv("config")
	decoded, _ := base64.URLEncoding.DecodeString("eyJkYkhvc3QiOiIxMDMuMjAuMTUwLjIyNDoyNzAxNyIsImRiVXNlciI6ImFkbWluIiwiZGJQYXNzd29yZCI6IjQzMjE5OTkwMDA5OTkxMjM0In0=")
	FromJson(decoded, &c)

	data := EnvConfig{}
	data.Env = "stg"
	data.Protocol = "HTTP"
	data.ConfigMap = c

	return data
}

type ENV struct {
	Prd string
	Uat string
	Stg string
	Dev string
}

var EnumEnv = &ENV{
	Prd: "prd",
	Uat: "uat",
	Stg: "stg",
	Dev: "dev",
}

const THRIFT = "THRIFT"
const HTTP = "HTTP"
