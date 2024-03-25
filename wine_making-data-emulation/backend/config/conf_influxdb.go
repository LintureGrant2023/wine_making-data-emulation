package config

type Influxdb struct {
	Ip           string `json:"ip" yaml:"ip"`
	Port         string `json:"port" yaml:"port"`
	Organization string `json:"organization" yaml:"organization"`
	Bucket       string `json:"bucket" yaml:"bucket"`
	Token        string `json:"token" yaml:"token"`
}

func (influxdb Influxdb) GetUrl() string {
	return "http://" + influxdb.Ip + ":" + influxdb.Port
}
