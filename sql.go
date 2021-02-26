package gooq

import "strings"

// SQL sql
type SQL struct {
	sentence string
	f        Function
}

// GetSQL get full sql string
func (sql *SQL) GetSQL(columns ...string) string {
	return strings.TrimSpace(sql.sentence)
}

// Add other sql string
func (sql *SQL) Add(str string) *SQL {
	sql.sentence += " " + str
	return sql
}

// S other sql string
func (sql *SQL) S(str string) *SQL {
	sql.sentence += " " + str
	return sql
}

// C add comma
func (sql *SQL) C(str string) *SQL {
	sql.sentence += ", " + str
	return sql
}

// Insert INSERT INTO table ( column1, column2, ... )
func (sql *SQL) Insert(table string, columns ...string) *SQL {
	sql.sentence += Insert(table, columns...)
	return sql
}

// Values VALUES ( value1, value2, ... )
func (sql *SQL) Values(values ...interface{}) *SQL {
	sql.sentence += Values(values...)
	return sql
}

// Update UPDATE table
func (sql *SQL) Update(table string) *SQL {
	sql.sentence += Update(table)
	return sql
}

// Set SET condition1, condition2, ...
func (sql *SQL) Set(conditions ...string) *SQL {
	sql.sentence += Set(conditions...)
	return sql
}

// Delete DELETE FROM table
func (sql *SQL) Delete(table string) *SQL {
	sql.sentence += Delete(table)
	return sql
}

// Into INTO table
func (sql *SQL) Into(table string) *SQL {
	sql.sentence += Into(table)
	return sql
}

// Select SELECT column1, column2, ...
func (sql *SQL) Select(columns ...string) *SQL {
	sql.sentence += Select(columns...)
	return sql
}

// SelectDistinct SELECT DISTINCT column1, column2, ...
func (sql *SQL) SelectDistinct(columns ...string) *SQL {
	sql.sentence += SelectDistinct(columns...)
	return sql
}

// SelectTop SQL Server, SELECT TOP number nil|PERCENT column1, column2, ...
func (sql *SQL) SelectTop(number string, percent bool, columns ...string) *SQL {
	sql.sentence += SelectTop(number, percent, columns...)
	return sql
}

// With WITH table
func (sql *SQL) With(table string) *SQL {
	sql.sentence += With(table)
	return sql
}

// From FROM table
func (sql *SQL) From(table string) *SQL {
	sql.sentence += From(table)
	return sql
}

// Where WHERE condition
func (sql *SQL) Where(c string) *SQL {
	sql.sentence += Where(c)
	return sql
}

// And AND condition
func (sql *SQL) And(c string) *SQL {
	sql.sentence += And(c)
	return sql
}

// Or OR condition
func (sql *SQL) Or(c string) *SQL {
	sql.sentence += Or(c)
	return sql
}

// OrderBy ORDER BY column1 ASC|DESC, column2 ASC|DESC, ...
func (sql *SQL) OrderBy(columns ...string) *SQL {
	sql.sentence += OrderBy(columns...)
	return sql
}

// Limit MySQL, LIMIT index, number
func (sql *SQL) Limit(index, number string) *SQL {
	sql.sentence += Limit(index, number)
	return sql
}

// Join INNER JOIN table
func (sql *SQL) Join(table string) *SQL {
	sql.sentence += Join(table)
	return sql
}

// LeftJoin LEFT JOIN table
func (sql *SQL) LeftJoin(table string) *SQL {
	sql.sentence += LeftJoin(table)
	return sql
}

// RightJoin RIGHT JOIN table
func (sql *SQL) RightJoin(table string) *SQL {
	sql.sentence += RightJoin(table)
	return sql
}

// FullJoin FULL JOIN table
func (sql *SQL) FullJoin(table string) *SQL {
	sql.sentence += FullJoin(table)
	return sql
}

// CrossJoin CROSS JOIN table
func (sql *SQL) CrossJoin(table string) *SQL {
	sql.sentence += CrossJoin(table)
	return sql
}

// NaturalJoin NATURAL JOIN table
func (sql *SQL) NaturalJoin(table string) *SQL {
	sql.sentence += NaturalJoin(table)
	return sql
}

// On ON condition
func (sql *SQL) On(c string) *SQL {
	sql.sentence += On(c)
	return sql
}

// Using USING (column)
func (sql *SQL) Using(column string) *SQL {
	sql.sentence += Using(column)
	return sql
}

// Union UNION ...
func (sql *SQL) Union() *SQL {
	sql.sentence += Union()
	return sql
}

// UnionAll UNION ALL ...
func (sql *SQL) UnionAll() *SQL {
	sql.sentence += UnionAll()
	return sql
}

// Intersect INTERSECT ...
func (sql *SQL) Intersect() *SQL {
	sql.sentence += Intersect()
	return sql
}

// Minus MINUS ...
func (sql *SQL) Minus() *SQL {
	sql.sentence += Minus()
	return sql
}

// Exists EXISTS ...
func (sql *SQL) Exists() *SQL {
	sql.sentence += Exists()
	return sql
}

// NotExists NOT EXISTS ...
func (sql *SQL) NotExists() *SQL {
	sql.sentence += NotExists()
	return sql
}

// GroupBy GROUP BY column1, column2, ...
func (sql *SQL) GroupBy(columns ...string) *SQL {
	sql.sentence += GroupBy(columns...)
	return sql
}

// Having HAVING ...
func (sql *SQL) Having(c string) *SQL {
	sql.sentence += Having(c)
	return sql
}

// Lp left parenthesis
func (sql *SQL) Lp() *SQL {
	sql.sentence += Lp()
	return sql
}

// Rp right parenthesis
func (sql *SQL) Rp() *SQL {
	sql.sentence += Rp()
	return sql
}

// As AS ...
func (sql *SQL) As(as string) *SQL {
	sql.sentence += " " + as
	return sql
}

// Condition

// Eq = value
func (sql *SQL) Eq(value interface{}) *SQL {
	sql.sentence += "=" + convert2String(value)
	return sql
}

// Gt > value
func (sql *SQL) Gt(value interface{}) *SQL {
	sql.sentence += ">" + convert2String(value)
	return sql
}

// Lt < value
func (sql *SQL) Lt(value interface{}) *SQL {
	sql.sentence += "<" + convert2String(value)
	return sql
}

// GtEq >= value
func (sql *SQL) GtEq(value interface{}) *SQL {
	sql.sentence += ">=" + convert2String(value)
	return sql
}

// LtEq column <= value
func (sql *SQL) LtEq(value interface{}) *SQL {
	sql.sentence += "<=" + convert2String(value)
	return sql
}

// In IN ( value1, value2, ...)
func (sql *SQL) In(values ...interface{}) *SQL {
	sql.sentence += " IN ( " + mergeValues(values) + " )"
	return sql
}

// NotIn NOT IN ( value1, value2, ...)
func (sql *SQL) NotIn(values ...interface{}) *SQL {
	sql.sentence += " NOT IN ( " + mergeValues(values) + " )"
	return sql
}

// Between BETWEEN value1 AND value2
func (sql *SQL) Between(value1, value2 interface{}) *SQL {
	sql.sentence += " BETWEEN " + convert2String(value1) + " AND " + convert2String(value2)
	return sql
}

// NotBetween NOT BETWEEN value1 AND value2
func (sql *SQL) NotBetween(value1, value2 interface{}) *SQL {
	sql.sentence += " NOT BETWEEN " + convert2String(value1) + " AND " + convert2String(value2)
	return sql
}

// Like LIKE pattern
func (sql *SQL) Like(pattern interface{}) *SQL {
	sql.sentence += " LIKE " + convert2String(pattern)
	return sql
}

// NotLike NOT LIKE pattern
func (sql *SQL) NotLike(pattern interface{}) *SQL {
	sql.sentence += " NOT LIKE " + convert2String(pattern)
	return sql
}

// IsNull IS NULL
func (sql *SQL) IsNull() *SQL {
	sql.sentence += " IS NULL"
	return sql
}

// IsNotNull IS NOT NULL
func (sql *SQL) IsNotNull() *SQL {
	sql.sentence += " IS NOT NULL"
	return sql
}

// Function

// Avg AVG(column)
func (sql *SQL) Avg(column string) *SQL {
	sql.sentence += sql.f.Avg(column)
	return sql
}

// Count COUNT(column)
func (sql *SQL) Count(column string) *SQL {
	sql.sentence += sql.f.Count(column)
	return sql
}

// Max MAX(column)
func (sql *SQL) Max(column string) *SQL {
	sql.sentence += sql.f.Max(column)
	return sql
}

// Min MIN(column)
func (sql *SQL) Min(column string) *SQL {
	sql.sentence += sql.f.Min(column)
	return sql
}

// Sum SUM(column)
func (sql *SQL) Sum(column string) *SQL {
	sql.sentence += sql.f.Sum(column)
	return sql
}

// ASCII ASCII('column')
func (sql *SQL) ASCII(column string) *SQL {
	sql.sentence += sql.f.ASCII(column)
	return sql
}

// Char CHAR(column)
func (sql *SQL) Char(column string) *SQL {
	sql.sentence += sql.f.Char(column)
	return sql
}

// Concat CONCAT(column1, column2, ...)
func (sql *SQL) Concat(columns ...string) *SQL {
	sql.sentence += sql.f.Concat(columns...)
	return sql
}

// Length LENGTH(column)
func (sql *SQL) Length(column string) *SQL {
	sql.sentence += sql.f.Length(column)
	return sql
}

// Replace REPLACE(column, 'from', 'to')
func (sql *SQL) Replace(column, from, to string) *SQL {
	sql.sentence += sql.f.Replace(column, from, to)
	return sql
}

// Ucase UCASE(column)
func (sql *SQL) Ucase(column string) *SQL {
	sql.sentence += sql.f.Ucase(column)
	return sql
}

// Lcase LCASE(column)
func (sql *SQL) Lcase(column string) *SQL {
	sql.sentence += sql.f.Lcase(column)
	return sql
}

// Mid MID(column, start, [,length])
func (sql *SQL) Mid(column, start string, length ...string) *SQL {
	sql.sentence += sql.f.Mid(column, start, length...)
	return sql
}

// Abs ABS(number)
func (sql *SQL) Abs(number string) *SQL {
	sql.sentence += sql.f.Abs(number)
	return sql
}

// Ceil CEIL(number)
func (sql *SQL) Ceil(number string) *SQL {
	sql.sentence += sql.f.Ceil(number)
	return sql
}

// Floor FLOOR(number)
func (sql *SQL) Floor(number string) *SQL {
	sql.sentence += sql.f.Floor(number)
	return sql
}

// Power POWER(number, power)
func (sql *SQL) Power(number, power string) *SQL {
	sql.sentence += sql.f.Power(number, power)
	return sql
}

// Round ROUND(column, decimals)
func (sql *SQL) Round(column, decimals string) *SQL {
	sql.sentence += sql.f.Round(column, decimals)
	return sql
}

// Sqrt SQRT(number)
func (sql *SQL) Sqrt(number string) *SQL {
	sql.sentence += sql.f.Sqrt(number)
	return sql
}

// Pi PI()
func (sql *SQL) Pi() *SQL {
	sql.sentence += sql.f.Pi()
	return sql
}

// Exp EXP(number)
func (sql *SQL) Exp(number string) *SQL {
	sql.sentence += sql.f.Exp(number)
	return sql
}

// Log LOG(base, number)
func (sql *SQL) Log(base, number string) *SQL {
	sql.sentence += sql.f.Log(base, number)
	return sql
}

// Log2 LOG2(number)
func (sql *SQL) Log2(number string) *SQL {
	sql.sentence += sql.f.Log2(number)
	return sql
}

// Log10 LOG10(number)
func (sql *SQL) Log10(number string) *SQL {
	sql.sentence += sql.f.Log10(number)
	return sql
}

// Ln LN(number)
func (sql *SQL) Ln(number string) *SQL {
	sql.sentence += sql.f.Ln(number)
	return sql
}
