package models

import (
	"time"

	"gorm.io/gorm"
)

type UserRepository struct {
	Id                 int
	Identity           string
	UserIdentity       string
	ParentId           int64
	RepositoryIdentity string
	Ext                string
	Name               string
	CreatedAt          time.Time      `xorm:"created"`
	UpdatedAt          time.Time      `xorm:"updated"`
	DeletedAt          gorm.DeletedAt `xorm:"deleted"`
}

func (table UserRepository) TableName() string {
	return "user_repository"
}
