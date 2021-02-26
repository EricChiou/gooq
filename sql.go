package gooq

import "strings"

// SQL sql
type SQL struct {
	sentence string
}

// GetSQL get full sql string
func (sql *SQL) GetSQL(columns ...string) string {
	return strings.TrimSpace(sql.sentence)
}

// Add other sql string
func (sql *SQL) Add(s string) *SQL {
	sql.sentence += " " + s
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
