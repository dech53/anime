package api

import (
	"anime_server/dao"
	"anime_server/model"
	"anime_server/utils"
	"context"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
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
func GetInfoById(c *gin.Context) {
	var anime_Info []model.Anime
	anime_id := c.Param("id")
	id, _ := strconv.Atoi(anime_id)
	err := dao.GetInfoById(id, &anime_Info)
	if err != nil {
		utils.ResponseFail(c, "获取信息失败", 404)
		return
	}
	utils.ResponseSuccess(c, anime_Info, 200)
}
func UploadVideoToMinIO(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.String(400, "Failed to get file: %v", err)
		return
	}
	animeInfo := strings.Split(file.Filename, "-")
	// 构建保存路径
	ext := filepath.Ext(file.Filename)
	localSavePath := filepath.Join("./uploads", animeInfo[0], animeInfo[1], animeInfo[2], animeInfo[3]+ext)
	c.SaveUploadedFile(file, localSavePath)
	hlsDir := filepath.Join("./hls", animeInfo[0], animeInfo[1], animeInfo[2], animeInfo[3])
	m3u8Filename := "output.m3u8"
	os.MkdirAll(hlsDir, 0755)
	cmd := exec.Command("ffmpeg", "-hwaccel", "cuvid", "-i", localSavePath, "-c:v", "h264_nvenc", "-c:a", "aac", "-start_number", "0", "-hls_time", "10", "-hls_list_size", "0", "-f", "hls", filepath.Join(hlsDir, m3u8Filename))
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("FFmpeg error: %v. Output: %s", err, string(output))
		c.String(500, "Failed to convert video: %v", err)
		return
	}
	hlsFiles, err := filepath.Glob(filepath.Join(hlsDir, "*"))
	for _, filePath := range hlsFiles {
		fileName := filepath.Base(filePath)
		objectName := "/hls" + "/" + animeInfo[0] + "/" + animeInfo[1] + "/" + animeInfo[2] + "/" + animeInfo[3] + "/" + fileName
		// 打开文件
		file, err := os.Open(filePath)
		if err != nil {
			log.Fatalf("Failed to open HLS file: %v", err)
		}
		defer file.Close()
		// 获取文件信息
		fileInfo, err := file.Stat()
		if err != nil {
			log.Fatalf("Failed to get file info: %v", err)
		}
		// 上传文件到 MinIO
		_, err = dao.MinioClient.PutObject(context.Background(), "test", objectName, file, fileInfo.Size(), minio.PutObjectOptions{
			ContentType: "application/octet-stream",
		})

		if err != nil {
			log.Fatalf("Failed to upload file to MinIO: %v", err)
		}
		file.Close()
	}
	os.RemoveAll("./hls")
	os.RemoveAll("./uploads")
	year, _ := strconv.Atoi(animeInfo[0])
	dao.SaveAnimeInfo(&model.Anime{
		Year:         year,
		Name:         animeInfo[1],
		Season:       animeInfo[2],
		EpisodeCount: 1,
		Author:       strings.Split(animeInfo[4], ".")[0],
		Path:         filepath.Join(hlsDir, m3u8Filename),
	})
	utils.ResponseSuccess(c, "上传文件成功", 200)
}
func GetM3U8(c *gin.Context) {
	year := c.Param("year")
	name := c.Param("name")
	season := c.Param("season")
	episode := c.Param("episode")
	filename := c.Param("filename")
	presignedURL, _ := dao.MinioClient.PresignedGetObject(context.Background(), "test", "hls/"+year+"/"+name+"/"+season+"/"+episode+"/"+filename, time.Minute*10, nil)
	log.Println(presignedURL)
	utils.ResponseSuccess(c, presignedURL.String(), 200)
}
