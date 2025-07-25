package lib

import (
	"fmt"
	"reflect"
	"strings"
)

type QueryOperator string

const (
	EqualTo              QueryOperator = "_eq"
	GreaterThanOrEqualTo QueryOperator = "_gte"
	LessThanOrEqualTo    QueryOperator = "_lte"
	Like                 QueryOperator = "_like"
	Not                  QueryOperator = "_not"
	Between              QueryOperator = "_between"
)

var operatorSymbols = map[QueryOperator]string{
	EqualTo:              "",
	GreaterThanOrEqualTo: ">=",
	LessThanOrEqualTo:    "<=",
	Like:                 "like",
	Not:                  "not",
	Between:              "between",
}

type QueryBuilder struct {
	filters     map[string]string
	fieldLookup map[string]string // maps struct field names to query field names
}

// NewQueryBuilder builds a new query builder for the given struct type
func NewQueryBuilder(model any) *QueryBuilder {
	fieldLookup := make(map[string]string)

	typ := reflect.TypeOf(model)
	if typ.Kind() == reflect.Pointer {
		typ = typ.Elem()
	}

	// Use `query` tag or field name
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		queryTag := field.Tag.Get("query")
		if queryTag == "" {
			queryTag = strings.ToLower(field.Name)
		}
		fieldLookup[field.Name] = queryTag
	}

	return &QueryBuilder{
		filters:     make(map[string]string),
		fieldLookup: fieldLookup,
	}
}

// Add adds a filter with operator suffix
func (q *QueryBuilder) Add(fieldName string, op QueryOperator, value any) *QueryBuilder {
	queryField, ok := q.fieldLookup[fieldName]
	if !ok {
		panic(fmt.Sprintf("field %s not found in model", fieldName))
	}

	operator, ok := operatorSymbols[op]
	if !ok {
		panic(fmt.Sprintf("unsupported operator: %s", op))
	}

	switch op {
	case Like:
		q.filters[queryField] = fmt.Sprintf("*%v*", value)
	case Not:
		q.filters[queryField] = fmt.Sprintf("~%v", value)
	case Between:
		vals := reflect.ValueOf(value)
		if vals.Kind() != reflect.Slice || vals.Len() != 2 {
			panic("between requires a slice/array of exactly two values")
		}
		q.filters[queryField] = fmt.Sprintf("%v..%v", vals.Index(0), vals.Index(1))
	default:
		q.filters[queryField] = fmt.Sprintf("%s%v", operator, value)
	}

	return q
}

func (q *QueryBuilder) ToParams() map[string]string {
	return q.filters
}
