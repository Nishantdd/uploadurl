package models

type Url struct {
	Model
	OriginalUrl string `gorm:"type:varchar(MAX)" json:"original_url"`
	ShortUrl    string `gorm:"type:varchar(MAX)" json:"short_url"`
	Slug        string `gorm:"type:varchar(MAX)" json:"slug"`
	Username    string `gorm:"type:varchar(100)" json:"username" binding:"max=100"`
	Type        string `gorm:"type:varchar(8)" json:"type"`
	UserId      uint64 `json:"user_id"`
}
