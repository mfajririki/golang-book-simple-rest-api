package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title  string `form:"title"`
	Author string `form:"author"`
	Desc   string `form:"desc"`
}
