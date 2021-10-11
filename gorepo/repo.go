package gorepo

import (
	"database/sql"
	"reflect"
)

type Repo struct {
	DB *sql.DB
}

func Start(db *sql.DB) (*Repo, error) {
	repo := Repo{DB: db}
	return &repo, nil
}

// All binds all matching terms to the Query's results
func (repo *Repo) All(iface *interface{}) error {
	return nil
}

// One calls All with a limit of 1 and gets the first sorted result
func (repo *Repo) One(iface *interface{}) error {
	return nil
}

// Get a single element by ID from the database
func (repo *Repo) Get(iface *interface{}) error {
	return nil
}

// Insert a record into the database
func (repo *Repo) Insert(iface *interface{}) error {
	return nil
}

// Update a record in the database by ID (must have a primary key)
func (repo *Repo) Update(iface *interface{}) error {
	return nil
}

// Delete a record from the database
func (repo *Repo) Delete(iface *interface{}) error {
	return nil
}

func (repo *Repo) Filter(iface interface{}) *Query {
	query := Query{
		Result: &iface,
	}
	table := reflect.TypeOf(iface).Elem().Name()
	return query.Table(table)
}
