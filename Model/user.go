package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username     string `gorm:"not null"`
	Password     string `gorm:"not null"`
	PasswordSalt string `gorm:"not null"`
	Email        string `gorm:"type:varchar(100);unique_index;not null"`
}
