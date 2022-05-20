package models

import (
	"time"

	"gorm.io/gorm"
)

type RepositoryPool struct {
	Id        int
	Identity  string
	Hash      string
	Name      string
	Ext       string
	Size      int64
	Path      string
	CreatedAt time.Time      `xorm:"created"`
	UpdatedAt time.Time      `xorm:"updated"`
	DeletedAt gorm.DeletedAt `xorm:"deleted"`
}

func (table RepositoryPool) TableName() string {
	return "repository_pool"
}
