package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gorilla/mux"
)

//App type to define this Application
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

//Initialize init app
func (a *App) Initialize(user, pass, db string) {
	connectingString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, pass, db)
	var err error
	a.DB, err = sql.Open("postgres", connectingString)
	if err != nil {
		log.Fatal(err)
	}
	a.Router = mux.NewRouter()
}

//Run start app
func (a *App) Run(port string) {}
