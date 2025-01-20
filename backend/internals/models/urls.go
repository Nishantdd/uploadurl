package models

type Url struct {
	Model
	OriginalUrl string `gorm:"type:text;not null" json:"original_url"`
	ShortUrl    string `gorm:"type:text;not null;unique" json:"short_url"`
	Type        string `gorm:"type:varchar(6)" json:"type"` // short or custom
	UserId      uint64 `json:"user_id"`
	User        User   `gorm:"foreignKey:UserId" json:"-"`
	Files       []File `gorm:"foreignKey:UrlId" json:"files,omitempty"` // One-to-Many
	Slug        string `gorm:"type:text;not null;unique" json:"slug"`   // One-to-One
}
