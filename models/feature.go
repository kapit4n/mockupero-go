package models

import "time"

type Feature struct {
	ID        uint       `gorm:"primary_key; AUTO_INCREMENT" json:"id" form:"id"`
	Name      string     `json:"name" form:"name"`
	CreatedAt *time.Time `json:"createdAt" form:"created_at"`
	UpdatedAt *time.Time `json:"updatedAt" form:"updated_at"`
}
