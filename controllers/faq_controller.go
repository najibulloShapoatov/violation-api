package controllers

import (
	"mobile-rest-api/models"
	"mobile-rest-api/utils"
	u "mobile-rest-api/utils"
	"net/http"
)

//GetListActiveFaq func
var GetListActiveFaq = func(w http.ResponseWriter, r *http.Request) {
	_, err := checkAuth(r.Context().Value("user"))
	if err {
		w.WriteHeader(http.StatusUnauthorized)
		utils.Respond(w, utils.Message(701, "Ошибка. попробуйте ещё раз  !!!"))
		return
	}
	
	resp := utils.Message(0, "")
	resp["faqs"] = models.GetListActiveFAQ()
	u.Respond(w, resp)
}
