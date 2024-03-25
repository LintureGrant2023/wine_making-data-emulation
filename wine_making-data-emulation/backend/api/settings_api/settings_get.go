package settings_api

import (
	"backend/config"
	"backend/core"
	"backend/gobal"
	"backend/models/res"
	"fmt"
	"log"
	"reflect"

	"github.com/gin-gonic/gin"
)

func (SettingsApi) GetAllThresholdView(c *gin.Context) {
	//处理预检请求
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(200)
		return
	}
	res.OKWithData(gobal.Config.Sensor, c)
}

func (SettingsApi) SetThresholdView(c *gin.Context) {
	var threshold_data config.Sensor
	//只能接受json格式
	err := c.ShouldBind(&threshold_data)
	if err != nil {
		//log.Fatal(err.Error())
		fmt.Println("失败")
		res.Error("失败", c)
		return
	}

	//fmt.Println(threshold_data)
	v := reflect.ValueOf(&threshold_data).Elem()
	gobalV := reflect.ValueOf(&gobal.Config.Sensor).Elem()
	// 遍历结构体字段
	for i := 0; i < v.NumField(); i++ {

		// 获取这个字段值
		value := v.Field(i)
		// fmt.Println(value.String() != "")
		// 判断是否为空
		if value.String() != "" {
			// 不为空,则重新赋值
			//fmt.Printf("%v\n", v.Field(i).Interface())
			//fmt.Println(gobalV.Field(i))
			gobalV.Field(i).Set(reflect.ValueOf(value.Interface()))
		}
	}
	//修改配置文件
	err = core.SetConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

	res.OKWithMsg("成功", c)
}
