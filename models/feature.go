package models

import "time"

type Feature struct {
	ID        uint       `gorm:"primary_key; AUTO_INCREMENT" json:"id" form:"id"`
	Name      string     `json:"name" form:"name"`
	CreatedAt *time.Time `json:"created_at" form:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" form:"updated_at"`
}
