package controllers

import (
	"encoding/json"
	"mobile-rest-api/models"
	u "mobile-rest-api/utils"
	"net/http"
	"time"
)

//CreateInsurance controller
var CreateInsurance = func(w http.ResponseWriter, r *http.Request) {

	phoneNo, err := checkAuth(r.Context().Value("user"))
	if err {
		w.WriteHeader(http.StatusUnauthorized)
		u.Respond(w, u.Message(701, "Ошибка. попробуйте ещё раз  !!!"))
		return
	}

	cust := models.GetCustomer(phoneNo)
	var iorders models.InsuranceOrder

	er := json.NewDecoder(r.Body).Decode(&iorders)
	if er != nil {
		u.Respond(w, u.Message(400, er.Error()))
		return
	}

	ins := models.GetInsurance(iorders.InsuranceID)

	if ins.ID == 0 {
		u.Respond(w, u.Message(404, "Страховая организация не найдено  !!!"))
		return
	}
	iorders.PhoneNo = cust.PhoneNo

	tyear := time.Now().Year()

	if (tyear - iorders.Birthday.Year()) < ins.VBirthday {
		u.Respond(w, u.Message(455, "Ваш возрасть не устраиваеть требования страховой организации !!!"))
		return
	}

	if iorders.Staj < ins.VStaj {
		u.Respond(w, u.Message(456, "Ваш стаж не устраиваеть требования страховой организации !!!"))
		return
	}
	if iorders.AutoYear < ins.VAutoYear {
		u.Respond(w, u.Message(457, "Ваш авто не устраиваеть требования страховой организации !!!"))
		return
	}
	iorders.PublishedAt = time.Now()

	iorders = iorders.Insert()
	res := u.Message(0, "OK")
	res["insurance_order"] = iorders
	u.Respond(w, res)
	return
}



//GetInsurance controller
var GetInsurance = func(w http.ResponseWriter, r *http.Request) {

	phoneNo, err := checkAuth(r.Context().Value("user"))
	if err {
		w.WriteHeader(http.StatusUnauthorized)
		u.Respond(w, u.Message(701, "Ошибка. попробуйте ещё раз  !!!"))
		return
	}
	res := u.Message(0, "OK")
	res["insurances"]=models.GetInsurances()
	res["insurance_orders"]=models.GetInsuranceOrders(phoneNo)
	u.Respond(w, res)
	return

}