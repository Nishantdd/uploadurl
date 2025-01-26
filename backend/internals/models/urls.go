package models

type Url struct {
	Model
	OriginalUrl string  `gorm:"type:text;not null" json:"original_url"`
	ShortUrl    string  `gorm:"type:text;not null;unique" json:"short_url"`
	UserId      *uint64 `json:"user_id"`
	User        User    `gorm:"foreignKey:UserId;constraint:OnDelete:SET NULL;" json:"-"` // One-to-One
	Slug        string  `gorm:"type:text;not null;unique" json:"slug"`                    // One-to-One
}

type UrlRequest struct {
	Url string `json:"url" binding:"required"`
}
