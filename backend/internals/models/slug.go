package models

type Slug struct {
	Model
	UrlId  uint64 `gorm:"uniqueIndex;not null" json:"url_id"`
	UserId uint64 `gorm:"not null" json:"user_id"`
	Slug   string `gorm:"type:text;not null;unique" json:"slug"`
	User   User   `gorm:"foreignKey:UserId" json:"-"`
	Url    Url    `gorm:"foreignKey:UrlId" json:"-"`
}
