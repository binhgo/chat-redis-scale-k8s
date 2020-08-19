package conf

type BaseConf struct {
	URL                  string
	Timeout              int
	MaxRetry             int
	WaitTimeBetweenRetry int
	LogName              string
	Debug                bool
}
type BMSConf struct {
	BaseConf
	BasicAuth string
}

func GetBMSConf() BMSConf {
	data := BMSConf{}

	data.URL = "172.17.13.23:3000/oms/v1/"
	data.Timeout = 10
	data.MaxRetry = 3
	data.WaitTimeBetweenRetry = 5
	data.LogName = "bms_client"
	data.BasicAuth = "Basic bGFzdG1pbGUtZnJvbnRlbmQ6elVzSjJmOFZ3NHh0eXJ6RXp6RjUyT3diNlU2aTQ3NkM="

	return data
}
