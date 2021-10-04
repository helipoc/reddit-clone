package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title     string  `json:"title"`
	Content   string  `json:"content"`
	Score     int     `gorm:"default:0" json:"score"`
	Opname    string  `json:"opname,omitempty"`
	Upvotes   []*User `gorm:"many2many:post_upvotes;" json:"upvotes"`
	Downvotes []*User `gorm:"many2many:post_downvotes;" json:"downvotes"`
}
