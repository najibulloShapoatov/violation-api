package models

import (
	"mobile-rest-api/db"
	"time"
)

//SendSMS database
type SendSMS struct {
	ID        int    `gorm:"AUTO_INCREMENT"`
	ServiceID int    `gorm:"not null"`
	PhoneNo   string `gorm:"not null"`
	Status    int    `gorm:"not null"`
	SmsText   string `gorm:"not null"`
	CreatedAt time.Time
	Content string
}

//Insert to DB
func (sendsms SendSMS) Insert() SendSMS {
	db.Init()
	db := db.GetDB()
	defer db.Close()
	db.Table("send_sms").Create(&sendsms)
	return sendsms
}
