package controllers

import (
	"mobile-rest-api/models"
	u "mobile-rest-api/utils"
	"net/http"
	"strconv"
)

//CustomerPayments controller
var CustomerPayments = func(w http.ResponseWriter, r *http.Request) {

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
	resp["payments"] = models.GetAllPaymentsCustomer(phoneNo, page, pageSize)
	u.Respond(w, resp)
	return
}
