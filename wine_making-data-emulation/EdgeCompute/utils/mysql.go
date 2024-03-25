package utils

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"
)

var DB *sql.DB

func ConnectMysql() {
	dsn := "root:root@tcp(192.168.120.100:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	//defer db.Close()
	fmt.Println("连接MySQL数据库成功...")
	DB = db
}

func WriteInitMysql(historyDataBatch []SystemHistoryData) {
	// 构建批量插入的SQL语句
	var valueStrings []string
	var valueArgs []interface{}
	for _, data := range historyDataBatch {
		valueStrings = append(valueStrings, "(?,?, ?,?,?, ?,?,?, ?,?,?, ?,?,?, ?,?,?)")
		valueArgs = append(valueArgs, 1, time.Now(),
			data.TemperatureMax, data.TemperatureMin, data.TemperatureMin,
			data.PHMax, data.PHMin, data.PHAvg,
			data.CO2Max, data.CO2Min, data.CO2Avg,
			data.O2Max, data.O2Min, data.O2Avg,
			data.AlcoholMax, data.AlcoholMin, data.AlcoholAvg,
		)
	}
	stmt := fmt.Sprintf("INSERT INTO SystemHistoryData (SystemID, CreateTime, "+
		"TemperatureMax, TemperatureMin, TemperatureAvg,"+
		"PHMax, PHMin, PHAvg,"+
		"CO2Max, CO2Min, CO2Avg,"+
		"O2Max, O2Min, O2Avg,"+
		"AlcoholMax, AlcoholMin, AlcoholAvg"+
		") VALUES %s",
		strings.Join(valueStrings, ","))
	//fmt.Println(stmt)

	// 开始事务
	tx, err := DB.Begin()
	if err != nil {
		log.Fatal(err)
	}

	// 预处理SQL语句
	prepStmt, err := tx.Prepare(stmt)
	if err != nil {
		log.Fatal(err)
	}
	defer prepStmt.Close()

	// 执行SQL语句
	_, err = prepStmt.Exec(valueArgs...)
	if err != nil {
		// 如果发生错误则回滚事务
		tx.Rollback()
		log.Fatal(err)
	}

	// 提交事务
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}

func WriteIntoMysql(data SystemHistoryData) {
	stmt, err := DB.Prepare("INSERT INTO SystemHistoryData (Model, CreateTime, " +
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

	res, err := stmt.Exec(data.Model, time.Now(),
		data.TemperatureMax, data.TemperatureMin, data.TemperatureAvg,
		data.PHMax, data.PHMin, data.PHAvg,
		data.CO2Max, data.CO2Min, data.CO2Avg,
		data.O2Max, data.O2Min, data.O2Avg,
		data.AlcoholMax, data.AlcoholMin, data.AlcoholAvg)
	if err != nil {
		panic(err)
	}

	_, err = res.LastInsertId()
	if err != nil {
		panic(err)
	}
	//fmt.Println("插入MySQL数据库成功, 插入id = ", id)

}

func CloseMysql() {
	DB.Close()
	fmt.Println("断开MySQL数据库连接...")
}
