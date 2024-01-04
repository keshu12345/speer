package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
	Notes    []Note

	// Name     string `json:"name" gorm:"index;type:varchar(100)"`
	// Email    string `json:"email" gorm:"uniqueIndex;type:varchar(100)"`
	// Password string `json:"-" gorm:"index"`
	// Notes    []Note `gorm:"foreignKey:UserID"`
}

type Note struct {
	gorm.Model
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	UserID    int
	User      User `gorm:"foreignkey:UserID"`
	// Content   string    `gorm:"index;type:text"`
	// CreatedAt time.Time `gorm:"index"`
	// UpdatedAt time.Time
	// DeletedAt gorm.DeletedAt `gorm:"index"`
	// UserID    uint
}
