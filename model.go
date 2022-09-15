package main

import (
	"time"

	"gorm.io/gorm"
)

type Menu struct {
	ParentId  int64          `json:"parentId" gorm:"column:parentId;not null;"`
	Title     string         `json:"title" gorm:"column:title;not null;"`
	Name      string         `json:"name" gorm:"column:name;not null;"`
	Sort      int64          `json:"sort" gorm:"column:sort;not null;"`
	Route     string         `json:"route" gorm:"column:route;not null;"`
	Component string         `json:"component" gorm:"column:component;not null;default:'';"`
	Icon      string         `json:"icon" gorm:"column:icon;not null;default:'';"`
	CreatedAt time.Time      `json:"createdAt" gorm:"column:createdAt;"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"column:updatedAt;"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index;comment:删除时间"`
}

func Insert(model *Menu) error {
	return DB.Model(model).Create(model).Error
}
