package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Message(status int, message string) map[string]interface{} {
	return map[string]interface{} {"status": status, "message":message}
}

func Respond(w http.ResponseWriter,  data map[string]interface{}) {

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
	
}

func JSONResponse(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	fmt.Fprintf(w, message)
}
