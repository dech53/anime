package model

import "github.com/dgrijalva/jwt-go"

type Anime struct {
	ID           int    `gorm:"primaryKey;autoIncrement;unique;not null" json:"id"`
	Year         int    `gorm:"not null"`
	Name         string `gorm:"size:255;not null"`
	Season       string `gorm:"not null"`
	EpisodeCount int    `gorm:"not null"`
	Author       string `gorm:"size:255;not null"`
	Path         string `gorm:"size:255;not null"`
	Cover        string `gorm:"size:255;not null"`
}
type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
