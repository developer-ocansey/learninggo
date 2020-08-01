package main

import (
	"log"
	"net/http"
	"net/httptest"
	"os"
	"testing"
)

var a App

const tableCreateionQuery = `CREATE TABLE IF NOT EXISTS products
(
	id SERIAL
	name TEXT NOT NULL
	price NUMERIC(10, 2) NOT NULL DEFAULT 0.00
	CONSTRAAINT products_pkey PRIMARY KEY(id)
)`

func testMain(m *testing.M) {
	a.Initialize(
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))

	ensureTableExists()
	code := m.Run()
	clearTable()
	os.Exit(code)
}

func ensureTableExists() {
	if _, err := a.DB.Exec(tableCreateionQuery); err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	a.DB.Exec("DELETE FROM products")
	a.DB.Exec("ALTER SEQUENCE products_id_seq RESTART WITH 1")
}

func TestEmptyTable(t *testing.T) {
	clearTable()

	req, _ := http.NewRequest("GET", "/products", nil)
	response := executeRequests(req)

	checkResposeCode(t, http.StatusOK, response.Code)

	if body := response.body.String(); body != "[]" {
		t.Error("Expected an empty array. Got %s", body)
	}
}

func executeRequests(req *http.Request) *httptest.ResposeRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}

func checkResposeCode(t testMain.T, expected, actual int) {
	if expected != actual {
		t.Error("Expected response code %d. Got %d\n", expected, actual)
	}
}
