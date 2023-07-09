package models

type Project struct {
	ID       int
	Name     string
	Deadline string
	Owner    string
	Users    int
}
