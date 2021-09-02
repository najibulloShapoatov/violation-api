package models

import (
	"mobile-rest-api/db"
	"time"
)

// History struct
type History struct {
	ID         int       `gorm:"AUTO_INCREMENT" json:"id"`
	Action     string    `gorm:"" json:"action"` // action text
	PhoneNo    string    `gorm:"not null" json:"phone_no"`
	Comment    string    `json:"comment"`
	Sum        float64   ` json:"sum"`
	ActionDate time.Time `json:"action_date"`
}

//CreateHistory func
func CreateHistory(action, phoneNo, comment string, s float64) {
	db.Init()
	db := db.GetDB()
	defer db.Close()
	var his History
	his.Action = action
	his.PhoneNo = phoneNo
	his.Comment = comment
	his.Sum = s
	his.ActionDate = time.Now()
	db.Create(&his)
}

//GetAllHistoriesCustomer func
func GetAllHistoriesCustomer(phoneNo string, page int, pageSize int) []History {
	db.Init()
	db := db.GetDB()
	defer db.Close()
	var data []History
	db.Offset(pageSize*page).Limit(pageSize).Order("action_date DESC", true).Where("phone_no = ? ", phoneNo).Find(&data)
	return data
}
