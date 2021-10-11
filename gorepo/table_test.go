package gorepo

import (
	"testing"
)

func TestStructToTable(t *testing.T) {
	var want string
	var got string

	want = "some_example_table_name"
	got = structToTableName("SOMEExampleTableNAME")
	if want != got {
		t.Errorf("Expected '%s', got '%s'", want, got)
	}

	want = "some_other_example"
	got = structToTableName("SomeOTHERExample")
	if want != got {
		t.Errorf("Expected '%s', got '%s'", want, got)
	}
}
