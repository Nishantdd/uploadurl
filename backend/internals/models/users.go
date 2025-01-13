package models

type User struct {
	Model
	Username string `gorm:"type:varchar(100)" json:"username" binding:"max=100"`
	Email    string `gorm:"type:varchar(50)" json:"email" binding:"max=50"`
	Password string `gorm:"type:varchar(30)" json:"password" binding:"min=8,max=30"`
}
