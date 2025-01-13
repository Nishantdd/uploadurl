package models

type Files struct {
	Model
	FileName string `gorm:"type:varchar(50)" json:"filename" binding:"min=1,max=50"`
	FileHash string `gorm:"type:varchar(64)" json:"filehash"`
	UrlId    uint64 `json:"url_id"`
	UserId   uint64 `json:"user_id"`
}
