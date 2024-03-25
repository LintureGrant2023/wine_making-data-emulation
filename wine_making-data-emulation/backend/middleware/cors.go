package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// 解决跨域共享资源问题
func InitCors() gin.HandlerFunc {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://192.168.163.128:8080"}
	corsConfig.AllowMethods = []string{"GET", "POST"}
	corsConfig.AllowHeaders = []string{"Origin", "Authorization", "Access-Control-Allow-Origin"}

	corsMiddleware := cors.New(corsConfig)
	return corsMiddleware
}
