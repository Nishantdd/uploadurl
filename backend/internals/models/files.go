package models

type File struct {
	Model
	FileName string `gorm:"type:varchar(50);not null" json:"filename" binding:"required,min=1,max=50"`
	FileHash string `gorm:"type:varchar(64);not null;unique" json:"filehash"`
	FileType string `gorm:"type:varchar(20);not null" json:"file_type"`
	FileSize int64  `gorm:"not null" json:"file_size"`
	UrlId    uint64 `json:"url_id"`
	UserId   uint64 `gorm:"index" json:"user_id"`
	User     User   `gorm:"foreignKey:UserId;constraint:OnDelete:SET NULL;" json:"-"`
	Url      Url    `gorm:"foreignKey:UrlId;constraint:OnDelete:SET NULL;" json:"-"`
}
