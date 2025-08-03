package lib

import (
	"fmt"
	"reflect"
	"strings"
	"unicode"

	"github.com/Bluestaq/udl-golang-sdk/option"
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
	EqualTo:              "=",
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

// toLowerCamelCase converts a string to lowerCamelCase
func toLowerCamelCase(s string) string {
	if s == "" {
		return s
	}

	// Handle the first character - make it lowercase
	result := strings.ToLower(string(s[0]))

	// Handle the rest of the string
	for i := 1; i < len(s); i++ {
		if unicode.IsUpper(rune(s[i])) {
			// For uppercase letters, keep them as is (this preserves camelCase)
			result += string(s[i])
		} else {
			// For lowercase letters, keep them as is
			result += string(s[i])
		}
	}

	return result
}

// NewQueryBuilder builds a new query builder for the given struct type
func NewQueryBuilder(model any) *QueryBuilder {
	fieldLookup := make(map[string]string)

	typ := reflect.TypeOf(model)
	if typ.Kind() == reflect.Pointer {
		typ = typ.Elem()
	}

	// Use `query` tag or convert field name to lowerCamelCase
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		queryTag := field.Tag.Get("query")
		if queryTag == "" {
			queryTag = toLowerCamelCase(field.Name)
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

	if _, ok := operatorSymbols[op]; !ok {
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
		// For standard operators, just use the value - the API handles operators differently
		q.filters[queryField] = fmt.Sprintf("%v", value)
	}

	return q
}

func (q *QueryBuilder) ToParams() map[string]string {
	return q.filters
}

// ToRequestOptions returns the query parameters as a slice of option.RequestOption
func (q *QueryBuilder) ToRequestOptions() []option.RequestOption {
	var opts []option.RequestOption
	for k, v := range q.filters {
		opts = append(opts, option.WithQuery(k, v))
	}
	return opts
}
