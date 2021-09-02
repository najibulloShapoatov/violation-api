package controllers

import (
	"mobile-rest-api/models"
	u "mobile-rest-api/utils"
	"net/http"
	"strconv"
)

//GetHistories controller
var GetHistories = func(w http.ResponseWriter, r *http.Request){
	
	phone, err := checkAuth(r.Context().Value("user"))
	if err {
		w.WriteHeader(http.StatusUnauthorized)
		u.Respond(w, u.Message(701, "Ошибка. попробуйте ещё раз  !!!"))
		return
	}
	q := r.URL.Query()
	page := 0
	pageSize := 25
	if s, err := strconv.Atoi(q.Get("page")); err == nil {
		page = Abs(s)
	}
	if s, err := strconv.Atoi(q.Get("pageSize")); err == nil {
		pageSize = Abs(s)
	}
	if page > 0 {
		page = page - 1
	}
	resp := u.Message(0,"")
	resp["events"] = models.GetAllHistoriesCustomer(phone, page, pageSize)
	u.Respond(w, resp)
}