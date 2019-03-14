package types

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Transaction struct {
	gorm.Model
	UserID          uint `gorm:"association_foreignkey:ID"`
	Amount          int
	TransactionType string
	BusinessName    string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
