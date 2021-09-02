package models

import (
	"fmt"
	"mobile-rest-api/db"
)

//Tarrif model
type Tarrif struct {
	//gorm.Model
	ID        int     `gorm:"AUTO_INCREMENT" json:"id"`
	Title     string  `gorm:"not null" json:"title"`
	ServiceID int     `gorm:"not null" json:"service_id"`
	Price     float64 `gorm:"not null" json:"price"`
	Days      int     `gorm:"not null" json:"days"`
	Comment   string  `gorm:" not null" json:"-"`
	WithImage int     `gorm:" not null" json:"with_image"`
	SortOrder int     `gorm:" not null" json:"sort_order"`
}

//GetTarrifs func
func GetTarrifs() []*Tarrif {

	db.Init()
	db := db.GetDB()
	defer db.Close()
	var tarrifs = make([]*Tarrif, 0)
	err := db.Table("tarrifs").Find(&tarrifs).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return tarrifs
}

//GetTarrifsByService func
func GetTarrifsByService(serviceID int) (tarrifs []Tarrif) {

	db.Init()
	db := db.GetDB()
	defer db.Close()
	err := db.Table("tarrifs").Where("service_id = ?", serviceID).Order("sort_order ASC", true).Find(&tarrifs).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return tarrifs
}

//GetTarrifByID func
func GetTarrifByID(tID int, serviceID int) (tarrif Tarrif) {

	db.Init()
	db := db.GetDB()
	defer db.Close()
	err := db.Table("tarrifs").Where("id = ? AND service_id = ?", tID, serviceID).Order("sort_order ASC", true).First(&tarrif).Error
	if err != nil {
		tarrif.ID = 0
		return
	}

	return tarrif
}
