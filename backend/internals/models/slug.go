package models

type Slug struct {
	Model
	UrlId  *uint64 `json:"url_id"`
	UserId *uint64 `json:"user_id"`
	Slug   string  `gorm:"type:text;not null;unique" json:"slug"`
	User   User    `gorm:"foreignKey:UserId;constraint:OnDelete:SET NULL;" json:"-"`
	Url    Url     `gorm:"foreignKey:UrlId;constraint:OnDelete:SET NULL;" json:"-"`
}
