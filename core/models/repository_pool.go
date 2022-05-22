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
	CreatedAt time.Time      `gorm:"created"`
	UpdatedAt time.Time      `gorm:"updated"`
	DeletedAt gorm.DeletedAt `gorm:"deleted"`
}

func (table RepositoryPool) TableName() string {
	return "repository_pool"
}
