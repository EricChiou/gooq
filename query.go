package gooq

import "strings"

// Select SELECT column1, column2, ...
func Select(columns ...string) string {
	return " SELECT " + strings.Join(columns, ", ")
}

// SelectDistinct SELECT DISTINCT column1, column2, ...
func SelectDistinct(columns ...string) string {
	return " SELECT DISTINCT " + strings.Join(columns, ", ")
}

// SelectTop SQL Server, SELECT TOP number nil|PERCENT column1, column2, ...
func SelectTop(number string, percent bool, columns ...string) string {
	p := " "
	if percent {
		p = " PERCENT "
	}
	return " SELECT TOP " + number + p + strings.Join(columns, ", ")
}

// With WITH table
func With(table string) string {
	return " WITH " + table
}

// From FROM table
func From(table string) string {
	return " FROM " + table
}

// Where WHERE condition
func Where(c string) string {
	return " WHERE " + c
}

// And AND condition
func And(c string) string {
	return " AND " + c
}

// Or OR condition
func Or(c string) string {
	return " OR " + c
}

// OrderBy ORDER BY column1 ASC|DESC, column2 ASC|DESC, ...
func OrderBy(columns ...string) string {
	return " ORDER BY " + strings.Join(columns, ", ")
}

// Limit MySQL, LIMIT index, number
func Limit(index, number string) string {
	return " LIMIT " + index + ", " + number
}

// Join INNER JOIN table
func Join(table string) string {
	return " INNER JOIN " + table
}

// LeftJoin LEFT JOIN table
func LeftJoin(table string) string {
	return " LEFT JOIN " + table
}

// RightJoin RIGHT JOIN table
func RightJoin(table string) string {
	return " RIGHT JOIN " + table
}

// FullJoin FULL JOIN table
func FullJoin(table string) string {
	return " FULL JOIN " + table
}

// CrossJoin CROSS JOIN table
func CrossJoin(table string) string {
	return " CROSS JOIN " + table
}

// NaturalJoin NATURAL JOIN table
func NaturalJoin(table string) string {
	return " NATURAL JOIN " + table
}

// On ON condition
func On(c string) string {
	return " ON " + c
}

// Using USING (column)
func Using(column string) string {
	return " USING (" + column + ")"
}

// Union UNION ...
func Union() string {
	return " UNION"
}

// UnionAll UNION ALL ...
func UnionAll() string {
	return " UNION ALL"
}

// Intersect INTERSECT ...
func Intersect() string {
	return " INTERSECT"
}

// Minus MINUS ...
func Minus() string {
	return " MINUS"
}

// Exists EXISTS ...
func Exists() string {
	return " EXISTS"
}

// NotExists NOT EXISTS ...
func NotExists() string {
	return " NOT EXISTS"
}

// GroupBy GROUP BY column1, column2, ...
func GroupBy(columns ...string) string {
	return " GROUP BY " + strings.Join(columns, ", ")
}

// Having HAVING condition
func Having(c string) string {
	return " HAVING " + c
}

// Lp left parenthesis
func Lp() string {
	return " ("
}

// Rp right parenthesis
func Rp() string {
	return " )"
}
