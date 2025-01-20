package models

type User struct {
	Model
	Username string `gorm:"type:varchar(100);unique;not null" json:"username" binding:"required,max=100"`
	Email    string `gorm:"type:varchar(50);unique;not null" json:"email" binding:"required,email,max=50"`
	Password string `gorm:"type:varchar(100);" json:"-" binding:"min=8,max=30"`
	Fullname string `gorm:"type:varchar(100);" json:"fullname,omitempty" binding:"max=100"`
	URLs     []Url  `gorm:"foreignKey:UserId" json:"urls,omitempty"`  // One-to-Many
	Files    []File `gorm:"foreignKey:UserId" json:"files,omitempty"` // One-to-Many
	Slugs    []Slug `gorm:"foreignKey:UserId" json:"slugs,omitempty"` // One-to-Many
}

type UserRequest struct {
	Username string `json:"username" binding:"required,max=100"`
	Email    string `json:"email" binding:"required,email,max=50"`
	Password string `json:"password" binding:"required,min=8,max=30"`
	Fullname string `json:"fullname" binding:"max=100"`
}

type GoogleUserInfo struct {
	Email   string `json:"email"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
}

type LoginRequest struct {
	UsernameOrEmail string `json:"username_or_email" binding:"required"`
	Password        string `json:"password" binding:"required"`
}

type SignupRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=30"`
	Fullname string `json:"fullname" binding:"max=100"`
}
