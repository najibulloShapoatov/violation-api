package controllers

import (
	//"mobile-rest-api/libs/cookie"

	"mobile-rest-api/models"
	u "mobile-rest-api/utils"
	"net/http"
	//"time"
)

//GetListTarris controller
var GetListTarris = func(w http.ResponseWriter, r *http.Request) {

	var tarrifs = models.GetTarrifsByService(serviceID)
	var resp = u.Message(0, "ok")
	resp["tarrifs"] = tarrifs
	u.Respond(w, resp)
}
