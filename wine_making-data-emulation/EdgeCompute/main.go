package main

import (
	"data-emulation/utils"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// 初始化的 节点数量
const Init_num = 1

// 初始化数据的个数
const Init_count = 20

// 初始化时间间隔
const Init_interval = 1

// 正常仿真的 节点数量
const Normal_num = 1200

// 正常仿真数据间隔
const Normal_interval = 5

// 每隔一秒，向influxdb写入一条数据
func main() {
	//打开文件
	utils.OpenAllFiles()
	fmt.Println("************************************************************************* [EdgeNode: v1.0] ************************************************************************")
	fmt.Println("连接到系统...")
	//连数据库
	utils.ConnectMysql()

	//获取influxdb客户端
	utils.NewClient()

	//测试初始仿真程序的执行时间
	fmt.Println("开始初始阶段的数据仿真...")
	// start := time.Now()
	utils.InitData(Init_num, Init_count, Init_interval) //间隔单位为s
	// duration := time.Since(start)
	// fmt.Println("初始化仿真程序共用时: ", duration)

	//正常情况下（非初始状态），产生仿真数据
	//fmt.Println("开始正常阶段的数据仿真...")
	fmt.Printf("开始连接到所有反应器...\n\n")
	//fmt.Printf("开始连接到[%d]个反应器...\n\n", Normal_num)

	utils.NormalData(Normal_num, Init_count, Normal_interval) //间隔单位为min

	//断开mysql数据库
	utils.CloseMysql()

	//断开influxdb
	utils.CloseClient()
}
