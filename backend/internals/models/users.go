package models

type User struct {
	Model
	Username string  `gorm:"type:varchar(100);unique;not null" json:"username" binding:"required,max=100"`
	Email    string  `gorm:"type:varchar(50);unique;not null" json:"email" binding:"required,email,max=50"`
	Password string  `gorm:"type:varchar(100);not null" json:"-" binding:"required,min=8,max=30"`
	URLs     []Url   `gorm:"foreignKey:UserId" json:"urls,omitempty"`  // One-to-Many
	Files    []Files `gorm:"foreignKey:UserId" json:"files,omitempty"` // One-to-Many
	Slugs    []Slug  `gorm:"foreignKey:UserId" json:"slugs,omitempty"` // One-to-Many
}

type UserWithPassword struct {
	Username string `json:"username" binding:"required,max=100"`
	Email    string `json:"email" binding:"required,email,max=50"`
	Password string `json:"password" binding:"required,min=8,max=30"`
}
