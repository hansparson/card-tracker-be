package dbcontroller

import (
	"gorm.io/gorm"
)

type HandlersController struct {
	db *gorm.DB
}

func Controller(db *gorm.DB) *HandlersController {
	return &HandlersController{
		db: db,
	}
}
