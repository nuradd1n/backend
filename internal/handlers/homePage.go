package handlers

import (
	"html/template"
	"log"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorPage(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		w.Header().Set("Allow", http.MethodGet)
		ErrorPage(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	if err := r.ParseForm(); err != nil {
		log.Println(err.Error())
	}

	tmpl, err := template.ParseFiles("./ui/templates/base.html", "./ui/templates/home.html")
	if err != nil {
		ErrorPage(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}
