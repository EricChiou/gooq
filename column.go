package gooq

func Column(column string) *condition {
	return &condition{column: column}
}

type condition struct {
	column string
}

// data = value
func (c *condition) Eq(value interface{}) Condition {
	return Eq(c.column, value)
}

// data > value
func (c *condition) Gt(value interface{}) Condition {
	return Gt(c.column, value)
}

// data < value
func (c *condition) Lt(value interface{}) Condition {
	return Lt(c.column, value)
}

// data >= value
func (c *condition) GtEq(value interface{}) Condition {
	return GtEq(c.column, value)
}

// data <= value
func (c *condition) LtEq(value interface{}) Condition {
	return LtEq(c.column, value)
}

// data IN ( value1, value2, ...)
func (c *condition) In(values ...interface{}) Condition {
	return In(c.column, values...)
}

// data NOT IN ( value1, value2, ...)
func (c *condition) NotIn(values ...interface{}) Condition {
	return NotIn(c.column, values...)
}

// data BETWEEN value1 AND value2
func (c *condition) Between(value1, value2 interface{}) Condition {
	return Between(c.column, value1, value2)
}

// data NOT BETWEEN value1 AND value2
func (c *condition) NotBetween(value1, value2 interface{}) Condition {
	return NotBetween(c.column, value1, value2)
}

// data LIKE pattern
func (c *condition) Like(pattern string) Condition {
	return Like(c.column, pattern)
}

// data NOT LIKE pattern
func (c *condition) NotLike(pattern string) Condition {
	return NotLike(c.column, pattern)
}

// data IS NULL
func (c *condition) IsNull() Condition {
	return IsNull(c.column)
}

// data IS NOT NULL
func (c *condition) IsNotNull() Condition {
	return IsNotNull(c.column)
}
