package models

import (
	"mobile-rest-api/db"
	"time"
)

//Violation model
type Violation struct {
	//gorm.Model
	ID            int        `gorm:"AUTO_INCREMENT"`
	BId           string     `gorm:"size:30;unique_index; not null" xml:"ID"`
	VehiclePlate  string     `gorm:"size:15" xml:"VEHICLE_PLATE"`
	VTime         *time.Time `xml:"TIME"`
	VLocation     string     `gorm:"size:255" xml:"VIOLATION_LOCATION"`
	VId           string     `gorm:"size:15" xml:"VIOLATION_ID"`
	VDescription  string     `gorm:"size:255" xml:"VIOLATION"`
	ProcessStatus int        `gorm:"default:1" xml:"PROCESS_STATUS"`
	PunishStatus  int        `gorm:"default:1" xml:"PUNISH_STATUS"`
	DateCreate    *time.Time
	DateUpdate    *time.Time
	IsPaid        int
	IsPublished   int
	CreateBy      int
	UpdateBy      int
	FileName      string `gorm:"size:255"`
	Images        map[string]interface{}
}

//GetAllViolations func
func GetAllViolations(plateNo string, page int, pageSize int) []Violation {
	db.InitSafeCityDB()
	db := db.GetSafeCityDB()
	defer db.Close()
	var violations []Violation
	db.Offset(pageSize*page).Limit(pageSize).Order("v_time DESC", true).Where("vehicle_plate = ? ", plateNo).Find(&violations)

	return violations

}

//SetImages func
func (v Violation) SetImages(mapI map[string]interface{}) {
	v.Images = mapI
}

//GetAllViolationsFilter from databse
func GetAllViolationsFilter(plateNo string, page int, pageSize int, paid, sts, viol int) (violations []Violation) {
	db.InitSafeCityDB()
	db := db.GetSafeCityDB()
	defer db.Close()
	//DB.LogMode(true)

	//tx := DB.Debug().Where("vehicle_plate = ? ", plateNo)
	tx := db.Where("vehicle_plate = ? ", plateNo)

	if paid != -1 {
		tx = tx.Where("is_paid = ?", paid)
	}

	if sts == 1 {
		tx = tx.Where("process_status = ?", sts)
	}
	if sts != 1 && sts != -1 {
		tx = tx.Where("process_status != ?", 1)
	}

	if viol != -1 {
		if viol == 1625 {
			var vIDS []int
			vIDS = append(vIDS, viol)
			vIDS = append(vIDS, 1302)
			tx = tx.Where("v_id IN (?)", vIDS)
		} else {
			tx = tx.Where("v_id = ?", viol)
		}

	}

	tx.Offset(pageSize*page).Limit(pageSize).Order("v_time DESC", true).Find(&violations)

	//DB.SetLogger(log.New(os.Stdout, "\r\n", 0))

	return violations

}

func getQNTViolations(Vp string) (total, approved, paid int) {
	db.InitSafeCityDB()
	db := db.GetSafeCityDB()
	defer db.Close()

	var vls []Violation
	db.Where("vehicle_plate = ?", Vp).Find(&vls).Count(&total)

	for _, i := range vls {

		if i.IsPaid == 0 && i.ProcessStatus == 1 {
			approved++
		}

		if i.IsPaid == 1 {
			paid++
		}
	}
	return total, approved, paid

}



//GetAllViolationsForSMS return
func GetAllViolationsForSMS(plateNo string) (vs []Violation, count int) {
	db.InitSafeCityDB()
	DB := db.GetSafeCityDB()
	defer db.CloseSafeCityDB()
	DB.Where("vehicle_plate = ? ", plateNo).Find(&vs).Order("process_status").Count(&count)
	return vs, count
}
