package models

import (
	"time"
)

type Model struct {
	ID        uint64    `gorm:"primary_key; auto_increment;index" json:"id"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
