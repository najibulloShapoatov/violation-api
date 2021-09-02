package models

import (
	"mobile-rest-api/db"
	"time"
)

//Event model
type Event struct {
	ID          int       `gorm:"AUTO_INCREMENT" json:"id"`
	DatePublish time.Time `json:"date_publish"`
	Title       string    `json:"title"`
	Text        string    `json:"text"`
	Image       string    `json:"image"`
}

//GetEvents func
func GetEvents(page, pageSize int) (events []Event) {
	db.Init()
	var db = db.GetDB()
	defer db.Close()
	db.Offset(pageSize*page).Limit(pageSize).Order("date_publish DESC", true).Find(&events)
	return
}
