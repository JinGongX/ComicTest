package routers

import (
	"gocmictest/middleware/gcors"
	"gocmictest/pkg/setting"
	v1 "gocmictest/routers/api/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(gcors.Cors())
	gin.SetMode(setting.RunMode)
	r.Static("/files", "./files")
	r.POST("/api/upload", v1.Uploadimgfile)
	r.POST("/api/SendComic", v1.SendComic)

	return r
}
