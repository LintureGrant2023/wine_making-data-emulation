package gobal

import (
	"backend/config"

	"github.com/gin-gonic/gin"
)

var (
	Config *config.Config
	Cors   gin.HandlerFunc
)
