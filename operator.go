package gooq

import "strings"

// Insert INSERT INTO table ( column1, column2, ... )
func Insert(table string, columns ...string) string {
	s := strings.Join(columns, ", ")
	if len(s) > 0 {
		return " INSERT INTO " + table + " ( " + s + " )"
	}
	return " INSERT INTO " + table
}

// Values VALUES ( value1, value2, ... )
func Values(values ...interface{}) string {
	return " VALUES ( " + mergeValues(values) + " )"
}

// Update Update table
func Update(table string) string {
	return " UPDATE " + table
}

// Set SET condition1, condition2, ...
func Set(conditions ...string) string {
	return " SET " + strings.Join(conditions, ", ")
}

// Delete DELETE FROM table
func Delete(table string) string {
	return " DELETE FROM " + table
}

// Into INTO table
func Into(table string) string {
	return " INTO " + table
}
