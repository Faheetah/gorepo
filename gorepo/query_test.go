package gorepo

import (
	"fmt"
	"os"
	_ "testing"

	_ "github.com/mattn/go-sqlite3"
)

type ExampleTable struct {
	Id   int
	Name string
}

func ExampleQuery() {
	example := ExampleTable{}
	repo := Repo{}

	q, err := repo.Filter(&example).Distinct().Count().Build()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(q.Statement)

	q, err = repo.Filter(&example).Select("id", "name").Where("name", LIKE, "$1", "bar").Limit(5).OrderBy("name", ASC).Build()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(q.Statement)
	fmt.Println("bindings:", q.Binds)
	// Output:
	// SELECT DISTINCT COUNT(*) FROM example_table;
	// SELECT id, name FROM example_table WHERE name LIKE $1 LIMIT 5 ORDER BY name ASC;
	// bindings: [bar]
}
