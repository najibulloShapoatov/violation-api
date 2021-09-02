package models

import (
	"mobile-rest-api/db"
	"time"
)

/*
- id (int)
- logo_list (varchar)
- logo_detail (varchar)
- name (varchar)
- product_title (text)
- content (text)
- email (varchar)
- phone_no (varchar)
- address (varchar)
- is_active (tinyInt 1-yes, 0-no)

- v_birthday (fill: 18)						// v_  - validate
- v_staj (fill: 1)
- v_auto_year (fill: 2001)

- published_at (datetime)
*/

//Insurance model
type Insurance struct {
	ID           int    `gorm:"AUTO_INCREMENT" json:"-"`
	LogoList     string `gorm:""`
	LogoDetail   string `gorm:""`
	Name         string ``
	ProductTitle string ``
	Content      string ``
	Email        string ``
	PhoneNo      string ``
	Address      string ``
	IsActive     int    `gorm:"default:'1';"`
	VBirthday    int
	VStaj        int
	VAutoYear    int
	PublishedAt  time.Time `json:"published_at"`
	Password     string
}

//Insert value to db
func (i Insurance) Insert() Insurance {
	db.Init()
	db := db.GetDB()
	defer db.Close()
	db.Create(&i)
	return i
}

//GetInsurance func
func GetInsurance(id int) (i Insurance) {
	db.Init()
	db := db.GetDB()
	defer db.Close()
	err := db.Where("id = ?", id).First(&i).Error
	if err != nil {
		i.ID = 0
	}
	return i
}

//GetInsurances func
func GetInsurances() (is []Insurance) {
	db.Init()
	db := db.GetDB()
	defer db.Close()
	db.Order("id ASC", true).Find(&is)
	return is
}

/* //GetInsurances from db by customer phone
func GetInsurances(ph string) (is []Insurance ){
	db.Init()
	db := db.GetDB()
	defer db.Close()
	db.Where("phone_no = ?", ph).Order("published_at DESC", true).Find(&is)
	return is
} */
