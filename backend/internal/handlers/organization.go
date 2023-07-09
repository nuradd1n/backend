package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/nuradd1n/backend/db"
)

func GetOrganizationPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/organization" {
		ErrorPage(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		w.Header().Set("Allow", http.MethodGet)
		ErrorPage(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	tmpl, err := template.ParseFiles("./ui/templates/base.html", "./ui/templates/organizations.html")
	if err != nil {
		ErrorPage(w, err.Error(), http.StatusInternalServerError)
		return
	}
	organization, err := db.GetOrganizations()
	if err != nil {
		ErrorPage(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, organization)
}

func RegisterOrganizationPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/organization/register" {
		ErrorPage(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	switch r.Method {
	case "GET":
		tmpl, err := template.ParseFiles("./ui/templates/registerOrganization.html")
		if err != nil {
			ErrorPage(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)

	case "POST":
		err := r.ParseForm()
		if err != nil {
			ErrorPage(w, err.Error(), http.StatusInternalServerError)
			return
		}

		bin := r.FormValue("bin")
		name := r.FormValue("name")
		address := r.FormValue("address")
		fmt.Println(bin, name, address, r.FormValue("number"), "--")
		workers, err := strconv.Atoi(r.FormValue("number"))
		if err != nil {
			fmt.Println("error atoi")
			ErrorPage(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = db.AddOrganization(bin, name, address, workers)
		if err != nil {
			ErrorPage(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/organization", http.StatusSeeOther)

	default:
		w.Header().Set("Allow", http.MethodGet)
		ErrorPage(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
}
