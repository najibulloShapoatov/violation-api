package controllers

import (
	"fmt"
	"mobile-rest-api/libs/uploader"
	u "mobile-rest-api/utils"
	"net/http"
)

// UploadFile uploads a file to the server
var UploadFile = func(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	var file, handle, err = r.FormFile("file")
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
