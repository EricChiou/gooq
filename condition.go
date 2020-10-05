package gooq

type condition string

// data = value
func Eq(data string, value interface{}) condition {
	return condition(data + " = " + convert2String(value))
}

// data > value
func Gt(data string, value interface{}) condition {
	return condition(data + " > " + convert2String(value))
}

// data < value
func Lt(data string, value interface{}) condition {
	return condition(data + " < " + convert2String(value))
}

// data >= value
func GtEq(data string, value interface{}) condition {
	return condition(data + " >= " + convert2String(value))
}

// data <= value
func LtEq(data string, value interface{}) condition {
	return condition(data + " <= " + convert2String(value))
}

// data IN ( value1, value2, ...)
func In(data string, values ...interface{}) condition {
	return condition(data + " IN ( " + mergeValues(values) + " )")
}

// data NOT IN ( value1, value2, ...)
func NotIn(data string, values ...interface{}) condition {
	return condition(data + " NOT IN ( " + mergeValues(values) + " )")
}

// data BETWEEN value1 AND value2
func Between(data string, value1, value2 interface{}) condition {
	return condition(data + " BETWEEN " + convert2String(value1) + " AND " + convert2String(value2))
}

// data NOT BETWEEN value1 AND value2
func NotBetween(data string, value1, value2 interface{}) condition {
	return condition(data + " NOT BETWEEN " + convert2String(value1) + " AND " + convert2String(value2))
}

// data LIKE pattern
func Like(data, pattern string) condition {
	return condition(data + " LIKE " + pattern)
}

// data NOT LIKE pattern
func NotLike(data, pattern string) condition {
	return condition(data + " NOT LIKE " + pattern)
}

// data IS NULL
func IsNull(data string) condition {
	return condition(data + " IS NULL")
}

// data IS NOT NULL
func IsNotNull(data string) condition {
	return condition(data + " IS NOT NULL")
}

// column ASC
func Asc(column string) string {
	return column + " ASC"
}

// column DESC
func Desc(column string) string {
	return column + " DESC"
}

// column AS as
func As(column, as string) string {
	return column + " AS " + as
}
