package routers

import (
	"backend/api"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	//要在配置路由之前，使用cors中间件 ==> 中间件的使用，写死了AllowOrigins，其他ip地址发来的请求都会被forbidden
	//cors := middleware.InitCors()
	//router.Use(cors)

	//## 系统配置接口 ##
	settingsApi := api.ApiGroupApp.SettingsApi
	//获取、更新配置
	router.GET("/api/settings", settingsApi.GetAllThresholdView)
	router.POST("/api/settings", settingsApi.SetThresholdView)

	//## dashboard接口 ##
	dashboardApi := api.ApiGroupApp.DashboardApi
	//获取、更新传感器数据
	router.GET("/api/dashboard/sensor", dashboardApi.GetSensor)
	router.GET("/api/dashboard/sensor-update", dashboardApi.UpdateSensor)

	//获取温度
	router.GET("/api/dashboard/temperature", dashboardApi.GetTemp)
	//获取ph
	router.GET("/api/dashboard/ph", dashboardApi.GetPh)
	//获取酒精浓度
	router.GET("/api/dashboard/alcohol", dashboardApi.GetAlcohol)
	//获取氧气浓度
	router.GET("/api/dashboard/o2", dashboardApi.GetO2)
	//获取二氧化碳浓度
	router.GET("/api/dashboard/co2", dashboardApi.GetCO2)
	//更新上面五个属性
	router.GET("/api/dashboard/all-update", dashboardApi.UpdateAll)

	//获取、更新环境数据
	router.GET("/api/dashboard/env", dashboardApi.GetEnv)
	router.GET("/api/dashboard/env-update", dashboardApi.UpdateEnv)

	//获取预测的温度
	router.GET("/api/dashboard/perdict-temp", dashboardApi.GetPerdictedTemp)
	//获取预测的ph
	router.GET("/api/dashboard/perdict-ph", dashboardApi.GetPerdictedPh)
	//更新预测数据
	router.GET("/api/dashboard/perdict-update", dashboardApi.UpdatePerdicted)

	//获取、更新异常信息
	router.GET("/api/dashboard/abnorm", dashboardApi.GetAbnorm)
	router.GET("/api/dashboard/abnorm-update", dashboardApi.UpdateAbnorm)

	//## 全局控制接口 ##
	globalctlApi := api.ApiGroupApp.GlobalctlApi
	//获取、更新"状态显示"
	router.GET("/api/globalctl/reactor", globalctlApi.GetReactor)
	router.GET("/api/globalctl/reactor-update", globalctlApi.UpdateReactor)

	//获取、更新传感器（仪表盘）
	router.GET("/api/globalctl/gauge", globalctlApi.GetGauge)
	router.GET("/api/globalctl/gauge-update", globalctlApi.UpdateGauge)

	//获取、更新 区域反应堆状态显示
	router.GET("/api/globalctl/sensor", globalctlApi.GetSensor)
	router.GET("/api/globalctl/sensor-update", globalctlApi.UpdateSensor)

	//获取、更新传感器告警信息
	router.GET("/api/globalctl/warnings", globalctlApi.GetWarnings)
	router.GET("/api/globalctl/warnings-update", globalctlApi.UpdateWarnings)

	//获取、更新预测数据
	router.POST("/api/globalctl/predict", globalctlApi.GetPerdictedSensor)
	router.POST("/api/globalctl/predict-update", globalctlApi.UpdatePerdicted)

	//## 局部控制接口 ##
	localctlApi := api.ApiGroupApp.LocalctlApi
	//获取、更新反应器状态
	router.GET("/api/localctl/reactor", localctlApi.GetReactor)
	router.GET("/api/localctl/reactor-update", localctlApi.UpdateReactor)

	//获取、更新"传感器数据"
	router.GET("/api/localctl/history-sensor", localctlApi.GetHistorySensor)
	router.GET("/api/localctl/current-sensor", localctlApi.GetCurrentSensor)

	//反应器控制 -- 温度控制
	router.POST("/api/localctl/start-heat", localctlApi.StartHeat)
	router.POST("/api/localctl/stop-heat", localctlApi.StopHeat)
	router.POST("/api/localctl/auto-heat", localctlApi.AutoHeat)

	//反应器控制 -- 氧气浓度控制
	router.POST("/api/localctl/start-ventilate", localctlApi.StartVentilate)
	router.POST("/api/localctl/stop-ventilate", localctlApi.StopVentilate)
	router.POST("/api/localctl/auto-ventilate", localctlApi.AutoVentilate)

	//获取、更新告警日志
	router.GET("/api/localctl/warnings", localctlApi.GetWarnings)
	router.GET("/api/localctl/warnings-update", localctlApi.UpdateWarnings)

	//控制流
	controlApi := api.ApiGroupApp.ControlApi
	router.POST("/api/control/all", controlApi.ControlALL)

	return router
}
