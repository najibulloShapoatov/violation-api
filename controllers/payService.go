package controllers

import (
	"mobile-rest-api/models"
	"mobile-rest-api/utils"
	"net/http"
)

//GetListPayServices controller
var GetListPayServices = func(w http.ResponseWriter, r *http.Request) {
	_, err := checkAuth(r.Context().Value("user"))
	if err {
		w.WriteHeader(http.StatusUnauthorized)
		utils.Respond(w, utils.Message(701, "Ошибка. попробуйте ещё раз  !!!"))
		return
	}
	resp := utils.Message(0, "")
	resp["pay_services"] = models.GetListPayServices()
	utils.Respond(w, resp)
}
