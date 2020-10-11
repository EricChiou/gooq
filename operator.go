package gooq

// INSERT INTO table ( column1, column2, ... )
func Insert(table string, columns ...string) string {
	s := mergeColumns(columns)
	if len(s) > 0 {
		return " INSERT INTO " + table + " ( " + s + " )"
	}
	return " INSERT INTO " + table
}

// VALUES ( value1, value2, ... )
func Values(values ...interface{}) string {
	return " VALUES ( " + mergeValues(values) + " )"
}

func Update(table string) string {
	return " UPDATE " + table
}

// SET condition1, condition2, ...
func Set(conditions ...Condition) string {
	return " SET " + mergeConditions(conditions)
}

// DELETE FROM table
func Delete(table string) string {
	return " DELETE FROM " + table
}

// INTO table
func INTO(table string) string {
	return " INTO " + table
}
