package models

type Url struct {
	Model
	OriginalUrl string `gorm:"type:text;not null" json:"original_url"`
	ShortUrl    string `gorm:"type:text;not null;unique" json:"short_url"`
	Type        string `gorm:"type:varchar(8);default:'url'" json:"type"`
	UserId      uint64 `gorm:"not null" json:"user_id"`
	User        User   `gorm:"foreignKey:UserId" json:"-"`
	Files       []File `gorm:"foreignKey:UrlId" json:"files,omitempty"` // One-to-Many
	Slug        *Slug  `gorm:"foreignKey:UrlId" json:"slug,omitempty"`  // One-to-One
}
