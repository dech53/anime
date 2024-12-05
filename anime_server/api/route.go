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
	//上传视频至minIO
	r.POST("/uploadVideoToMinIO",UploadVideoToMinIO)
	r.POST("/login",Login)
	r.POST("/regist",Regist)
	//获取播放源
	InfoGroup := r.Group("/info")
	{
		InfoGroup.Use(middleware.JWTAuthMiddleware())
		InfoGroup.GET("/getInfos", GetInfos)
		InfoGroup.GET("/getInfoById/:id",GetInfoById)
	}
	// 提供 HLS 文件
	r.GET("/hls/:year/:name/:season/:episode/:filename", GetM3U8)
	r.Run(":1226")
}
