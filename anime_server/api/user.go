package api

import (
	"anime_server/dao"
	"anime_server/middleware"
	"anime_server/model"
	"anime_server/utils"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var user model.User
	c.ShouldBindJSON(&user)
	saved_user, err := dao.GetInfoByPattern("username", user.Username)
	if err != nil || saved_user.Password != user.Password {
		utils.ResponseFail(c, "登陆失败请检查用户名和密码", 404)
		return
	}
	// 生成 JWT token
	claim := model.MyClaims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Issuer:    "YXH",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, _ := token.SignedString(middleware.Secret)
	// 返回新生成的 token
	utils.ResponseSuccess(c, "Bearer "+tokenString, http.StatusOK)
}
func Regist(c *gin.Context) {
	var new_user model.User
	c.ShouldBindJSON(&new_user)
	saved_user, err := dao.GetInfoByPattern("username", new_user.Username)
	log.Println(saved_user)
	log.Println(new_user)
	if err != nil || saved_user.Username == new_user.Username {
		utils.ResponseFail(c, "用户名已存在", 404)
		return
	}
	err = dao.RegistUser(new_user)
	if err != nil {
		utils.ResponseFail(c, err.Error(), 400)
		return
	}
	utils.ResponseSuccess(c, "注册成功", 200)
}
