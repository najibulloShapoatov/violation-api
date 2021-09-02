package models

import (
	"mobile-rest-api/db"
	"time"
)

//Subscription model
type Subscription struct {
	//gorm.Model
	ID              int       `gorm:"AUTO_INCREMENT" json:"id"`
	PhoneNo         string    `gorm:" not null" json:"phone_no"`
	Content         string    `gorm:"" json:"content"`
	VehiclePlate    string    `gorm:"size:15" json:"vehicle_plate"`
	TarrifID        int       `gorm:"not null" json:"tarrif_id"`
	DateStart       time.Time `json:"date_start"`
	DateEnd         time.Time `json:"date_end"`
	Status          int       `json:"status"`            // 0-non-active, 1-active, 2-waiting
	ISAutoProlonged int       `json:"is_auto_prolonged"` // 0-non-active, 1-active,
	ServiceID       int       `json:"service_id"`
	Total           int       `json:"total"`
	Approved        int       `json:"approved"`
	Paid            int       `json:"paid"`
}

//GetAllSubscribtionsCustomer func
func GetAllSubscribtionsCustomer(phoneNo string, page int, pageSize int) []Subscription {
	db.Init()
	db := db.GetDB()
	defer db.Close()
	var subs []Subscription
	db.Where("phone_no = ? AND service_id = ? ", phoneNo, serviceID).Order("date_start DESC", true).Offset(pageSize * page).Limit(pageSize).Find(&subs)
	return subs
}

//CustomerSubscribe func
func CustomerSubscribe(phoneNo, plateNo string, tID int) (subs Subscription, err bool) {
	db.Init()
	db := db.GetDB()
	defer db.Close()
	tarrif := GetTarrifByID(tID, serviceID)
	subs, err = CheckSubscribe(phoneNo, plateNo)

	if err && subs.Status == 1 {
		subs.DateEnd = subs.DateEnd.AddDate(0, 0, tarrif.Days)
		db.Save(&subs)
		return
	}
	if err && subs.Status == 0 {
		subs.DateEnd = time.Now().AddDate(0, 0, tarrif.Days)
		db.Save(&subs)
		return
	}
	subs.PhoneNo = phoneNo
	subs.VehiclePlate = plateNo
	subs.TarrifID = tID
	subs.DateStart = time.Now()
	subs.ServiceID = serviceID
	subs.DateEnd = time.Now().AddDate(0, 0, tarrif.Days)
	subs.Status = 1
	db.Create(&subs)
	return
}

//CheckSubscribe func
func CheckSubscribe(phoneNo, plateNo string) (Subscription, bool) {
	db.Init()
	db := db.GetDB()
	defer db.Close()
	var subs Subscription
	err := db.Where("phone_no = ? AND vehicle_plate = ? AND service_id = ?", phoneNo, plateNo, serviceID).First(&subs).RecordNotFound()
	if err {
		return subs, false
	}
	return subs, true
}

//GetSubscription func
func GetSubscription(phone, plateNo string, serviceID int) (subs Subscription) {
	db.Init()
	db := db.GetDB()
	defer db.Close()
	err := db.Where("phone_no = ? AND vehicle_plate = ? AND service_id = ?", phone, plateNo, serviceID).First(&subs).RecordNotFound()
	if err {
		subs.ID = 0
		return
	}
	return subs
}

//ChangeAutoProlonged func
func ChangeAutoProlonged(phone, plateNo string) (s Subscription) {
	s = GetSubscription(phone, plateNo, serviceID)
	if s.ID == 0 {
		return s
	}
	if s.ISAutoProlonged == 1 {
		s.ISAutoProlonged = 0
	} else {
		s.ISAutoProlonged = 1
	}
	s.Update()
	return s
}

//Update data
func (s Subscription) Update() Subscription {
	db.Init()
	db := db.GetDB()
	defer db.Close()
	db.Save(&s)
	return s
}
