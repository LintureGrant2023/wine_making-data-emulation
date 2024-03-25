package globalctl_api

import (
	"backend/models/res"
	"backend/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

var Pos Position

type Position struct {
	Row    int `json:"row"`
	Column int `json:"col"`
}

// 获取环境信息
func (GlobalctlApi) GetReactor(c *gin.Context) {
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(200)
		return
	}
	data := utils.GetReactorStatus()
	res.OKWithData(data, c)
}

// 更新环境信息
func (GlobalctlApi) UpdateReactor(c *gin.Context) {
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(200)
		return
	}
	data := utils.GetReactorStatus()
	res.OKWithData(data, c)
}

// 获取仪表盘数据
func (GlobalctlApi) GetGauge(c *gin.Context) {
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(200)
		return
	}
	data := utils.GetSensorData()
	res.OKWithData(data, c)
}

// 更新仪表盘数据
func (GlobalctlApi) UpdateGauge(c *gin.Context) {
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(200)
		return
	}
	data := utils.GetSensorData()
	res.OKWithData(data, c)
}

// 获取反应堆数据
func (GlobalctlApi) GetSensor(c *gin.Context) {
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(200)
		return
	}
	data := utils.GetSensorStatus()
	res.OKWithData(data, c)
}

// 更新反应堆数据
func (GlobalctlApi) UpdateSensor(c *gin.Context) {
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(200)
		return
	}
	data := utils.GetSensorStatus()
	res.OKWithData(data, c)
}

// 获取预测数据
func (GlobalctlApi) GetPerdictedSensor(c *gin.Context) {
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(200)
		return
	}
	//var pos Position
	if err := c.ShouldBindJSON(&Pos); err != nil {
		res.Error("绑定json失败", c)
		return
	}
	fmt.Println("pos = ", Pos)

	data := utils.GetPerdictedSensor(12)
	res.OKWithData(data, c)
}

// 更新预测数据
func (GlobalctlApi) UpdatePerdicted(c *gin.Context) {
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(200)
		return
	}
	var pos Position
	if err := c.ShouldBindJSON(&pos); err != nil {
		res.Error("绑定json失败", c)
		return
	}
	fmt.Println("pos = ", pos)

	data := utils.GetUpdatedSensor(12)
	res.OKWithData(data, c)
}

// 获取告警信息
func (GlobalctlApi) GetWarnings(c *gin.Context) {
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(200)
		return
	}
	data := utils.GetAbnorms()
	res.OKWithData(data, c)
}

// 更新告警信息
func (GlobalctlApi) UpdateWarnings(c *gin.Context) {
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(200)
		return
	}
	data := utils.GetAbnorms()
	res.OKWithData(data, c)
}
