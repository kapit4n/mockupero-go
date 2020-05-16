package models

import "time"

type Mockup struct {
	ID          uint       `gorm:"primary_key; AUTO_INCREMENT" json:"id" form:"id"`
	Name        string     `json:"name" form:"name"`
	State       string     `json:"state" form:"state"`
	Width       string     `json:"width" form:"width"`
	Height      string     `json:"height" form:"height"`
	Description string     `json:"description" form:"description"`
	ImgToShow   string     `json:"imgToShow" form:"imgToShow"`
	Project     Project    `json:"project" form:"project"`
	Owner       User       `json:"owner" form:"owner"`
	Feature     Feature    `json:"feature" form:"feature"`
	CreatedAt   *time.Time `json:"createdAt" form:"created_at"`
	UpdatedAt   *time.Time `json:"updatedAt" form:"updated_at"`
}
