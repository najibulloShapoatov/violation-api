package models

import (
	"mobile-rest-api/db"
	"time"
)

//Article model
type Article struct {
	ID          int       `gorm:"AUTO_INCREMENT" json:"id"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	ImageList   string    `json:"image_list"`
	ImageDetail string    `json:"image_detail"`
	IsActive    int       `gorm:"default:'1'" json:"is_active"`
	PublishedAt time.Time `json:"published_at"`
}

//GetListArticles func
func GetListArticles()(ars []Article){
	db.Init()
	db := db.GetDB()
	defer db.Close()
	db.Order("title ASC", true).Where("is_active = ?", 1).Find(&ars)
	return ars
}

//GetArticleByID func
func GetArticleByID(id int)(ar Article){
	db.Init()
	db := db.GetDB()
	defer db.Close()
	err := db.Order("title ASC", true).Where("is_active = ? and id = ?", 1, id).First(&ar).Error
	if err != nil{
		ar.ID=0
	}
	return ar
}
