package main

import (
	"github.com/asdine/storm"
)

// Application struct which houses a database
type Application struct {
	Database *storm.DB
}

// NewApplication Creates a new Application
func NewApplication() *Application {
	db := createDatabase()
	return &Application{Database: db}
}
