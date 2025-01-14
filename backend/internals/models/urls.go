package models

type Url struct {
	Model
	OriginalUrl string `gorm:"type:text" json:"original_url"`
	ShortUrl    string `gorm:"type:text" json:"short_url"`
	Slug        string `gorm:"type:text" json:"slug"`
	Username    string `gorm:"type:varchar(100)" json:"username" binding:"max=100"`
	Type        string `gorm:"type:varchar(8)" json:"type"`
	UserId      uint64 `json:"user_id"`
}
