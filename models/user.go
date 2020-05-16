package models

import "time"

type User struct {
	ID        uint       `gorm:"primary_key; AUTO_INCREMENT" json:"id" form:"id"`
	Name      string     `json:"name" form:"name"`
	UserName  string     `json:"userName" form:"userName"`
	Email     string     `json:"email" form:"email"`
	Password  string     `json:"password" form:"password"`
	CreatedAt *time.Time `json:"createdAt" form:"created_at"`
	UpdatedAt *time.Time `json:"updatedAt" form:"updated_at"`
}
