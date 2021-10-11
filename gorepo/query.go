package gorepo

import (
	"fmt"
	"strings"
)

type Query struct {
	table     string
	fields    []string
	distinct  bool
	count     bool
	orderBy   string
	where     string
	limit     int
	preloads  interface{}
	Binds     []string
	Error     error
	Result    interface{}
	Statement string
}

type Order string

const (
	ASC  Order = "ASC"
	DESC Order = "DESC"
)

type Cond string

const (
	LIKE Cond = "LIKE"
)

func (query *Query) Table(table string) *Query {
	query.table = structToTableName(table)
	return query
}

func (query *Query) Select(fields ...string) *Query {
	if query.Error != nil {
		return nil
	}

	query.fields = fields

	return query
}

func (query *Query) Distinct() *Query {
	if query.Error != nil {
		return nil
	}

	query.distinct = true

	return query
}

func (query *Query) Count() *Query {
	if query.Error != nil {
		return query
	}

	query.count = true

	return query
}

func (query *Query) OrderBy(field string, asc Order) *Query {
	if query.Error != nil {
		return query
	}

	query.orderBy = field + " " + string(asc)

	return query
}

func (query *Query) Where(left string, comp Cond, right string, binds ...string) *Query {
	if query.Error != nil {
		return query
	}

	query.where = fmt.Sprintf("%s %s %s", left, string(comp), right)
	query.Binds = append(query.Binds, binds...)
	return query
}

func (query *Query) Limit(limit int) *Query {
	if query.Error != nil {
		return query
	}

	query.limit = limit
	return query
}

func (query *Query) Preload(relation *interface{}) *Query {
	if query.Error != nil {
		return query
	}

	return query
}

func (query *Query) Build() (*Query, error) {
	statement := []string{"SELECT"}

	if query.distinct {
		statement = append(statement, "DISTINCT")
	}

	if query.count {
		statement = append(statement, "COUNT(*)")
	} else if len(query.fields) == 0 {
		statement = append(statement, "*")
	} else {
		statement = append(statement, strings.Join(query.fields, ", "))
	}

	statement = append(statement, "FROM")
	statement = append(statement, query.table)

	if query.where != "" {
		statement = append(statement, "WHERE "+query.where)
	}

	if query.limit != 0 {
		statement = append(statement, "LIMIT "+fmt.Sprintf("%d", query.limit))
	}

	if query.orderBy != "" {
		statement = append(statement, "ORDER BY "+query.orderBy)
	}

	query.Statement = strings.Join(statement, " ") + ";"
	return query, nil
}
