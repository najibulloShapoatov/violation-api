package uploader

import (
	"io/ioutil"
	"log"
	"math/rand"
	"mime/multipart"
	"os"
	"strings"
	"time"
)

var err error




func UploadFileServer(file multipart.File, handle *multipart.FileHeader, path string) string {

	var data, err = ioutil.ReadAll(file)
	if err != nil {
		return "not readble data !!!"
	}

	var name = strings.Split(handle.Filename, ".")
	var ext = name[len(name)-1]
	var fileName = UniqidString() + "." + ext
	var Path = "./public/uploads/" + path
	_, err = os.Stat(Path)
	if os.IsNotExist(err) {
		errDir := os.MkdirAll(Path, 0755)
		if errDir != nil {
			log.Fatal(err)
		}

	}

	err = ioutil.WriteFile(Path+fileName, data, 0666)
	if err != nil {
		return "not saved from folder!!!"
	}

	return fileName
}

func UniqidString() string {
	var t1 = time.Now().Format("20060102150405")
	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		"0123456789")
	length := 8
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	var s = b.String()
	return t1 + "__" + s
}


/*
//###########################Example Controller########################################in controller
import "mobile-rest-api/uploader"
//////////////////////////////////////////////////////
var UploadFile = func(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	var file, handle, err = r.FormFile("file") //file key from request

	if err != nil {
		fmt.Fprintf(w, "%v", err.Error())
		return
	}
	defer file.Close()

	mimeType := handle.Header.Get("Content-Type")
	switch mimeType {
	case "image/jpeg", "image/jpg", "image/png":
		var fileName = uploader.UploadFileServer(file, handle, "test/6/")
		u.JSONResponse(w, http.StatusCreated, "Succesfully uploaded !!!\n"+fileName)
	default:
		u.JSONResponse(w, http.StatusBadRequest, "Not valid format file !!!")
	}
}
#####################################################################################
*/
