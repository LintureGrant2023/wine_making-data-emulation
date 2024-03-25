package utils

import (
	"backend/gobal"
	"context"
	"fmt"
	"os"
	"time"

	influxdb "github.com/influxdata/influxdb-client-go/v2"
)

// influxdb client config
type Influxdb2ClientConfig struct {
	Url    string `json:"url,omitempty"`
	Org    string `json:"org,omitempty"`
	Bucket string `json:"bucket,omitempty"`
}

// influxdb data config
type Influxdb2DataConfig struct {
	Measurement string            `json:"measurement,omitempty"`
	Tag         map[string]string `json:"tag,omitempty"`
	FieldKey    string            `json:"fieldKey,omitempty"`
}

type DataBaseConfig struct {
	Influxdb2ClientConfig *Influxdb2ClientConfig `json:"influxdb2ClientConfig,omitempty"`
	Influxdb2DataConfig   *Influxdb2DataConfig   `json:"influxdb2DataConfig,omitempty"`
}

type DataModel struct {
	SensorData SensorData
}

func NewDataBaseClient() (*DataBaseConfig, error) {
	// parse influx database config data
	influxdb2ClientConfig := Influxdb2ClientConfig{
		Url:    gobal.Config.Influxdb.GetUrl(),
		Org:    gobal.Config.Influxdb.Organization,
		Bucket: gobal.Config.Influxdb.Bucket,
	}
	influxdb2DataConfig := Influxdb2DataConfig{
		Measurement: "test-write",
		Tag:         map[string]string{"type": "Reported"},
		FieldKey:    "co2",
	}

	return &DataBaseConfig{
		Influxdb2ClientConfig: &influxdb2ClientConfig,
		Influxdb2DataConfig:   &influxdb2DataConfig,
	}, nil
}

func (d *DataBaseConfig) InitClient() influxdb.Client {
	var usrtoken string
	usrtoken = os.Getenv("TOKEN")
	fmt.Println("TOKEN: ", usrtoken)
	client := influxdb.NewClient(d.Influxdb2ClientConfig.Url, gobal.Config.Influxdb.Token)
	//fmt.Println(d.Influxdb2ClientConfig.Url, gobal.Config.Influxdb.Token)
	return client
}

func (d *DataBaseConfig) CloseSession(client influxdb.Client) {
	client.Close()
}

func (d *DataBaseConfig) AddData(data *DataModel, client influxdb.Client) error {
	// write device data to influx database
	writeAPI := client.WriteAPIBlocking(d.Influxdb2ClientConfig.Org, d.Influxdb2ClientConfig.Bucket)
	//fmt.Println(d.Influxdb2ClientConfig.Org, d.Influxdb2ClientConfig.Bucket)
	p := influxdb.NewPoint(d.Influxdb2DataConfig.Measurement,
		d.Influxdb2DataConfig.Tag,
		map[string]interface{}{d.Influxdb2DataConfig.FieldKey: data.SensorData.CO2},
		time.Now())
	// write point immediately
	err := writeAPI.WritePoint(context.Background(), p)
	if err != nil {
		return err
	}
	return nil
}

func NewClient() influxdb.Client {
	token := "JVPfR4sWY3f9W93tJYaN_2lVfdozqSKMuxZpU2WlFUKuDntsdO9HcPLaqBcTOZ1Pw1_p3srVM0VSpphRkPFegQ=="
	url := "http://172.17.0.2:8086"
	client := influxdb.NewClient(url, token)
	os.Getenv("TOKEN")
	return client
}
