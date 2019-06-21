package models

import "time"

type Project struct {
	Id          uint       `gorm:"primary_key; AUTO_INCREMENT" json:"id" form:"id"`
	Name        string     `json:"name" form:"name"`
	State       string     `json:"state" form:"state"`
	Type        string     `json:"type" form:"type"`
	ImgtoShow   string     `json:"imgToShow" form:"imgToShow"`
	Description string     `json:"description" form:"description"`
	Mockups     []Mockup   `json:"mockups" form:"mockups"`
	CreatedAt   *time.Time `json:"created_at" form:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at" form:"updated_at"`
}
