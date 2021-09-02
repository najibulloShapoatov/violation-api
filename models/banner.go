package models

import "mobile-rest-api/db"

//Banner model
type Banner struct {
	ID         int    `gorm:"AUTO_INCREMENT" json:"id"`
	Title      string `json:"title"`
	HTTPLink   string `json:"http_link"`
	LocalLink  string `json:"local_link"`
	BannerType string `json:"banner_type"`
	IsActive   int    `gorm:"default:'1'" json:"is_active"`
}

//GetListBanners func
func GetListBanners()(brs []Banner){
	db.Init()
	db := db.GetDB()
	defer db.Close()
	db.Order("title ASC", true).Where("is_active = ?", 1).Find(&brs)
	return brs
}


//GetBannerByID func
func GetBannerByID(id int)(br Banner){
	db.Init()
	db := db.GetDB()
	defer db.Close()
	err := db.Order("title ASC", true).Where("is_active = ? and id = ?", 1, id).First(&br)
	if err!=nil {
		br.ID=0
	}
	return br
}

//GetBannerByType func
func GetBannerByType(t string)(br Banner){
	db.Init()
	db := db.GetDB()
	defer db.Close()
	err := db.Order("title ASC", true).Where("is_active = ? and banner_type = ?", 1, t).First(&br)
	if err!=nil {
		br.ID=0
	}
	return br
}
