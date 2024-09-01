package dao

import (
	"anime_server/model"
	"gorm.io/gorm"
)

func SaveAnimeInfo(anime *model.Anime) error {
	var existingAnime model.Anime
	err := DB.Where("name = ? AND season = ?", anime.Name, anime.Season).First(&existingAnime).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// 记录不存在，创建新的记录
			if err := DB.Create(anime).Error; err != nil {
				return err
			}
		} else {
			// 处理其他查询错误
			return err
		}
	} else {
		DB.Exec("update animes set episode_count = episode_count + 1 where name = ? and season = ?", anime.Name, anime.Season)
	}
	return nil
}
func GetAnimes(anime_Info *[]model.Anime) error {
	//后续数量多了使用分页查询
	return DB.Find(anime_Info).Error
}
func GetInfoById(id int, anime_Info *[]model.Anime) error {
	return DB.Where("id = ?", id).First(anime_Info).Error
}
