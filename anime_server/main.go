package main

import (
	"anime_server/api"
	"anime_server/dao"
)

func main() {
	dao.Initminio()
	dao.InitDB()
	api.InitRoute()
}
