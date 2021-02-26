package gooq

// Column set column condition
func Column(column string) *ColCondition {
	return &ColCondition{column: column}
}

// ColCondition column condition struct
type ColCondition struct {
	column string
}

// Eq column = value
func (c *ColCondition) Eq(value interface{}) string {
	return Eq(c.column, value)
}

// Gt column > value
func (c *ColCondition) Gt(value interface{}) string {
	return Gt(c.column, value)
}

// Lt column < value
func (c *ColCondition) Lt(value interface{}) string {
	return Lt(c.column, value)
}

// GtEq column >= value
func (c *ColCondition) GtEq(value interface{}) string {
	return GtEq(c.column, value)
}

// LtEq column <= value
func (c *ColCondition) LtEq(value interface{}) string {
	return LtEq(c.column, value)
}

// In column IN ( value1, value2, ...)
func (c *ColCondition) In(values ...interface{}) string {
	return In(c.column, values...)
}

// NotIn column NOT IN ( value1, value2, ...)
func (c *ColCondition) NotIn(values ...interface{}) string {
	return NotIn(c.column, values...)
}

// Between column BETWEEN value1 AND value2
func (c *ColCondition) Between(value1, value2 interface{}) string {
	return Between(c.column, value1, value2)
}

// NotBetween column NOT BETWEEN value1 AND value2
func (c *ColCondition) NotBetween(value1, value2 interface{}) string {
	return NotBetween(c.column, value1, value2)
}

// Like column LIKE pattern
func (c *ColCondition) Like(pattern string) string {
	return Like(c.column, pattern)
}

// NotLike column NOT LIKE pattern
func (c *ColCondition) NotLike(pattern string) string {
	return NotLike(c.column, pattern)
}

// IsNull column IS NULL
func (c *ColCondition) IsNull() string {
	return IsNull(c.column)
}

// IsNotNull column IS NOT NULL
func (c *ColCondition) IsNotNull() string {
	return IsNotNull(c.column)
}
