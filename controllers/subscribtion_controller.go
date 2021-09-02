package controllers

import (
	"log"
	"mobile-rest-api/libs/validator"
	"mobile-rest-api/models"
	u "mobile-rest-api/utils"
	"net/http"
	"strconv"
)

//CustomerSubscribtionsAll controller
var CustomerSubscribtionsAll = func(w http.ResponseWriter, r *http.Request) {

	phoneNo, err := checkAuth(r.Context().Value("user"))
	if err {
		w.WriteHeader(http.StatusUnauthorized)
		u.Respond(w, u.Message(701, "Ошибка. попробуйте ещё раз  !!!"))
		return
	}
	q := r.URL.Query()
	page := 0
	pageSize := 15
	if s, err := strconv.Atoi(q.Get("page")); err == nil {
		page = Abs(s)
	}
	if s, err := strconv.Atoi(q.Get("pageSize")); err == nil {
		pageSize = Abs(s)
	}
	if page > 0 {
		page = page - 1
	}
	var resp = u.Message(0, "OK")
	sbs := models.GetAllSubscribtionsCustomer(phoneNo, page, pageSize)
	resp["subscribtions"] = sbs
	u.Respond(w, resp)
	return
}

//CustomerSubscribe controller
var CustomerSubscribe = func(w http.ResponseWriter, r *http.Request) {

	//today := time.Now()
	phoneNo, err := checkAuth(r.Context().Value("user"))
	if err {
		w.WriteHeader(http.StatusUnauthorized)
		u.Respond(w, u.Message(701, "Ошибка. попробуйте ещё раз !!!"))
		return
	}
	q := r.URL.Query()
	plateNo := q.Get("plateNo")
	var tID int
	if plateNo == "" || !validator.ValidVehiclePlate(plateNo) {
		u.Respond(w, u.Message(706, "Ошибка. Номер автомобиля (например: 1234AB01) !!!"))
		return
	}
	if models.CheckVehicleBlackList(plateNo) {
		u.Respond(w, u.Message(718, "Ошибка. Доступ закрыть !!!"))
		return
	}
	if s, err := strconv.Atoi(q.Get("tID")); err == nil {
		tID = Abs(s)
	} else {
		u.Respond(w, u.Message(700, "Ошибка. Выберите тариф !!!"))
		return
	}

	customer := models.GetCustomer(phoneNo)
	if customer.PhoneNo == "NO DATA FOUND" {
		w.WriteHeader(http.StatusUnauthorized)
		u.Respond(w, u.Message(701, "Ошибка. попробуйте ещё раз  !!!"))
		return
	}
	tarif := models.GetTarrifByID(tID, serviceID)
	if tarif.ID == 0 {
		u.Respond(w, u.Message(715, "Ошибка. Выберите тариф  !!!"))
		return
	}

	/* err = models.CheckSubscribe(phoneNo, plateNo)
	if err {
		u.Respond(w, u.Message(705, "Ошибка. Вы уже подписанны на этот номер !!!"))
		return
	} */

	var s = customer.Balance - tarif.Price
	if (s) < 0.00 {
		u.Respond(w, u.Message(704, "Ошибка. У вас не достаточно средств на балансе !!!"))
		return
	}

	customer.Balance = s
	customer.Update()

	res, err := models.CustomerSubscribe(phoneNo, plateNo, tID)

	models.CreatePaymentSystem(phoneNo, 0, tarif.Price)

	action := "Подписка на тариф:\t" + tarif.Title
	comment := ""
	models.CreateHistory(action, phoneNo, comment, tarif.Price)
	log.Println(action, phoneNo, comment, tarif.Price)

	var resp = u.Message(0, "OK")
	resp["subscribtion"] = res
	u.Respond(w, resp)
	return
}

//SubscribeChangeAutoprolonged func
var SubscribeChangeAutoprolonged = func(w http.ResponseWriter, r *http.Request) {
	////Do source
	phoneNo, err := checkAuth(r.Context().Value("user"))
	if err {
		w.WriteHeader(http.StatusUnauthorized)
		u.Respond(w, u.Message(701, "Ошибка. попробуйте ещё раз!!!"))
		return
	}
	q := r.URL.Query()
	plateNo := q.Get("plateNo")

	subs := models.ChangeAutoProlonged(phoneNo, plateNo)
	if subs.ID == 0 {
		u.Respond(w, u.Message(744, "Подписка не найдено !!!"))
	}
	res := u.Message(0, "ok")
	res["subscription"] = subs
	u.Respond(w, res)

}
