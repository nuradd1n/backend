package db

import (
	"database/sql"
	"log"
	"time"

	// "github.com/gorilla/mux"
	// "net/http"
	// "strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/nuradd1n/backend/models"
)

var DB *sql.DB

func InitDB() error {
	db, err := sql.Open("mysql", "root:8755@tcp(localhost:3306)/1project")
	if err != nil {
		log.Fatal(err)
	}
	DB = db

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func createTables() error {
	_, err := DB.Exec(`
		CREATE TABLE IF NOT EXISTS organizations (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			bin VARCHAR(12),
			name VARCHAR(20),
			address VARCHAR(20),
			users VARCHAR(100)
		);

		CREATE TABLE IF NOT EXISTS projects (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name VARCHAR(20),
			deadline DATE,
			owner VARCHAR(20),
			users VARCHAR(100)
		);
	`)
	if err != nil {
		log.Println("Error creating tables:", err)
		return err
	}

	return nil
}

func GetOrganizations() ([]models.Organization, error) {
	query := "SELECT id, bin, name, address, users FROM organizations"
	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}

	var organizations []models.Organization
	for rows.Next() {
		var organization models.Organization
		err := rows.Scan(&organization.ID, &organization.BIN, &organization.Name, &organization.Address, &organization.Users)
		if err != nil {
			return nil, err
		}
		organizations = append(organizations, organization)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return organizations, nil
}

func AddOrganization(bin, name, address, users string) error {
	query := "INSERT INTO organizations (bin, name, address, users) VALUES (?, ?, ?, ?)"
	_, err := DB.Exec(query, bin, name, address, users)
	if err != nil {
		return err
	}

	return nil
}

func GetProject() ([]models.Project, error) {
	query := "SELECT id, name, deadline, owner, users FROM projects"
	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}
	var projects []models.Project
	for rows.Next() {
		var project models.Project
		var deadlineStr string
		err := rows.Scan(&project.ID, &project.Name, &deadlineStr, &project.Owner, &project.Users)
		if err != nil {
			return nil, err
		}
		project.Deadline, err = time.Parse("2006-01-02", deadlineStr)
		if err != nil {
			return nil, err
		}
		projects = append(projects, project)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return projects, nil
}

func CreateProject(name string, deadline time.Time, owner, users string) error {
	query := "INSERT INTO projects (name, deadline, owner, users) VALUES (?, ?, ?, ?)"
	_, err := DB.Exec(query, name, deadline, owner, users)
	if err != nil {
		return err
	}

	return nil
}

func GetOrganizationByName(name string) (models.Organization, error) {
	query := "SELECT id, bin, name, address, users FROM organizations WHERE name = ?"
	var organization models.Organization
	err := DB.QueryRow(query, name).Scan(&organization.ID, &organization.BIN, &organization.Name, &organization.Address, &organization.Users)
	if err != nil {
		return organization, err
	}
	return organization, nil
}

func GetOrganizationByID(orgID int) (models.Organization, error) {
	query := "SELECT id, bin, name, address, users FROM organizations WHERE id = ?"
	var organization models.Organization
	err := DB.QueryRow(query, orgID).Scan(&organization.ID, &organization.BIN, &organization.Name, &organization.Address, &organization.Users)
	if err != nil {
		return organization, err
	}
	return organization, nil
}

func GetProjectByID(orgID int) (models.Project, error) {
	query := "SELECT id, name, deadline, owner, users FROM projects WHERE id = ?"
	var project models.Project
	var deadlineStr string
	err := DB.QueryRow(query, orgID).Scan(&project.ID, &project.Name, &deadlineStr, &project.Owner, &project.Users)
	if err != nil {
		return project, err
	}
	project.Deadline, err = time.Parse("2006-01-02", deadlineStr)
	if err != nil {
		return project, err
	}
	return project, nil
}

func GetUsers() ([]models.User, error) {
	query := "SELECT id, full_name, email, phone FROM users"
	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.FullName, &user.Email, &user.Phone)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func AddUser(fullName, email, phone string) error {
	query := "INSERT INTO users (full_name, email, phone) VALUES (?, ?, ?)"
	_, err := DB.Exec(query, fullName, email, phone)
	if err != nil {
		return err
	}

	return nil
}
