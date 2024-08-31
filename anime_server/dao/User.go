package dao

import "anime_server/model"

func GetInfoByPattern(pattern, value string) (model.User, error) {
	var user model.User
	err := DB.Where(pattern+" = ?", value).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
func RegistUser(user model.User)error{
	return DB.Save(&user).Error
}