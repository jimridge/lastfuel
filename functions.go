package main

import (
	"time"
	"net/http"
)

func GetCurrentTime(timeFormat string) (string){
	time := time.Now().Local()
	return time.Format(timeFormat)
}

func redirect(w http.ResponseWriter, r *http.Request, url string) {

	var currentPath string = r.URL.Path
    http.Redirect(w, r, url + "?redirect=" + currentPath, 303)
}

