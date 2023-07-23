package handlers

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"github.com/nuradd1n/backend/db"
	"github.com/nuradd1n/backend/models"
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

		// Получите список пользователей, не связанных с какой-либо организацией
		users, err := db.GetUsers()
		if err != nil {
			ErrorPage(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// users, err := db.GetUsers()
		// if err != nil {
		// 	ErrorPage(w, err.Error(), http.StatusInternalServerError)
		// 	return
		// }

		data := struct {
			Users []models.User
		}{
			Users: users,
		}

		tmpl.Execute(w, data)

	case "POST":

		if err := r.ParseForm(); err != nil {
			ErrorPage(w, err.Error(), http.StatusInternalServerError)
			return
		}

		bin := r.FormValue("bin")
		name := r.FormValue("name")
		address := r.FormValue("address")
		selectedUsers := r.Form["users[]"]
		users := strings.Join(selectedUsers, ", ")

		fmt.Println(bin, name, address, users, "--")

		// for _, userFullName := range selectedUsers {
		// 	// Если пользователь уже выбран для другой организации, выдаем ошибку
		// 	if selectedUsersMap[userFullName] {
		// 		ErrorPage(w, "User is already selected for another organization", http.StatusBadRequest)
		// 		return
		// 	}
		// 	// Иначе, помечаем пользователя как выбранного
		// 	selectedUsersMap[userFullName] = true
		// }

		// // Формируем строку с именами пользователей
		// users := strings.Join(selectedUsers, ", ")

		err := db.AddOrganization(bin, name, address, users)
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

func GetOrganizationPageByID(w http.ResponseWriter, r *http.Request) {
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

	orgIDStr := parts[2]
	orgID, err := strconv.Atoi(orgIDStr)
	if err != nil {
		ErrorPage(w, "Invalid Organization ID", http.StatusBadRequest)
		return
	}

	organization, err := db.GetOrganizationByID(orgID)
	if err != nil {
		if err == sql.ErrNoRows {
			ErrorPage(w, "The organization does not exist", http.StatusNotFound)
			return
		}
		ErrorPage(w, "Database query error", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("./ui/templates/organizationByID.html")
	if err != nil {
		ErrorPage(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, organization)
}
