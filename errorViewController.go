package main

import (
	"net/http"
	// "html/template"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, status int) {

    type PageVariables struct {
        Title string
        Year string
        Content string
    }

    w.WriteHeader(status)
    if status == http.StatusNotFound {
        p := PageVariables{Title: "404 Not Found", Year: GetCurrentTime("2006")}

        err := templates.ExecuteTemplate(w, "404.html", p)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
    }
    if status == http.StatusUnauthorized {
        redirect(w, r, "/login")
    }
}