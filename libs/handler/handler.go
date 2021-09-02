package handlers

import (
	"regexp"
)

// ValidVehiclePlate function
func ValidVehiclePlate(vehiclePlate string) bool {

	// 4444АА01
	validID := regexp.MustCompile(`^[0-9]{4}[A-Z]{2}[0-9]{2}`)
	if validID.MatchString(vehiclePlate) {
		return true
	}
	//АА4444РТ01
	validID = regexp.MustCompile(`^[A-Z]{2}[0-9]{4}PT0[1-5]{1}`)
	if validID.MatchString(vehiclePlate) {
		return true
	}
	// АА444401
	validID = regexp.MustCompile(`^[A-Z]{2}[0-9]{6}`)
	if validID.MatchString(vehiclePlate) {
		return true
	}
	// 4444А01 давлати
	validID = regexp.MustCompile(`^[0-9]{4}[A-Z]{1}[0-9]{2}`)
	if validID.MatchString(vehiclePlate) {
		return true
	}
	// 444ХА01
	validID = regexp.MustCompile(`^[0-9]{3}[A-Z]{2}[0-9]{2}`)
	if validID.MatchString(vehiclePlate) {
		return true
	}

	return false

}
