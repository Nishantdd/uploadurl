package models

type Url struct {
	Model
	OriginalUrl string  `gorm:"type:text;not null" json:"original_url"`
	ShortUrl    string  `gorm:"type:text;not null;unique" json:"short_url"`
	Type        string  `gorm:"type:varchar(6)" json:"type"`
	UserId      *uint64 `json:"user_id"`
	User        User    `gorm:"foreignKey:UserId;constraint:OnDelete:SET NULL;" json:"-"`              // One-to-One
	Files       []File  `gorm:"foreignKey:UrlId;constraint:OnDelete:SET NULL;" json:"files,omitempty"` // One-to-Many
	Slug        string  `gorm:"type:text;not null;unique" json:"slug"`                                 // One-to-One
}

type UrlRequest struct {
	Url string `json:"url" binding:"required"`
}
