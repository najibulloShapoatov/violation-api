package controllers

import (
	"mobile-rest-api/models"
	"mobile-rest-api/utils"
	"net/http"
	"strconv"
)

//GetListArticles contoller
var GetListArticles = func (w http.ResponseWriter, r *http.Request)  {
	_, err := checkAuth(r.Context().Value("user"))
	if err {
		w.WriteHeader(http.StatusUnauthorized)
		utils.Respond(w, utils.Message(701, "Ошибка. попробуйте ещё раз  !!!"))
		return
	}
	resp := utils.Message(0, "")
	resp["articles"] = models.GetListArticles()
	utils.Respond(w, resp)
}




//GetArticleByID contoller
var GetArticleByID = func (w http.ResponseWriter, r *http.Request)  {
	_, err := checkAuth(r.Context().Value("user"))
	if err {
		w.WriteHeader(http.StatusUnauthorized)
		utils.Respond(w, utils.Message(701, "Ошибка. попробуйте ещё раз  !!!"))
		return
	}
	id :=0
	q := r.URL.Query()
	if s, err := strconv.Atoi(q.Get("id")); err == nil{
		id=s
	}

	article := models.GetArticleByID(id)
	if article.ID == 0{
		utils.Respond(w, utils.Message(790, "Article с таким ID не найдено  !!!"))
		return
	}
	resp := utils.Message(0, "")
	resp["articles"] = article
	utils.Respond(w, resp)
}