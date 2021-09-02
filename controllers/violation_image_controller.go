package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"log"
	u "mobile-rest-api/utils"
	"net/http"
)

//Data xml
type Data struct {
	XMLName xml.Name `xml:"xml" json:"-"`
	Error   string   `xml:"err" json:"error"`
	Info    *Info    `xml:"info" json:"info"`
}
//Info xml
type Info struct {
	XMLName xml.Name `xml:"info" json:"-"`
	Rows    Rows     `xml:"rows" json:"rows"`
}

//Rows xml
type Rows struct {
	XMLName xml.Name `xml:"rows" json:"-"`
	Row     []Row    `xml:"row" json:"row"`
}
//Row xml
type Row struct {
	URL string `xml:"url,attr" json:"url"`
}

//CustomerImages controller
var CustomerImages = func(w http.ResponseWriter, r *http.Request) {

	pin := "4698$p0ytAkht"
	username := "poytakht"
	info := "info"
	q := r.URL.Query()
	externalID := q.Get("BId")
	if externalID == "" {
		u.Respond(w, u.Message(720, "External ID Not found !!!"))
		return
	}
	hash := GetMD5Hash(username + externalID + pin + info)
	log.Println("hash (md5)===>", hash)

	res, err := http.Get("http://download1.safecity.tj/get.aspx?username=" + username + "&ExternalID=" + externalID + "&key=" + hash + "&action=info")

	if err != nil {
		log.Println("errror==>", err)
		u.Respond(w, u.Message(721, "Сервер изображения не отвечает !!!"))
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		u.Respond(w, u.Message(700, "Ошибка. попробуйте ещё раз !!!"))
		return
	}
	var data Data
	xml.Unmarshal([]byte(body), &data)
	jsonData, _ := json.Marshal(data)

	var result map[string]interface{}
	json.Unmarshal([]byte(string(jsonData)), &result)

	response := u.Message(0, "")
	response["data_Status"] = result["error"]
	var urls interface{}
	if result["error"] != "ERROR" && result["error"] != "FILE_NOT_FOUND" {
		urls = result["info"].(map[string]interface{})["rows"].(map[string]interface{})["row"]
	} else {
		response["status"] = false
		response["message"] = "В данный момент изображения не доступны !!!"
	}
	response["Image_Links"] = urls

	u.Respond(w, response)

	return
}

//GetMD5Hash func
func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

//GetImageViolation func
func GetImageViolation(externalID string) map[string]interface{} {
	response := make(map[string]interface{})

	pin := "4698$p0ytAkht"
	username := "poytakht"
	info := "info"
	hash := GetMD5Hash(username + externalID + pin + info)

	res, err := http.Get("http://download1.safecity.tj/get.aspx?username=" + username + "&ExternalID=" + externalID + "&key=" + hash + "&action=info")
	if err != nil {
		log.Println("errror==>", err)
		response["status"] = false
		response["message"] = "В данный момент изображения не доступны !!!"
		return response

	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		response["status"] = false
		response["message"] = "В данный момент изображения не доступны !!!"
		return response
	}
	var data Data
	xml.Unmarshal([]byte(body), &data)
	jsonData, _ := json.Marshal(data)

	var result map[string]interface{}
	json.Unmarshal([]byte(string(jsonData)), &result)

	if result["error"] == "ERROR" || result["error"] == "FILE_NOT_FOUND" {
		response["status"] = false
	} else {
		response["status"] = true
	}
	var urls interface{}
	if result["error"] != "ERROR" && result["error"] != "FILE_NOT_FOUND" {
		urls = result["info"].(map[string]interface{})["rows"].(map[string]interface{})["row"]
	} else {
		response["status"] = 722
		response["message"] = "В данный момент изображения не доступны !!!"
	}
	response["message"] = ""
	response["Image_Links"] = urls

	return response
}
