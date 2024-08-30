package main

import (
	"anime_server/api"
	"anime_server/dao"
)

func main() {
	dao.InitDB()
	api.InitRoute()
}
