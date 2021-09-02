package models

import (
	"mobile-rest-api/db"
	"time"
)

//Service model
type Service struct {
	ID         int       `gorm:"AUTO_INCREMENT" json:"id"`
	Name       string    `json:"title"`
	Login      string    `json:"login"`
	Password   string    `json:"password"`
	DateCreate time.Time `json:"date_create"`
	IsActive   int       `json:"is_active"`
}

//GetServiceFirst func
func GetServiceFirst(id string) (ser Service) {
	db.Init()
	var db = db.GetDB()
	defer db.Close()
	err := db.Table("services").Where("id = ?", id).First(&ser).Error
	if err != nil {
		ser.ID = 0
		return ser
	}

	return ser
}
