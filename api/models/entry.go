package models

type Entry struct {
	ID     *uint  `gorm:"primaryKey,column:id" json:"id"`
	URL    string `gorm:"column:url;not null" json:"url"`
	Title  string `gorm:"column:title;not null" json:"title"`
	IsDone bool   `gorm:"column:is_done;default:false;not null" json:"is_done"`
}
