package data

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	Id        int64 `gorm:"primarykey,autoIncrement:false" json:"id,omitempty"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}
