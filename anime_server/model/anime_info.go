package model

type Anime struct {
	Year         int    `gorm:"not null"`
	Name         string `gorm:"size:255;not null"`
	Season       string `gorm:"not null"`
	EpisodeCount int    `gorm:"not null"`
	Author       string `gorm:"size:255;not null"`
	Path         string `gorm:"size:255;not null"`
	Cover        string `gorm:"size:255;not null"`
}
