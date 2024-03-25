package main

import (
	"backend/core"
	"backend/gobal"
	"backend/routers"
	"backend/utils"

	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("******backend: v2.7, apiversios: v1beta1, msg: 修改了接口******")
	gobal.Config = core.InitConfig()
	utils.ConnectMysql()
	router := routers.InitRouter()
	router.Run(":8089")

}
