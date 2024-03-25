package main

import (
	"data-emulation/utils"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// 每隔一秒，向influxdb写入一条数据
func main() {
	//fmt.Println(time.Now())
	//fmt.Println(utils.InitO2(0))
	client := utils.NewClient()
	//shanghaiT, _ := time.LoadLocation("Asia/Shanghai")
	//duration time about ?s
	start := time.Now()
	utils.InitData(client, 100, 24, 1)
	end := time.Since(start)
	fmt.Println("初始化仿真程序共用时: ", end)
	//utils.GenerateCO2()

	//WriteIntoMysql(compute_res)

	//间隔时间产生
	// for {
	// 	//fmt.Println(time.Now().In(shanghaiT).Format("2006-01-02 15:04:05"))
	// 	//产生仿真数据
	// 	WriteData(client, GetSimulationData(1000), time.Now().Round(time.Second))
	// 	time.Sleep(5 * 60 * time.Second) // separate points by 1 second
	// }

}
