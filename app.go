package main

import (
	"database/sql"

	"github.com/gorilla/mux"
)

//App type to define this Application
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

//Initialize init app
func (a *App) Initialize(user, pass, db string) {}

//Run start app
func (a *App) Run(port string) {}
