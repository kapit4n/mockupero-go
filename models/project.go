package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Project struct {
	gorm.Model
	Name        string     `json:"name" form:"name"`
	State       string     `json:"state" form:"state"`
	Type        string     `json:"type" form:"type"`
	ImgtoShow   string     `json:"imgToShow" form:"imgToShow"`
	Description string     `json:"description" form:"description"`
	Mockups     []Mockup   `gorm:"association_autoupdate:false;association_autocreate:false" json:"mockups" form:"mockups"`
	CreatedAt   *time.Time `json:"createdAt" form:"created_at"`
	UpdatedAt   *time.Time `json:"updatedAt" form:"updated_at"`
}
