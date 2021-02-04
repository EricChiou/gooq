package gooq

// Condition condition enum
type Condition string

// Eq column = value
func Eq(column string, value interface{}) Condition {
	return Condition(column + "=" + convert2String(value))
}

// Gt column > value
func Gt(column string, value interface{}) Condition {
	return Condition(column + ">" + convert2String(value))
}

// Lt column < value
func Lt(column string, value interface{}) Condition {
	return Condition(column + "<" + convert2String(value))
}

// GtEq column >= value
func GtEq(column string, value interface{}) Condition {
	return Condition(column + ">=" + convert2String(value))
}

// LtEq column <= value
func LtEq(column string, value interface{}) Condition {
	return Condition(column + "<=" + convert2String(value))
}

// In column IN ( value1, value2, ...)
func In(column string, values ...interface{}) Condition {
	return Condition(column + " IN ( " + mergeValues(values) + " )")
}

// NotIn column NOT IN ( value1, value2, ...)
func NotIn(column string, values ...interface{}) Condition {
	return Condition(column + " NOT IN ( " + mergeValues(values) + " )")
}

// Between column BETWEEN value1 AND value2
func Between(column string, value1, value2 interface{}) Condition {
	return Condition(column + " BETWEEN " + convert2String(value1) + " AND " + convert2String(value2))
}

// NotBetween column NOT BETWEEN value1 AND value2
func NotBetween(column string, value1, value2 interface{}) Condition {
	return Condition(column + " NOT BETWEEN " + convert2String(value1) + " AND " + convert2String(value2))
}

// Like column LIKE pattern
func Like(column, pattern string) Condition {
	return Condition(column + " LIKE '" + pattern + "'")
}

// NotLike column NOT LIKE pattern
func NotLike(column, pattern string) Condition {
	return Condition(column + " NOT LIKE '" + pattern + "'")
}

// IsNull column IS NULL
func IsNull(column string) Condition {
	return Condition(column + " IS NULL")
}

// IsNotNull column IS NOT NULL
func IsNotNull(column string) Condition {
	return Condition(column + " IS NOT NULL")
}
