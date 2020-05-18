package models

import "time"

type MockupItem struct {
	ID        uint       `gorm:"primary_key; AUTO_INCREMENT" json:"id" form:"id"`
	Name      string     `json:"name" form:"name"`
	Type      string     `json:"type" form:"type"`
	Width     float64    `json:"width" form:"width"`
	Height    float64    `json:"height" form:"height"`
	Top       float64    `json:"top" form:"top"`
	Left      float64    `json:"left" form:"left"`
	Position  string     `json:"position" form:"position"`
	IdHtml    string     `json:"idHtml" form:"idHtml"`
	Src       string     `json:"src" form:"src"`
	CreatedAt *time.Time `json:"createdAt" form:"created_at"`
	UpdatedAt *time.Time `json:"updatedAt" form:"updated_at"`
}
