package gooq

func Column(column string) *Condition {
	return &Condition{column: column}
}

type Condition struct {
	column string
}

// data = value
func (c *Condition) Eq(value interface{}) condition {
	return Eq(c.column, value)
}

// data > value
func (c *Condition) Gt(value interface{}) condition {
	return Gt(c.column, value)
}

// data < value
func (c *Condition) Lt(value interface{}) condition {
	return Lt(c.column, value)
}

// data >= value
func (c *Condition) GtEq(value interface{}) condition {
	return GtEq(c.column, value)
}

// data <= value
func (c *Condition) LtEq(value interface{}) condition {
	return LtEq(c.column, value)
}

// data IN ( value1, value2, ...)
func (c *Condition) In(values ...interface{}) condition {
	return In(c.column, values...)
}

// data NOT IN ( value1, value2, ...)
func (c *Condition) NotIn(values ...interface{}) condition {
	return NotIn(c.column, values...)
}

// data BETWEEN value1 AND value2
func (c *Condition) Between(value1, value2 interface{}) condition {
	return Between(c.column, value1, value2)
}

// data NOT BETWEEN value1 AND value2
func (c *Condition) NotBetween(value1, value2 interface{}) condition {
	return NotBetween(c.column, value1, value2)
}

// data LIKE pattern
func (c *Condition) Like(pattern string) condition {
	return Like(c.column, pattern)
}

// data NOT LIKE pattern
func (c *Condition) NotLike(pattern string) condition {
	return NotLike(c.column, pattern)
}

// data IS NULL
func (c *Condition) IsNull() condition {
	return IsNull(c.column)
}

// data IS NOT NULL
func (c *Condition) IsNotNull() condition {
	return IsNotNull(c.column)
}
