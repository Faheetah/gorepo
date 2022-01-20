package gorepo

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"
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
func (repo *Repo) Insert(iface interface{}) error {
	t := reflect.TypeOf(iface).Elem()
	v := reflect.ValueOf(iface).Elem()
	tableName := structToTableName(t.Name())
	var columns []string
	var values []string

	for i := 0; i < v.NumField(); i++ {
		columns = append(columns, structToTableName(t.Field(i).Name))
		f := v.Field(i)
		switch f.Interface().(type) {
		case string:
			values = append(values, fmt.Sprintf("'%v'", f))
		default:
			values = append(values, fmt.Sprintf("%v", f))
		}
	}

	statement := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s);", tableName, strings.Join(columns, ", "), strings.Join(values, ", "))
	fmt.Println(statement)
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
