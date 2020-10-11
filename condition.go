package gooq

type condition string

// column = value
func Eq(column string, value interface{}) condition {
	return condition(column + " = " + convert2String(value))
}

// column > value
func Gt(column string, value interface{}) condition {
	return condition(column + " > " + convert2String(value))
}

// column < value
func Lt(column string, value interface{}) condition {
	return condition(column + " < " + convert2String(value))
}

// column >= value
func GtEq(column string, value interface{}) condition {
	return condition(column + " >= " + convert2String(value))
}

// column <= value
func LtEq(column string, value interface{}) condition {
	return condition(column + " <= " + convert2String(value))
}

// column IN ( value1, value2, ...)
func In(column string, values ...interface{}) condition {
	return condition(column + " IN ( " + mergeValues(values) + " )")
}

// column NOT IN ( value1, value2, ...)
func NotIn(column string, values ...interface{}) condition {
	return condition(column + " NOT IN ( " + mergeValues(values) + " )")
}

// column BETWEEN value1 AND value2
func Between(column string, value1, value2 interface{}) condition {
	return condition(column + " BETWEEN " + convert2String(value1) + " AND " + convert2String(value2))
}

// column NOT BETWEEN value1 AND value2
func NotBetween(column string, value1, value2 interface{}) condition {
	return condition(column + " NOT BETWEEN " + convert2String(value1) + " AND " + convert2String(value2))
}

// column LIKE pattern
func Like(column, pattern string) condition {
	return condition(column + " LIKE " + pattern)
}

// column NOT LIKE pattern
func NotLike(column, pattern string) condition {
	return condition(column + " NOT LIKE " + pattern)
}

// column IS NULL
func IsNull(column string) condition {
	return condition(column + " IS NULL")
}

// column IS NOT NULL
func IsNotNull(column string) condition {
	return condition(column + " IS NOT NULL")
}
