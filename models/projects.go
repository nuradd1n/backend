package models

import "time"

type Project struct {
	ID       int
	Name     string
	Deadline time.Time
	Owner    string
	Users    string
}
