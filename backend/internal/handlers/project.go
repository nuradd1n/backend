package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/nuradd1n/backend/db"
)

func GetProjectsPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/projects" {
		ErrorPage(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		w.Header().Set("Allow", http.MethodGet)
		ErrorPage(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	tmpl, err := template.ParseFiles("./ui/templates/base.html", "./ui/templates/projects.html")
	if err != nil {
		ErrorPage(w, err.Error(), http.StatusInternalServerError)
		return
	}

	projects, err := db.GetProject()
	if err != nil {
		ErrorPage(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, projects)
}

func CreateProjectPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/projects/create" {
		ErrorPage(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	switch r.Method {
	case "GET":
		tmpl, err := template.ParseFiles("./ui/templates/CreateProjects.html")
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

		name := r.FormValue("name")
		deadline := r.FormValue("deadline")
		owner := r.FormValue("owner")
		fmt.Println(name, deadline, owner, r.FormValue("number"), "--")
		users, err := strconv.Atoi(r.FormValue("number"))
		if err != nil {
			fmt.Println("error atoi")
			ErrorPage(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = db.AddOrganization(name, deadline, owner, users)
		if err != nil {
			ErrorPage(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/projects", http.StatusSeeOther)

	default:
		w.Header().Set("Allow", http.MethodGet)
		ErrorPage(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
}
