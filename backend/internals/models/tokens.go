package models

import "time"

type Token struct {
	Token     string    `gorm:"type:varchar(100);unique;not null" json:"token" binding:"required"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}
