package gorepo

import (
	"fmt"
	_ "testing"

	_ "github.com/mattn/go-sqlite3"
)

type ExampleRepoTable struct {
	Id      int
	Name    string
	IsAdmin bool
}

func ExampleRepo() {
	example := ExampleRepoTable{
		Name: "foo",
	}
	repo := Repo{}

	err := repo.Insert(&example)
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// INSERT INTO example_repo_table (id, name, is_admin) VALUES (0, 'foo', false);
}
