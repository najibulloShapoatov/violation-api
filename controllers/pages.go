package controllers

import (
	"mobile-rest-api/models"
	"mobile-rest-api/utils"
	"net/http"
)

//GetListPages func
var GetListPages = func(w http.ResponseWriter, r *http.Request) {
	_, err := checkAuth(r.Context().Value("user"))
	if err {
		w.WriteHeader(http.StatusUnauthorized)
		utils.Respond(w, utils.Message(701, "Ошибка. попробуйте ещё раз  !!!"))
		return
	}
	resp := utils.Message(0, "")
	resp["pages"] = models.GetListPages()
	utils.Respond(w, resp)
}

//GetPageBySlug func
var GetPageBySlug = func(w http.ResponseWriter, r *http.Request) {
	_, err := checkAuth(r.Context().Value("user"))
	if err {
		w.WriteHeader(http.StatusUnauthorized)
		utils.Respond(w, utils.Message(701, "Ошибка. попробуйте ещё раз  !!!"))
		return
	}
	q := r.URL.Query()
	s := q.Get("slug")
	resp := utils.Message(0, "")
	resp["page"] = models.GetPagesBySlug(s)
	utils.Respond(w, resp)
}
