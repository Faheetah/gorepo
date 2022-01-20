package gorepo

import (
	"database/sql"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

var dbFile *os.File
var db *sql.DB

func TestMain(m *testing.M) {
	setup()
	defer teardown()
	os.Exit(m.Run())
}

func setup() {
	var err error

	dbFile, err = ioutil.TempFile(os.TempDir(), "gorepo-")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	db, err = sql.Open("sqlite3", dbFile.Name())
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func teardown() {
	db.Close()
	os.Remove(dbFile.Name())
}
