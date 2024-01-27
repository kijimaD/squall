package models

import "gorm.io/gorm"

type Entry struct {
	gorm.Model
	ID     *uint  `gorm:"primaryKey,column:id" json:"id"`
	URL    string `gorm:"column:url;not null;unique" json:"url"`
	IsDone bool   `gorm:"column:is_done;default:false;not null" json:"is_done"`
}
