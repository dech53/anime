package utils

//响应工具类
import (
	"anime_server/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ResponseSuccess(c *gin.Context, data interface{}, code int) {
	c.JSON(http.StatusOK, model.Response{
		Code:    code,
		Message: "成功",
		Result:  data,
	})
}
func ResponseFail(c *gin.Context, data interface{}, code int) {
	c.JSON(http.StatusOK, model.Response{
		Code:    code,
		Message: "失败",
		Result:  data,
	})
}
