package models

import (
	"time"

	"gorm.io/gorm"
)

type UserBasic struct {
	Id        int
	Identity  string
	Name      string
	Password  string
	Email     string
	CreatedAt time.Time      `xorm:"created"`
	UpdatedAt time.Time      `xorm:"updated"`
	DeletedAt gorm.DeletedAt `xorm:"deleted"`
}

func (table UserBasic) TableName() string {
	return "user_basic"
}
