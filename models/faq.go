package models

import "mobile-rest-api/db"

//Faq model
type Faq struct {
	ID        int    `gorm:"AUTO_INCREMENT"`
	Title     string `gorm:"not null" json:"title"`
	Content   string `gorm:"" json:"content"`
	SortOrder int    `gorm:"default:'1'; not null" json:"sort_order"`
	IsActive  int    `gorm:"default:'1'; not null" json:"is_active"`
}

//GetListActiveFAQ func
func GetListActiveFAQ() (faqs []Faq) {
	db.Init()
	db := db.GetDB()
	defer db.Close()

	db.Where("is_active = ?", 1).Order("sort_order ASC", true).Find(&faqs)
	return faqs
}
