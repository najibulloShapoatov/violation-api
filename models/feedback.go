package models

import (
	"mobile-rest-api/db"
	"time"
)

//Feedback model
type Feedback struct {
	ID          int       `gorm:"AUTO_INCREMENT"`
	Fio         string    `json:"fio"`
	Phone       string    `json:"phone"`
	Content     string    `json:"content"`
	PublishedAt time.Time `json:"published_at"`
}

//Insert func
func (f Feedback) Insert() Feedback {
	db.Init()
	db := db.GetDB()
	defer db.Close()
	f.PublishedAt = time.Now()
	db.Create(&f)
	return f
}

//GetFeedback func
func GetFeedback(ph string) (fs []Feedback) {
	db.Init()
	db := db.GetDB()
	defer db.Close()
	db.Where("phone = ?", ph).Order("published_at DESC", true).Find(&fs)
	return fs
}
