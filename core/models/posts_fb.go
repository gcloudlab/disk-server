package models

import (
	"time"

	"gorm.io/gorm"
)

type PostsFeedback struct {
	Id            int
	Identity      string
	UserIdentity  string
	PostsIdentity string
	Type          string
	Count         int
	Read          int
	CreatedAt     time.Time      `gorm:"created"`
	UpdatedAt     time.Time      `gorm:"updated"`
	DeletedAt     gorm.DeletedAt `gorm:"deleted"`
}

func (table PostsFeedback) TableName() string {
	return "posts_fb"
}
