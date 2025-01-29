package models

type UrlHits struct {
	Model
	Hits  uint64 `json:"hits" gorm:"default:0"`
	UrlId uint64 `json:"url_id"`
	Url   Url    `gorm:"foreignKey:UrlId;constraint:OnDelete:SET NULL;" json:"-"` // One-to-One
	Slug  string `gorm:"type:text;not null" json:"slug"`                          // One-to-One
}
