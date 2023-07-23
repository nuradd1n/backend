package route

import (
	"net/http"

	h "github.com/nuradd1n/backend/internal/handlers"
)

func Router() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", h.HomePage)
	mux.HandleFunc("/organization", h.GetOrganizationPage)
	mux.HandleFunc("/organization/register", h.RegisterOrganizationPage)
	mux.HandleFunc("/projects", h.GetProjectsPage)
	mux.HandleFunc("/projects/create", h.CreateProjectPage)
	mux.HandleFunc("/users", h.GetUsersPage)
	mux.HandleFunc("/users/register", h.RegisterUserPage)
	// mux.HandleFunc("/projects/submit", h.SubmitProject)
	mux.HandleFunc("/organization/", h.GetOrganizationPageByID)
	mux.HandleFunc("/projects/", h.GetProjectPageByID)
	// mux.HandleFunc("/projects/id/users", h.GetProjectUserPage)
	mux.Handle("/ui/static/", http.StripPrefix("/ui/static/", http.FileServer(http.Dir("./ui/static/"))))
	return mux
}
