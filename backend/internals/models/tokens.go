package models

import "time"

type Token struct {
	Token     string    `gorm:"type:varchar(100);unique;not null" json:"token" binding:"required"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UserId    *uint64   `json:"user_id"`
	User      User      `gorm:"foreignKey:UserId;constraint:OnDelete:SET NULL;" json:"-"` // One-to-One
}
