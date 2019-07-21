package model

import (
	"github.com/jinzhu/gorm"
)

type Application struct {
	DB *gorm.DB
}
