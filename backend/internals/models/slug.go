package models

type Slug struct {
	Model
	UrlId  uint64 `json:"url_id"`
	UserId uint64 `json:"user_id"`
	Slug   string `gorm:"type:text" json:"slug"`
}
