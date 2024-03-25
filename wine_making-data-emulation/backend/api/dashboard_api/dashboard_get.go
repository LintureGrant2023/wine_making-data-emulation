package dashboard_api

import (
	"backend/models/res"
	"backend/utils"

	"github.com/gin-gonic/gin"
)

// 获取温度数据
func (DashboardApi) GetTemp(c *gin.Context) {
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(200)
		return
	}
	data := utils.GetTemps(24)

	//前端没有设置代理，则需要在后端设置接受跨域
	// c.Header("Access-Control-Allow-Origin", "http://192.168.120.100:8089")
	res.OKWithData(data, c)
}

// 获取PH数据
func (DashboardApi) GetPh(c *gin.Context) {
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(200)
		return
	}
	data := utils.GetPhs(24)
	res.OKWithData(data, c)
}

// 获取二氧化碳浓度数据
func (DashboardApi) GetCO2(c *gin.Context) {
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(200)
		return
	}
	data := utils.GetCO2()
	res.OKWithData(data, c)
}

// 获取氧气浓度数据
func (DashboardApi) GetO2(c *gin.Context) {
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(200)
		return
	}
	data := utils.GetO2()
	res.OKWithData(data, c)
}

// 获取酒精浓度数据
func (DashboardApi) GetAlcohol(c *gin.Context) {
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(200)
		return
	}
	data := utils.GetAlcohols(24)
	res.OKWithData(data, c)
}

// 获取传感器信息数据
func (DashboardApi) GetSensor(c *gin.Context) {
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(200)
		return
	}
	data := utils.GetSensorData()
	res.OKWithData(data, c)
}

// 更新传感器信息数据
func (DashboardApi) UpdateSensor(c *gin.Context) {
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(200)
		return
	}
	data := utils.GetSensorData()
	res.OKWithData(data, c)
}

// 获取温度预测数据
func (DashboardApi) GetPerdictedTemp(c *gin.Context) {
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(200)
		return
	}
	data := utils.GetPerdictedTemps(12)
	res.OKWithData(data, c)
}

// 获取PH预测数据
func (DashboardApi) GetPerdictedPh(c *gin.Context) {
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(200)
		return
	}
	data := utils.GetPerdictedPhs(12)
	res.OKWithData(data, c)
}

// 获取异常信息
func (DashboardApi) GetAbnorm(c *gin.Context) {
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(200)
		return
	}
	data := utils.GetAbnorms()
	res.OKWithData(data, c)
}

// 更新异常信息
func (DashboardApi) UpdateAbnorm(c *gin.Context) {
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(200)
		return
	}
	data := utils.GetAbnorms()
	res.OKWithData(data, c)
}

// 更新五个属性数据
func (DashboardApi) UpdateAll(c *gin.Context) {
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(200)
		return
	}
	data := utils.GetCurrentAll()
	res.OKWithData(data, c)
}

// 获取环境数据
func (DashboardApi) GetEnv(c *gin.Context) {
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(200)
		return
	}
	data := utils.GetEnv()
	res.OKWithData(data, c)
}

// 更新环境数据
func (DashboardApi) UpdateEnv(c *gin.Context) {
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(200)
		return
	}
	data := utils.GetEnv()
	res.OKWithData(data, c)
}

// 更新预测数据
func (DashboardApi) UpdatePerdicted(c *gin.Context) {
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(200)
		return
	}
	data := utils.GetPerdictUpdated()
	res.OKWithData(data, c)
}
