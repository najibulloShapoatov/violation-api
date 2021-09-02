package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strings"

	"mobile-rest-api/apis"
	"mobile-rest-api/libs/cookie"
	"mobile-rest-api/libs/validator"
	"mobile-rest-api/models"
	u "mobile-rest-api/utils"
	"net/http"
)

const (
	serviceID = 1
)

//Err error
var Err error

//CreateCustomer controller
var CreateCustomer = func(w http.ResponseWriter, r *http.Request) {

	q := r.URL.Query()
	phoneNo := q.Get("phone")

	if phoneNo == "" || !validator.ValidatePhone("+992"+phoneNo) {
		u.Respond(w, u.Message(709, "Неправильно введен  номер.  Введите номер телефона, пример: 900112233 !!!"))
		return
	}

	resp, cust := models.CreateCustomer("+992" + phoneNo)
	if resp["status"] == 0 {
		smsText := "Код потверждения: " + cust.SmsCode
		apis.CreateSendSMS(serviceID, cust.PhoneNo, 1, smsText)
	} else {
		log.Println(serviceID, cust.PhoneNo, 1, "Код потверждения: "+cust.SmsCode, "\tnot send SMS")
	}

	u.Respond(w, resp)
}

//CustomerCheckSMS controller
var CustomerCheckSMS = func(w http.ResponseWriter, r *http.Request) {

	q := r.URL.Query()
	smsCode := q.Get("sms")
	phoneNo := q.Get("phone")
	var validsms = regexp.MustCompile(`^[0-9]{4}`)
	if smsCode == "" || phoneNo == "" || !validsms.MatchString(smsCode) || !validator.ValidatePhoneWithoutCode(phoneNo) {
		u.Respond(w, u.Message(708, "Неправильно введен  смс !!!"))
		return
	}
	phoneNo = "+992" + phoneNo
	log.Println("Query >>", smsCode, phoneNo)
	var resp = models.CheckSmsCodeCustomer(phoneNo, smsCode)
	if resp["status"] == 0 {
		tk := resp["customer"].(models.Customer).Token
		cookie.AddCookie(w, "token", tk)
	}
	u.Respond(w, resp)
	return
}

func checkAuth(ctx interface{}) (phoneNo string, err bool) {

	if !models.CheckAuthCustomer(ctx) {
		err = true
		return
	}
	var s = (fmt.Sprintf("%v", ctx))
	s = s[1 : len(s)-1]
	//ID = strings.Split(s, " ")[0]
	phoneNo = strings.Split(s, " ")[1]
	//sms_code = strings.Split(s, " ")[2]
	return
}

//Abs func
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}


//CheckAuthCustomer controller
var CheckAuthCustomer = func(w http.ResponseWriter, r *http.Request) {
	phoneNo, err := checkAuth(r.Context().Value("user"))
	if err {
		w.WriteHeader(http.StatusUnauthorized)
		u.Respond(w, u.Message(701, "Ошибка. попробуйте ещё раз!!!"))
		return
	}
	resp := u.Message(0, "ok")
	resp["phone"] = phoneNo
	u.Respond(w, resp)
}




//GetCustomer controller
var GetCustomer = func(w http.ResponseWriter, r *http.Request) {
	phoneNo, err := checkAuth(r.Context().Value("user"))
	if err {
		w.WriteHeader(http.StatusUnauthorized)
		u.Respond(w, u.Message(701, "Ошибка. попробуйте ещё раз!!!"))
		return
	}
	resp := u.Message(0, "ok")
	resp["customer"] = models.GetCustomer(phoneNo)
	u.Respond(w, resp)
}



//CustomerUpdate controller
var CustomerUpdate = func(w http.ResponseWriter, r *http.Request) {

	phoneNo, err := checkAuth(r.Context().Value("user"))
	if err {
		w.WriteHeader(http.StatusUnauthorized)
		u.Respond(w, u.Message(701, "Ошибка. попробуйте ещё раз!!!"))
		return
	}

	customer := models.GetCustomer(phoneNo)
	Err = json.NewDecoder(r.Body).Decode(&customer)
	if Err != nil {
		u.Respond(w, u.Message(700, Err.Error()))
		return
	}
	customer.PhoneNo = phoneNo
	customer.Update()
	resp := u.Message(0, "")
	resp["customer"] = customer
	u.Respond(w, resp)
	return
}
