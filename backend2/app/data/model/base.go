package model

import "gorm.io/gorm"

type Base struct {
	gorm.Model
}

func (base *Base) BeforeCreate(tx *gorm.DB) error {
	base.ID = 0
	return nil
}
