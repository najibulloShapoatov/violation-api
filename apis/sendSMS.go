package apis

import (
	"log"
	"mobile-rest-api/config"
	"mobile-rest-api/models"
	"net/http"
	"net/url"
	"time"
)

//CreateSendSMS SendSMS sendin sms via kannel
func CreateSendSMS(sID int, PhoneNo string, sts int, text string) {

	var sentsms = &models.SendSMS{}
	sentsms.PhoneNo = PhoneNo
	sentsms.ServiceID = sID
	sentsms.SmsText = text
	sentsms.CreatedAt = time.Now()
	SentSMS(sentsms)
	sentsms.Insert()

}

//CreateSendSMSWithContent SendSMS sendin sms via kannel
func CreateSendSMSWithContent(sID int, PhoneNo string, sts int, text, content string) {

	var sentsms = &models.SendSMS{}

	sentsms.PhoneNo = PhoneNo
	sentsms.ServiceID = sID
	sentsms.SmsText = text
	sentsms.CreatedAt = time.Now()
	sentsms.Content = content
	SentSMS(sentsms)
	sentsms.Insert()

}

//SentSMS func
func SentSMS(sms *models.SendSMS) {
	config.Init()
	var cfg = config.Get()
	var link = cfg.Kannel.Link + "?username=" + cfg.Kannel.Username + "&password=" + cfg.Kannel.Password + "&charset=utf-8&smsc=" + cfg.Kannel.Smsc + "&to=" + sms.PhoneNo + "&from=" + cfg.Kannel.From + "&text=" + url.QueryEscape(sms.SmsText)

	log.Println("\nsending SMS\n ")
	log.Println("\n", link)

	client := &http.Client{}

	req, err := http.NewRequest("GET", cfg.Kannel.Link, nil)
	if err != nil {
		log.Println(err)
	}
	q := url.Values{}
	q.Add("username", cfg.Kannel.Username)
	q.Add("password", cfg.Kannel.Password)
	q.Add("charset", "utf-8")
	q.Add("coding", "2")
	q.Add("from", cfg.Kannel.From)
	q.Add("smsc", cfg.Kannel.Smsc)
	q.Add("to", sms.PhoneNo)
	q.Add("text", sms.SmsText)
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp.StatusCode, resp.Status)
	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		log.Println("\n\t>>>>>>>>>>>>>>sms sent>>>>>>>>>>>>>>>\t\n\n ")
		sms.Status = 1
	} else {
		log.Println("sms don`t sent to", sms.PhoneNo)
		sms.Status = 0
	}
}
