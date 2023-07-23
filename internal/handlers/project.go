package handlers

import (
	"database/sql"

	// "fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/nuradd1n/backend/db"
	"github.com/nuradd1n/backend/models"
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
		// Заполняем шаблон и передаем список организаций
		tmpl, err := template.ParseFiles("./ui/templates/CreateProjects.html")
		if err != nil {
			ErrorPage(w, err.Error(), http.StatusInternalServerError)
			return
		}

		organizations, err := db.GetOrganizations()
		if err != nil {
			ErrorPage(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := struct {
			Organizations []models.Organization
		}{
			Organizations: organizations,
		}

		tmpl.Execute(w, data)

	case "POST":

		name := r.FormValue("name")
		deadlineStr := r.FormValue("deadline")
		owner := r.FormValue("owner")

		// Получаем организацию по имени (Name)
		organization, err := db.GetOrganizationByName(owner)
		if err != nil {
			ErrorPage(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Заполняем поле Users данными из организации
		users := organization.Users

		// fmt.Println(name, deadline, owner, users, "--")
		// fmt.Println(name, deadline, owner, r.FormValue("number"), "--")
		// users, err := strconv.Atoi(r.FormValue("number"))
		// if err != nil {
		// 	fmt.Println("error atoi")
		// 	ErrorPage(w, err.Error(), http.StatusInternalServerError)
		// 	return
		// }
		deadline, err := time.Parse("2006-01-02", deadlineStr)
		if err != nil {
			ErrorPage(w, "Error parsing deadline", http.StatusInternalServerError)
			return
		}

		err = db.CreateProject(name, deadline, owner, users)
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

func GetProjectPageByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		ErrorPage(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 3 {
		ErrorPage(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	projectIDStr := parts[2]
	projectID, err := strconv.Atoi(projectIDStr)
	if err != nil {
		ErrorPage(w, "Invalid Project ID", http.StatusBadRequest)
		return
	}

	project, err := db.GetProjectByID(projectID)
	if err != nil {
		if err == sql.ErrNoRows {
			ErrorPage(w, "The project does not exist", http.StatusNotFound)
			return
		}
		ErrorPage(w, "Database query error", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("./ui/templates/projectByID.html")
	if err != nil {
		ErrorPage(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, project)
}
