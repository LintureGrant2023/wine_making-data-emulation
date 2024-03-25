package utils

import (
	"backend/gobal"
	"database/sql"
	"fmt"
)

func ConnectMysql() {
	dsn := gobal.Config.MySQL.GetDB()
	fmt.Println(dsn)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	//err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("成功连接到MySQL数据库")
}
