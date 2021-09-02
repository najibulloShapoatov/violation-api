package cookie

import (
	"net/http"
	"time"
)

func AddCookie(w http.ResponseWriter, name, value string) http.ResponseWriter {
	var ttl = (64 * 8192) * time.Hour

	expire := time.Now().Add(ttl)
	cookie := http.Cookie{
		Name:    name,
		Value:   value,
		Expires: expire,
	}
	http.SetCookie(w, &cookie)
	return w
}

func AddCookieWithTime(w http.ResponseWriter, name, value string, ttl time.Duration) http.ResponseWriter {

	//var ttl = (256 * 8192) * time.Hour
	expire := time.Now().Add(ttl)
	cookie := http.Cookie{
		Name:    name,
		Value:   value,
		Expires: expire,
	}
	http.SetCookie(w, &cookie)
	return w
}
