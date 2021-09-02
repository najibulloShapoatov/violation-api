package models

import (
	"mobile-rest-api/db"
	"time"
)

//BlackList modal
type BlackList struct {
	ID           int `gorm:"AUTO_INCREMENT"`
	VehiclePlate string
	Note         string
	CreatedAt    time.Time
}

//CheckVehicleBlackList func
func CheckVehicleBlackList(VehiclePlate string) bool {
	db.Init()
	var dB = db.GetDB()
	defer db.CloseDB()
	bL := BlackList{}
	err := dB.Table("black_list").Where(" vehicle_plate = ? ", VehiclePlate).First(&bL).RecordNotFound()
	if err {
		return false
	}
	return true
}
