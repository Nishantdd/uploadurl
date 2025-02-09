package models

type File struct {
	Model
	FileName string `gorm:"type:varchar(100);not null" json:"filename" binding:"required,min=1,max=100"`
	FileHash string `gorm:"type:varchar(64);not null;unique" json:"filehash"`
	FileType string `gorm:"type:varchar(20);not null" json:"file_type"`
	FileSize int64  `gorm:"not null" json:"file_size"`
	Location string `gorm:"not null" json:"location"`
	UserId   uint64 `gorm:"index" json:"user_id"`
	User     User   `gorm:"foreignKey:UserId;constraint:OnDelete:SET NULL;" json:"-"`
}
