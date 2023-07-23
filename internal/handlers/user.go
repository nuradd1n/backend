package handlers

import (
	"html/template"
	"net/http"

	"github.com/nuradd1n/backend/db"
)

func GetUsersPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/users" {
		ErrorPage(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		w.Header().Set("Allow", http.MethodGet)
		ErrorPage(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	users, err := db.GetUsers()
	if err != nil {
		ErrorPage(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("./ui/templates/base.html", "./ui/templates/GetUsers.html")
	if err != nil {
		ErrorPage(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, users)
}

func RegisterUserPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/users/register" {
		ErrorPage(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		tmpl, err := template.ParseFiles("./ui/templates/registerUser.html")
		if err != nil {
			ErrorPage(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, nil)

	case "POST":
		fullName := r.FormValue("full_name")
		email := r.FormValue("email")
		phone := r.FormValue("phone")

		err := db.AddUser(fullName, email, phone)
		if err != nil {
			ErrorPage(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/users", http.StatusSeeOther)

	default:
		w.Header().Set("Allow", http.MethodGet)
		ErrorPage(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
}
