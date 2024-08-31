package api

import (
	"anime_server/middleware"
	"github.com/gin-gonic/gin"
)

func InitRoute() {
	r := gin.Default()
	r.Use(middleware.CORS())
	// 上传视频接口
	r.POST("/uploadVideo", UpLoadVideo)
	r.POST("/login",Login)
	r.POST("/regist",Regist)
	InfoGroup := r.Group("/info")
	{
		InfoGroup.Use(middleware.JWTAuthMiddleware())
		InfoGroup.GET("/getInfos", GetInfos)
	}
	// 提供 HLS 文件
	r.Static("/hls", "./hls")
	r.Run(":1226")
}
