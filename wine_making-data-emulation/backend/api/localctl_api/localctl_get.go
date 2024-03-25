package localctl_api

import (
	"backend/api/globalctl_api"
	"backend/models/res"
	"backend/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

type HeaterSwitcher struct {
	Swticher int `json:"switcher"`
}

type AutoHeater struct {
	Status         int     `json:"status"`
	UpperThreshold float64 `json:"upper-threshold"`
	LowerThreshold float64 `json:"lower-threshold"`
}

type VentilaterSwitcher struct {
	Swticher int `json:"switcher"`
}

type AutoVentilater struct {
	Status         int     `json:"status"`
	UpperThreshold float64 `json:"upper-threshold"`
	LowerThreshold float64 `json:"lower-threshold"`
}

// 获取反应堆信息
func (LocalctlApi) GetReactor(c *gin.Context) {
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(200)
		return
	}
	id := (globalctl_api.Pos.Row-1)*50 + globalctl_api.Pos.Column
	data := struct {
		Reactor utils.ReactorStatus
		Status  int
		ID      int
	}{
		Reactor: utils.GetReactorStatus(),
		Status:  int(utils.GenerateRandomFloat64Range(0, 1)),
		ID:      id,
	}
	res.OKWithData(data, c)
}

// 更新反应堆信息
func (LocalctlApi) UpdateReactor(c *gin.Context) {
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(200)
		return
	}
	id := (globalctl_api.Pos.Row-1)*50 + globalctl_api.Pos.Column
	data := struct {
		Reactor utils.ReactorStatus
		Status  int
		ID      int
	}{
		Reactor: utils.GetReactorStatus(),
		Status:  int(utils.GenerateRandomFloat64Range(0, 1)),
		ID:      id,
	}
	res.OKWithData(data, c)
}

// 获取传感器数据
func (LocalctlApi) GetHistorySensor(c *gin.Context) {
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(200)
		return
	}
	data := utils.GetHistorySensor()
	res.OKWithData(data, c)
}

// 更新传感器数据
func (LocalctlApi) GetCurrentSensor(c *gin.Context) {
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(200)
		return
	}
	data := utils.GetSensorData()
	res.OKWithData(data, c)
}

// 开始加热
func (LocalctlApi) StartHeat(c *gin.Context) {
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(200)
		return
	}
	var data HeaterSwitcher
	if err := c.ShouldBindJSON(&data); err != nil {
		res.Error("绑定json失败", c)
		return
	}
	fmt.Println("swticher = ", data)
}

// 停止加热
func (LocalctlApi) StopHeat(c *gin.Context) {
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(200)
		return
	}
	var data HeaterSwitcher
	if err := c.ShouldBindJSON(&data); err != nil {
		res.Error("绑定json失败", c)
		return
	}
	fmt.Println("swticher = ", data)
}

// 自动加热
func (LocalctlApi) AutoHeat(c *gin.Context) {
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(200)
		return
	}
	var data AutoHeater
	if err := c.ShouldBindJSON(&data); err != nil {
		res.Error("绑定json失败", c)
		return
	}
	fmt.Println("data = ", data)
}

// 开始通风
func (LocalctlApi) StartVentilate(c *gin.Context) {
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(200)
		return
	}
	var data VentilaterSwitcher
	if err := c.ShouldBindJSON(&data); err != nil {
		res.Error("绑定json失败", c)
		return
	}
	fmt.Println("swticher = ", data)
}

// 停止通风
func (LocalctlApi) StopVentilate(c *gin.Context) {
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(200)
		return
	}
	var data VentilaterSwitcher
	if err := c.ShouldBindJSON(&data); err != nil {
		res.Error("绑定json失败", c)
		return
	}
	fmt.Println("swticher = ", data)
}

// 自动通风
func (LocalctlApi) AutoVentilate(c *gin.Context) {
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(200)
		return
	}
	var data AutoVentilater
	if err := c.ShouldBindJSON(&data); err != nil {
		res.Error("绑定json失败", c)
		return
	}
	fmt.Println("data = ", data)
}

// 获取告警信息
func (LocalctlApi) GetWarnings(c *gin.Context) {
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(200)
		return
	}
	data := utils.GetAbnorms()
	res.OKWithData(data, c)
}

// 更新告警信息
func (LocalctlApi) UpdateWarnings(c *gin.Context) {
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(200)
		return
	}
	data := utils.GetAbnorms()
	res.OKWithData(data, c)
}
