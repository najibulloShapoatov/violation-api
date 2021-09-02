package models

import (
	"mobile-rest-api/db"
	"time"
)

//PaymentSystem model
type PaymentSystem struct {
	//gorm.Model
	ID          int       `gorm:"AUTO_INCREMENT" json:"id"`
	Action      int       `gorm:"" json:"action"` // '+' add balance '-' subtract balance
	PhoneNo     string    `gorm:" not null" json:"phone_no"`
	PaymentType string    `gorm:"" json:"payment_type"`
	Sum         float64   `gorm:"default:'0.00'" json:"sum"`
	OperDate    time.Time `json:"oper_date"`
}

//GetAllPaymentsCustomer func
func GetAllPaymentsCustomer(phoneNo string, page int, pageSize int) []PaymentSystem {
	db.Init()
	db := db.GetDB()
	defer db.Close()
	var subs []PaymentSystem
	db.Offset(pageSize*page).Limit(pageSize).Order("oper_date DESC", true).Where("phone_no = ? ", phoneNo).Find(&subs)
	return subs
}

//CreatePaymentSystem func
func CreatePaymentSystem(phoneNo string, action int, s float64) {
	db.Init()
	db := db.GetDB()
	defer db.Close()
	var payS PaymentSystem
	payS.Action = action
	payS.PhoneNo = phoneNo
	payS.Sum = s
	payS.OperDate = time.Now()
	db.Create(&payS)
}
