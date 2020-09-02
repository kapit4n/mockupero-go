package models

import (
	"github.com/jinzhu/gorm"
)

type MockupItem struct {
	gorm.Model
	Name     string  `json:"name" form:"name"`
	Type     string  `json:"type" form:"type"`
	Width    float64 `json:"width" form:"width"`
	Height   float64 `json:"height" form:"height"`
	Top      float64 `json:"top" form:"top"`
	Left     float64 `json:"left" form:"left"`
	Position string  `json:"position" form:"position"`
	Text     string  `json:"text" form:"text"`
	IdHtml   string  `json:"idHtml" form:"idHtml"`
	Src      string  `json:"src" form:"src"`
	MockupID uint    `json:"mockupId" form:"mockupId"`
	Mockup   Mockup  `gorm:"association_autoupdate:false;association_autocreate:false" json:"mockup" form:"mockup"`
}
