package db

import (
	"database/sql"
	"log"

	// "github.com/gorilla/mux"
	// "net/http"
	// "strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/nuradd1n/backend/models"
)

var DB *sql.DB

func InitDB() error {
	db, err := sql.Open("mysql", "root:8755@tcp(localhost:3306)/myproject")
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
			workers INT
		);

		CREATE TABLE IF NOT EXISTS projects (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name VARCHAR(20),
			deadline VARCHAR(6),
			owner VARCHAR(20),
			users INT
		);
	`)
	if err != nil {
		log.Println("Error creating tables:", err)
		return err
	}

	return nil
}

func GetOrganizations() ([]models.Organization, error) {
	query := "SELECT id, bin, name, address, workers FROM organizations"
	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}

	var organizations []models.Organization
	for rows.Next() {
		var organization models.Organization
		err := rows.Scan(&organization.ID, &organization.BIN, &organization.Name, &organization.Address, &organization.Workers)
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

func AddOrganization(bin, name, address string, workers int) error {
	query := "INSERT INTO organizations (bin, name, address, workers) VALUES (?, ?, ?, ?)"
	_, err := DB.Exec(query, bin, name, address, workers)
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
		err := rows.Scan(&project.ID, &project.Name, &project.Deadline, &project.Owner, &project.Users)
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

func CreateProject(name, deadline, owner string, users int) error {
	query := "INSERT INTO projects (name, deadline, owner, users) VALUES (?, ?, ?, ?)"
	_, err := DB.Exec(query, name, deadline, owner, users)
	if err != nil {
		return err
	}

	return nil
}
