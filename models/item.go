package models

import (
	"time"
	"gorm.io/gorm"
)

type Item struct {
	ItemId		uint			`gorm:"primaryKey" json:"lineItemId"`
	ItemCode	string			`gorm:"not null; type:varchar(50)" json:"itemCode"`
	Description	string			`gorm:"not null; type:varchar(255)" json:"description"`
	Quantity	uint			`gorm:"not null" json:"quantity"`
	OrderId		uint			`json:"-"`
	CreatedAt	time.Time		`json:"-"`
	UpdatedAt	time.Time		`json:"-"`
	DeletedAt	gorm.DeletedAt	`json:"-"`
}