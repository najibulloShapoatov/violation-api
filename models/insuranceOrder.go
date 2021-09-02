package models

import (
	"mobile-rest-api/db"
	"time"
)

/*
- id (int)
- insurance_id (int)
- fio (varchar)
- phone_no (varchar)
- birthday (date)
- staj (varchar)
- auto_year (varchar)
- status - int (0-novaya zayavka, 1-zayavka obrabotana, 2-zayavka otkazana)
- published_at (datetime)
*/

//InsuranceOrder model
type InsuranceOrder struct {
	ID          int       `gorm:"AUTO_INCREMENT"`
	InsuranceID int       `gorm:"not null" json:"insurance_id"`
	Fio         string    `gorm:"not null" json:"fio"`
	PhoneNo     string    `json:"phone"`
	Birthday    time.Time `json:"birthday"`
	Staj        int       `gorm:"not null" json:"staj"`
	AutoYear    int       `gorm:"not null" json:"auto_year"`
	Status      int       `gorm:"not null; default:'0'" json:"status"`
	PublishedAt time.Time `json:"published_at"`
}

//Insert value to db
func (i InsuranceOrder) Insert() InsuranceOrder {
	db.Init()
	db := db.GetDB()
	defer db.Close()
	db.Create(&i)
	return i
}

//GetInsuranceOrders from db by customer phone
func GetInsuranceOrders(ph string) (is []InsuranceOrder) {
	db.Init()
	db := db.GetDB()
	defer db.Close()
	db.Where("phone_no = ?", ph).Order("published_at DESC", true).Find(&is)
	return is
}
