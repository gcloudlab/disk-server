package models

import (
	"time"

	"gorm.io/gorm"
)

type ShareBasic struct {
	Id                     int
	Identity               string
	UserIdentity           string
	UserRepositoryIdentity string
	RepositoryIdentity     string
	ExpiredTime            int
	ClickNum               int
	CreatedAt              time.Time      `xorm:"created"`
	UpdatedAt              time.Time      `xorm:"updated"`
	DeletedAt              gorm.DeletedAt `xorm:"deleted"`
}

func (table ShareBasic) TableName() string {
	return "share_basic"
}
