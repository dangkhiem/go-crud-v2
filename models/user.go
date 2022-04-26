package models

import "time"

type User struct {
	ID        uint      `gorm:"primary_key;auto_increment" json:"id"`
	UserName  string    `gorm:"size:255;not null;unique" json:"name" validate:"required"`
	Email     string    `gorm:"size:100;not null;unique" json:"email" validate:"required,email"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
