package models

import (
	"time"

	"gorm.io/gorm"
)

type PostsCommentBasic struct {
	Id            int
	Identity      string
	UserIdentity  string
	PostsIdentity string
	ReplyIdentity string
	ReplyName     string
	Content       string
	Mention       string
	Like          int
	Dislike       int
	Read          int
	CreatedAt     time.Time      `gorm:"created"`
	UpdatedAt     time.Time      `gorm:"updated"`
	DeletedAt     gorm.DeletedAt `gorm:"deleted"`
}

func (table PostsCommentBasic) TableName() string {
	return "posts_comment_basic"
}
