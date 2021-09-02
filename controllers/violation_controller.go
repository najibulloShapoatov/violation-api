package controllers

import (
	"mobile-rest-api/models"
	u "mobile-rest-api/utils"
	"net/http"
	"strconv"
	"time"
)

//CustomerGetAllViolations controller
var CustomerGetAllViolations = func(w http.ResponseWriter, r *http.Request) {

	phoneNo, err := checkAuth(r.Context().Value("user"))
	if err {
		w.WriteHeader(http.StatusUnauthorized)
		u.Respond(w, u.Message(701, "Ошибка. попробуйте ещё раз!!!"))
		return
	}

	q := r.URL.Query()
	page := 0
	pageSize := 25
	if s, err := strconv.Atoi(q.Get("page")); err == nil {
		page = s
	}
	if s, err := strconv.Atoi(q.Get("pageSize")); err == nil {
		pageSize = s
	}
	if page > 0 {
		page = page - 1
	}
	plateNo := q.Get("plateNo")
	if plateNo == "" {
		u.Respond(w, u.Message(714, "Ошибка. нет номера машины !!!"))
	}

	susb := models.GetSubscription(phoneNo, plateNo, serviceID)

	if susb.ID == 0 {
		u.Respond(w, u.Message(714, "Ошибка. подписка не обнаружено !!!"))
		return
	}

	if time.Now().After(susb.DateEnd) {
		u.Respond(w, u.Message(715, "Ошибка. срок подписки истек !!!"))
		return
	}

	paid := -1
	sts := -1
	viol := -1

	if s, err := strconv.Atoi(q.Get("paid")); err == nil {
		paid = s
	}
	if s, err := strconv.Atoi(q.Get("sts")); err == nil {
		sts = s
	}
	if s, err := strconv.Atoi(q.Get("viol")); err == nil {
		viol = s
	}

	var resp = u.Message(0, "")
	var violat = models.GetAllViolationsFilter(plateNo, page, pageSize, paid, sts, viol)
	var violations []models.Violation

	for _, i := range violat {
		imgs := GetImageViolation(i.BId)
		violations = append(violations,
			models.Violation{ID: i.ID,
				BId:           i.BId,
				VehiclePlate:  i.VehiclePlate,
				VTime:         i.VTime,
				VLocation:     i.VLocation,
				VId:           i.VId,
				VDescription:  i.VDescription,
				ProcessStatus: i.ProcessStatus,
				PunishStatus:  i.PunishStatus,
				DateCreate:    i.DateCreate,
				DateUpdate:    i.DateUpdate,
				IsPaid:        i.IsPaid,
				IsPublished:   i.IsPublished,
				CreateBy:      i.CreateBy,
				UpdateBy:      i.UpdateBy,
				FileName:      i.FileName,
				Images:        imgs,
			})
	}
	resp["violations"] = violations
	if page <= 1 {
		resp["violationsQNT"] = getViolationQnt(plateNo)
	}

	u.Respond(w, resp)
}

//getViolationQnt
func getViolationQnt(vp string) map[string]string {
	violations, count := models.GetAllViolationsForSMS(vp)
	totalQnt := count
	approvedQnt := 0
	paidQnt := 0
	processQnt := 0
	if count > 0 {
		for _, i := range violations {
			// is paid
			if i.IsPaid == 1 {
				paidQnt++
			}
			// approved
			if i.IsPaid == 0 && i.ProcessStatus == 1 {
				approvedQnt++
			}
			// process status
			if i.IsPaid == 0 && i.ProcessStatus != 1 {
				processQnt++
			}
		}
	}
	res := make(map[string]string)

	res["total_qnt"] = strconv.Itoa(totalQnt)
	res["approved_qnt"] = strconv.Itoa(approvedQnt)
	res["paid_qnt"] = strconv.Itoa(paidQnt)
	res["process_qnt"] = strconv.Itoa(processQnt)
	return res
}
