package utils

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	influxdb "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
)

var client influxdb.Client

// 填充数据点
var dataPoints []*write.Point

// 传感器数据结构体
type SensorData struct {
	Temperature float64
	O2          float64
	CO2         float64
	PH          float64
	Alcohol     float64
}

type Data struct {
	ReactorID  int
	EdgeID     int
	SystemID   int
	Model      string
	SensorData SensorData
}

type SystemHistoryData struct {
	Model          int
	CreateTime     time.Time
	TemperatureMax float64
	TemperatureMin float64
	TemperatureAvg float64
	PHMax          float64
	PHMin          float64
	PHAvg          float64
	O2Max          float64
	O2Min          float64
	O2Avg          float64
	CO2Max         float64
	CO2Min         float64
	CO2Avg         float64
	AlcoholMax     int
	AlcoholMin     int
	AlcoholAvg     int
}

func NewClient() {
	token := "kNRh2ho2bL4TMo41IBE3V2sI45BDNQxLXP2l98jq-ksFr6Q1V5bNsn3aCyi4CytCUdqv-bSesPAXJ514Ufdobg=="
	url := "http://192.168.120.100:8086"
	//url := "http://10.244.0.68:8086"
	//token := "QS-dR2xz5rLlk6Jp5TM32svCDjpDCByQExqAAqaBIU5zPKinz3J4x9LrzOyEIuMkN-RxgJDE89tNiIN3lOZbHQ=="

	client = influxdb.NewClient(url, token)
	fmt.Println("连接Influxdb数据库成功...")
}

func GenerateDataPoint(fields map[string]interface{}, t time.Time) {
	//start := time.Now()

	tags := map[string]string{
		//type": "Desired",
		"type": "Reported",
		//方案一：增加tags字段，记录插入操作的时间，不支持按tags排序
		"insert_time": strconv.FormatInt(time.Now().UnixNano(), 10),
	}
	point := write.NewPoint("test-v17", tags, fields, t)

	org := "uestcydri"
	bucket := "test"
	//WriteAPI：异步、不加锁的可写客户端
	//WriteAPIBlocking：同步、加锁的可写客户端
	writer := client.WriteAPI(org, bucket)

	// 写入数据点
	writer.WritePoint(point)

	// 一次性刷新所有缓冲区，并关闭客户端
	writer.Flush()
	//fmt.Println("GenerateDataPoint用时: ", time.Since(start))
	//return point
}

func WriteInitInfluxdb() {
	org := "uestcydri"
	bucket := "test"
	//WriteAPI：异步、不加锁的可写客户端
	//WriteAPIBlocking：同步、加锁的可写客户端
	writer := client.WriteAPI(org, bucket)

	// 写入数据点
	for _, point := range dataPoints {
		writer.WritePoint(point)
	}

	// 一次性刷新所有缓冲区，并关闭客户端
	writer.Flush()
}

func WriteIntoInfluxdb(flag int, fields map[string]interface{}, t time.Time) {
	org := "uestcydri"
	bucket := "test"
	tags := map[string]string{
		//type": "Desired",
		"type": "Reported",
		//方案一：增加tags字段，记录插入操作的时间，不支持按tags排序
		"insert_time": strconv.FormatInt(time.Now().UnixNano(), 10),
	}

	writer := client.WriteAPI(org, bucket)
	point := write.NewPoint("test-v17", tags, fields, t)
	//WriteAPI：异步、不加锁的可写客户端
	//WriteAPIBlocking：同步、加锁的可写客户端
	writer.WritePoint(point)
	writer.Flush()
	// if err := writer.WritePoint(context.Background(), point); err != nil {
	// 	log.Fatal(err)
	// }
	if flag == 2 {
		fmt.Println("写入Influxdb数据库成功...")
	}

}

func QueryData() *api.QueryTableResult {
	//query from influxdb
	//loc, _ := time.LoadLocation("Asia/Shanghai")
	org := "uestcydri"
	bucket := "test"
	queryApi := client.QueryAPI(org)
	query := `from(bucket:` + bucket + `)
	|> range(start: -24h)
	|> filter(fn: (r) => r._measurement == "test-v17")
	|> last()`
	res, err := queryApi.Query(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}
	// for res.Next() {
	// 	fmt.Println(res.Record().Field(), res.Record().Time().In(loc).Round(time.Second).Format("2006-01-02 15:04:05"), res.Record().Value())
	// }
	if err := res.Err(); err != nil {
		log.Fatal(err)
	}
	return res
}

func CloseClient() {
	client.Close()
	fmt.Println("断开Influxdb数据库连接...")
}
