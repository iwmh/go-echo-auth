package model

import (
	"github.com/jinzhu/gorm"
)

type Session struct {
	gorm.Model
	Username string
	Agent    string
	Message  string
}
