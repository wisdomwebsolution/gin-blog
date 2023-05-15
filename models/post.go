package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title    string `json:"title" gorm:"required"`
	Content  string `json:"content"`
	Tags     []*Tag `json:"tags" gorm:"many2many:post_tags"`
	AuthorID int
	Author   User `json:"author"`
}
