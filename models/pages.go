package models

import "mobile-rest-api/db"

//Pages model
type Pages struct {
	ID      int    `gorm:"AUTO_INCREMENT" json:"-"`
	Slug    string `json:"slug"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

//GetListPages func
func GetListPages()(pgs []Pages){
	db.Init()
	db := db.GetDB()
	defer db.Close()
	db.Order("slug ASC", true).Find(&pgs)
	return pgs
}

//GetPagesBySlug func
func GetPagesBySlug(s string)(pg Pages){
	db.Init()
	db := db.GetDB()
	defer db.Close()
	err := db.Order("slug ASC", true).Where("slug = ?", s).First(&pg).Error
	if err !=nil {
		pg.ID=0
	}
	return pg
}
