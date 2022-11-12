package models

import (
	"time"

	"gorm.io/gorm"
)

type PostsBasic struct {
	Id           int
	Identity     string
	UserIdentity string
	Title        string
	Tags         string
	Content      string
	Mention      string
	Cover        string
	ClickNum     int
	CreatedAt    time.Time      `gorm:"created"`
	UpdatedAt    time.Time      `gorm:"updated"`
	DeletedAt    gorm.DeletedAt `gorm:"deleted"`
}

func (table PostsBasic) TableName() string {
	return "posts_basic"
}
