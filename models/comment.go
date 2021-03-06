package models

import "time"

type Comment struct {
	ID           uint       `gorm:"primary_key; AUTO_INCREMENT" json:"id" form:"id"`
	UserID       uint       `json:"userId" form:"userId"`
	RelationId   uint       `json:"relationId" form:"relationId"`
	Comment      string     `json:"comment" form:"comment"`
	RelationType string     `json:"relationType" form:"relationType"`
	CreatedAt    *time.Time `json:"createdAt" form:"created_at"`
	UpdatedAt    *time.Time `json:"updatedAt" form:"updated_at"`
}
