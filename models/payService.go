package models

import "mobile-rest-api/db"

//PayService model
type PayService struct {
	ID        int    `gorm:"AUTO_INCREMENT" json:"id"`
	Title     string `json:"title"`
	Login     string `gorm:"unique_index;" json:"-"`
	Password  string `gorm:"unique_index;" json:"-"`
	SortOrder int    `json:"sort_order"`
	Content   string `json:"content"`
}


//GetListPayServices func
func GetListPayServices()(pss []PayService){
	db.Init()
	db := db.GetDB()
	defer db.Close()
	db.Order("sort_order ASC", true).Find(&pss)
	return pss
}
