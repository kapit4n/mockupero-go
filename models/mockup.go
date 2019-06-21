package models

import "time"

type Mockup struct {
	ID          uint       `gorm:"primary_key; AUTO_INCREMENT" json:"id" form:"id"`
	Name        string     `json:"name" form:"name"`
	State       string     `json:"state" form:"state"`
	Description string     `json:"description" form:"description"`
	ImgToShow   string     `json:"imgToShow" form:"imgToShow"`
	Project     Project    `json:"project" form:"project"`
	Owner       User       `json:"owner" form:"owner"`
	Feature     Feature    `json:"feature" form:"feature"`
	CreatedAt   *time.Time `json:"created_at" form:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at" form:"updated_at"`
}
