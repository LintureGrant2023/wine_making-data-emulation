package utils

import (
	"fmt"
	"time"
)

func InitData(nums int, hours int, interval float64) {

	for i := hours; i > 0; i-- {
		WriteIntoInfluxdb(1, InitSimulationData(nums, hours-i), time.Now().Add(-time.Duration(i)*time.Hour))
		ProcessInitData()
		//time.Sleep(time.Duration(interval) * time.Second)

		//插入异常数据
		if i%3 != 0 {
			GenerateAbnormal()
		}
	}
}

func NormalData(nums int, start int, interval int) {
	//fmt.Println("开始检测传感器...")
	for i := start; true; i++ {
		//每五分钟 产生仿真数据
		//GenerateData(nums, i, 1)
		fmt.Printf("开始第%d次采集数据...\n", i-20+1)
		start := time.Now()
		WriteIntoInfluxdb(2, GetSimulationData(nums, i), time.Now())
		GenerateData()
		duration := time.Since(start)
		fmt.Printf("系统数据采集&呈现总时间: 一共采集了%d个反应器,用时 %v\n", nums, duration)
		fmt.Println("前端呈现完成...")
		//fmt.Println("检测完成，传感器参数种类分别为: 温度, pH值, 酒精气体浓度, 二氧化碳浓度, 氧气浓度")
		fmt.Println()
		//fmt.Println("正常阶段仿真一次数据用时: ", duration)
		time.Sleep(time.Duration(interval) * 60 * time.Second) // separate points by 1 second
		if i%3 == 0 {
			GenerateAbnormal()
		}
	}
}

// provide inited data for webs
func GenerateData() {
	//从influxdb查询数据
	query_res := QueryData()

	//处理查询结果
	compute_res := ComputeNormalData(query_res)

	//写入mysql
	WriteIntoMysql(compute_res)

}

// provide inited data for webs
func ProcessInitData() {
	//start1 := time.Now()

	//从influxdb查询数据
	query_res := QueryData()

	//处理查询结果
	compute_res := ComputeInitData(query_res)

	//写入mysql
	WriteIntoMysql(compute_res)

	//fmt.Println("ProcessInitData用时: ", time.Since(start1))
}
