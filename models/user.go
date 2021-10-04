package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username  string  `gorm:"uniqueIndex" json:"username,omitempty"`
	Password  string  `json:"-"`
	Upvoted   []*Post `gorm:"many2many:user_upvoted;" json:"upvoted"`
	Downvoted []*Post `gorm:"many2many:user_downvoted;" json:"downvoted"`
}
