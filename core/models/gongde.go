package models

import (
	"time"

	"gorm.io/gorm"
)

type GongdeBasic struct {
	Id        int
	Count     int
	CreatedAt time.Time      `gorm:"created"`
	UpdatedAt time.Time      `gorm:"updated"`
	DeletedAt gorm.DeletedAt `gorm:"deleted"`
}

func (table GongdeBasic) TableName() string {
	return "gongde_basic"
}
