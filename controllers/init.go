package controllers

import (
	"gorm.io/gorm"
)

type Controller struct {
	DB	*gorm.DB
}

func New(db *gorm.DB) *Controller {
	return &Controller {
		DB: db,
	}
}