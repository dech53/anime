package api

import (
	"anime_server/dao"
	"anime_server/model"
	"anime_server/utils"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func UpLoadVideo(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.String(400, "Failed to get file: %v", err)
		return
	}
	// 保存上传的文件
	filename := filepath.Base(file.Filename)
	// 获取原始文件的扩展名
	ext := filepath.Ext(filename)
	animeInfo := strings.Split(filename, "-")
	savePath := filepath.Join("./uploads", animeInfo[0], animeInfo[1], animeInfo[2], animeInfo[3]+ext)
	if err := c.SaveUploadedFile(file, savePath); err != nil {
		c.String(500, "Failed to save file: %v", err)
		return
	}
	// 转码为 HLS 格式
	hlsDir := filepath.Join("./hls", animeInfo[0], animeInfo[1], animeInfo[2], animeInfo[3])
	m3u8Filename := "output.m3u8"
	err = os.MkdirAll(hlsDir, 0755)
	if err != nil {
		log.Fatalf("Failed to create directory: %v", err)
	}
	cmd := exec.Command("ffmpeg", "-i", savePath, "-c:v", "libx264", "-c:a", "aac", "-start_number", "0", "-hls_time", "10", "-hls_list_size", "0", "-f", "hls", filepath.Join(hlsDir, m3u8Filename))
	// 获取标准输出和标准错误
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("FFmpeg error: %v. Output: %s", err, string(output))
		c.String(500, "Failed to convert video: %v", err)
		return
	}
	year, _ := strconv.Atoi(animeInfo[0])
	dao.SaveAnimeInfo(&model.Anime{
		Year:         year,
		Name:         animeInfo[1],
		Season:       animeInfo[2],
		EpisodeCount: 1,
		Author:       strings.Split(animeInfo[4], ".")[0],
		Path:         filepath.Join(hlsDir, m3u8Filename),
	})
	c.String(200, "File uploaded and converted successfully")
}
func GetInfos(c *gin.Context) {
	var anime_Infos []model.Anime
	err := dao.GetAnimes(&anime_Infos)
	if err != nil {
		utils.ResponseFail(c, "获取信息失败", 404)
		return
	}
	utils.ResponseSuccess(c, anime_Infos, 200)
}
