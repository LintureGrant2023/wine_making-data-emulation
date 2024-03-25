package utils

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"

	influxdb "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
)

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
	SystemID       int
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
	AlcoholMax     float64
	AlcoholMin     float64
	AlcoholAvg     float64
}

func NewClient() influxdb.Client {
	token := "kNRh2ho2bL4TMo41IBE3V2sI45BDNQxLXP2l98jq-ksFr6Q1V5bNsn3aCyi4CytCUdqv-bSesPAXJ514Ufdobg=="
	url := "http://192.168.120.100:8086"
	//url := "http://10.244.0.68:8086"
	//token := "QS-dR2xz5rLlk6Jp5TM32svCDjpDCByQExqAAqaBIU5zPKinz3J4x9LrzOyEIuMkN-RxgJDE89tNiIN3lOZbHQ=="

	client := influxdb.NewClient(url, token)
	return client
}

func WriteData(client influxdb.Client, fields map[string]interface{}, t time.Time) {
	org := "uestcydri"
	bucket := "test"
	//WriteAPI：异步、不加锁的可写客户端
	//WriteAPIBlocking：同步、加锁的可写客户端
	writer := client.WriteAPIBlocking(org, bucket)
	tags := map[string]string{
		//type": "Desired",
		"type": "Reported",
		//方案一：增加tags字段，记录插入操作的时间，不支持按tags排序
		"insert_time": strconv.FormatInt(time.Now().UnixNano(), 10),
	}

	point := write.NewPoint("test-v13", tags, fields, t)

	if err := writer.WritePoint(context.Background(), point); err != nil {
		log.Fatal(err)
	}
}

func QueryData(client influxdb.Client) *api.QueryTableResult {
	//query from influxdb
	//loc, _ := time.LoadLocation("Asia/Shanghai")
	org := "uestcydri"
	queryApi := client.QueryAPI(org)
	query := `from(bucket: "test")
	|> range(start: -24h)
	|> filter(fn: (r) => r._measurement == "test-v13")
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

// provide inited data for webs
func InitData(client influxdb.Client, nums int, hours int, interval int) {
	// loc, _ := time.LoadLocation("Asia/Shanghai")
	// t := time.Now().In(loc)

	for i := hours; i > 0; i-- {
		//写入influxdb，一次只写一组
		fmt.Println(time.Now().Add(-time.Duration(i) * time.Hour))
		WriteData(client, GetSimulationData(nums, hours-i), time.Now().Add(-time.Duration(i)*time.Hour))

		//从influxdb查询数据
		query_res := QueryData(client)

		//处理查询结果
		compute_res := ComputeData(query_res)

		//fmt.Println(compute_res)

		//写入mysql
		WriteIntoMysql(compute_res)

		//休息
		//time.Sleep(time.Duration(interval) * time.Second)
	}

}

func ComputeData(data *api.QueryTableResult) SystemHistoryData {
	fmt.Println("called ComputeData...")
	//res := make([]SystemHistoryData, 100)
	var temp SystemHistoryData
	// temp.TemperatureMax = 0
	// temp.TemperatureMin = math.MaxFloat64
	// temp.TemperatureAvg = 0

	//fmt.Println(data.Next())

	for data.Next() {
		//fmt.Println(data.Record().Value())
		value, _ := ConvertStringToSlice(data.Record().Value().(string))
		//fmt.Println("field = ", data.Record().Field(), " time = ", data.Record().Time(), " value = ", value)
		if data.Record().Field() == "Temperature" {
			temp.TemperatureMax = maxFloats(value)
			temp.TemperatureMin = minFloats(value)
			temp.TemperatureAvg = avgFloats(value)
		}
		if data.Record().Field() == "PH" {
			temp.PHMax = maxFloats(value)
			temp.PHMin = minFloats(value)
			temp.PHAvg = avgFloats(value)
		}
		if data.Record().Field() == "CO2" {
			temp.CO2Max = maxFloats(value)
			temp.CO2Min = minFloats(value)
			temp.CO2Avg = avgFloats(value)
		}
		if data.Record().Field() == "O2" {
			temp.O2Max = maxFloats(value)
			temp.O2Min = minFloats(value)
			temp.O2Avg = avgFloats(value)
		}
		if data.Record().Field() == "Alcohol" {
			temp.AlcoholMax = maxFloats(value)
			temp.AlcoholMin = minFloats(value)
			temp.AlcoholAvg = avgFloats(value)
		}
		//fmt.Println(temp)
		//res = append(res, temp)
	}
	//fmt.Println(temp)
	return temp
}

func WriteIntoMysql(data SystemHistoryData) {
	//write into mysql
	dsn := "root:root@tcp(192.168.120.100:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Println("connect mysql successfully...")
	// TemperatureMax := strconv.FormatFloat(data.TemperatureMax, 'f', -1, 64)
	// TemperatureMin := strconv.FormatFloat(data.TemperatureMin, 'f', -1, 64)
	// TemperatureAvg := strconv.FormatFloat(data.TemperatureAvg, 'f', -1, 64)
	// insert_sql := "insert into SystemHistoryData (" + "TemperatureMax, TemperatureMin, TemperatureAvg" + ") values (" +
	// 	TemperatureMax + "," + TemperatureMin + "," + TemperatureAvg + ")"
	// fmt.Println(insert_sql)
	//stmt, err := db.Prepare(insert_sql)
	stmt, err := db.Prepare("INSERT INTO SystemHistoryData (SystemID, CreateTime, " +
		"TemperatureMax, TemperatureMin, TemperatureAvg," +
		"PHMax, PHMin, PHAvg," +
		"CO2Max, CO2Min, CO2Avg," +
		"O2Max, O2Min, O2Avg," +
		"AlcoholMax, AlcoholMin, AlcoholAvg" +
		") VALUES (?,?, ?,?,?, ?,?,?, ?,?,?, ?,?,?, ?,?,?)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(1, time.Now(),
		data.TemperatureMax, data.TemperatureMin, data.TemperatureAvg,
		data.PHMax, data.PHMin, data.PHAvg,
		data.CO2Max, data.CO2Min, data.CO2Avg,
		data.O2Max, data.O2Min, data.O2Avg,
		data.AlcoholMax, data.AlcoholMin, data.AlcoholAvg)
	if err != nil {
		panic(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Println("insert successfully...,   id = ", id)
}
