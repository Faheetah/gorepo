package gorepo

import (
	"regexp"
	"strings"
)

var table struct {
	Name string
}

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

// func Create(iface interface{}) error {
// name := structToTableName(iface)
// return nil
// }

// structToTable converts a struct to snake case
func structToTableName(table string) string {
	// This approach is a bit heavy handed, but works for now
	name := matchFirstCap.ReplaceAllString(table, "${1}_${2}")
	name = matchAllCap.ReplaceAllString(name, "${1}_${2}")
	return strings.ToLower(name)
}
