package controllers

import (
	"encoding/json"
	"mobile-rest-api/models"
	"mobile-rest-api/utils"
	"net/http"
)



//CreateFeedBack func
var CreateFeedBack = func (w http.ResponseWriter, r *http.Request)  {
	_, err := checkAuth(r.Context().Value("user"))
	if err {
		w.WriteHeader(http.StatusUnauthorized)
		utils.Respond(w, utils.Message(701, "Ошибка. попробуйте ещё раз  !!!"))
		return
	}
	var f models.Feedback

	Err = json.NewDecoder(r.Body).Decode(&f)
	if Err != nil {
		utils.Respond(w, utils.Message(700, Err.Error()))
		return
	}
	f = f.Insert()
	res := utils.Message(0, "")
	res["feedback"] = f
	utils.Respond(w, res)
		return
}
//GetFeedBack func
var GetFeedBack = func (w http.ResponseWriter, r *http.Request)  {
	phone, err := checkAuth(r.Context().Value("user"))
	if err {
		w.WriteHeader(http.StatusUnauthorized)
		utils.Respond(w, utils.Message(701, "Ошибка. попробуйте ещё раз  !!!"))
		return
	}
	
	res := utils.Message(0, "")
	res["feedbacks"] = models.GetFeedback(phone)
	utils.Respond(w, res)
		return
}