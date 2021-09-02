package controllers

import (
	"mobile-rest-api/models"
	"mobile-rest-api/utils"
	"net/http"
	"strconv"
)

//GetListBanners controller
var GetListBanners = func (w http.ResponseWriter, r *http.Request)  {
	_, err := checkAuth(r.Context().Value("user"))
	if err {
		w.WriteHeader(http.StatusUnauthorized)
		utils.Respond(w, utils.Message(701, "Ошибка. попробуйте ещё раз  !!!"))
		return
	}
	resp := utils.Message(0, "")
	resp["banners"] = models.GetListBanners()
	utils.Respond(w, resp)
}



//GetBannerByID controller
var GetBannerByID = func (w http.ResponseWriter, r *http.Request)  {
	_, err := checkAuth(r.Context().Value("user"))
	if err {
		w.WriteHeader(http.StatusUnauthorized)
		utils.Respond(w, utils.Message(701, "Ошибка. попробуйте ещё раз  !!!"))
		return
	}
	id:=0
	q := r.URL.Query()
	if s, err := strconv.Atoi(q.Get("id")); err == nil {
		id=s
	}
	banner := models.GetBannerByID(id)
	if banner.ID == 0 {
		utils.Respond(w, utils.Message(788, "Баннер с таким ID не найдено  !!!"))
		return
	}
	resp := utils.Message(0, "")
	resp["banner"] = banner
	utils.Respond(w, resp)
}

//GetBannerByType controller
var GetBannerByType = func (w http.ResponseWriter, r *http.Request)  {
	_, err := checkAuth(r.Context().Value("user"))
	if err {
		w.WriteHeader(http.StatusUnauthorized)
		utils.Respond(w, utils.Message(701, "Ошибка. попробуйте ещё раз  !!!"))
		return
	}
	q := r.URL.Query()
	t := q.Get("type")
	banner := models.GetBannerByType(t)
	if banner.ID == 0 {
		utils.Respond(w, utils.Message(789, "Баннер с таким Type не найдено  !!!"))
		return
	}
	resp := utils.Message(0, "")
	resp["banner"] = banner
	utils.Respond(w, resp)
}