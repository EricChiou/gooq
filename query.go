package gooq

// SELECT column1, column2, ...
func Select(columns ...string) string {
	return " SELECT " + mergeColumns(columns)
}

// SELECT DISTINCT column1, column2, ...
func SelectDistinct(columns ...string) string {
	return " SELECT DISTINCT " + mergeColumns(columns)
}

// SQL Server, SELECT TOP number nil|PERCENT column1, column2, ...
func SelectTop(number string, percent bool, columns ...string) string {
	p := " "
	if percent {
		p = " PERCENT "
	}
	return " SELECT TOP " + number + p + mergeColumns(columns)
}

// WITH table
func With(table string) string {
	return " WITH " + table
}

func From(table string) string {
	return " FROM " + table
}

func Where(c condition) string {
	return " WHERE " + string(c)
}

func And(c condition) string {
	return " AND " + string(c)
}

func Or(c condition) string {
	return " OR " + string(c)
}

// ORDER BY column1 ASC|DESC, column2 ASC|DESC, ...
func OrderBy(columns ...string) string {
	return " ORDER BY " + mergeColumns(columns)
}

// MySQL, LIMIT index, number
func Limit(index, number string) string {
	return " LIMIT " + index + ", " + number
}

// INNER JOIN table
func Join(table string) string {
	return " INNER JOIN " + table
}

func LeftJoin(table string) string {
	return " LEFT JOIN " + table
}

func RightJoin(table string) string {
	return " RIGHT JOIN " + table
}

func FullJoin(table string) string {
	return " FULL JOIN " + table
}

func CrossJoin(table string) string {
	return " CROSS JOIN " + table
}

func NatiralJoin(table string) string {
	return " NATURAL JOIN " + table
}

func On(c condition) string {
	return " ON " + string(c)
}

func Using(column string) string {
	return " USING (" + column + ")"
}

func Union() string {
	return " UNION"
}

func UnionAll() string {
	return " UNION ALL"
}

func Intersect() string {
	return " INTERSECT"
}

func Minus() string {
	return " MINUS"
}

func Exists() string {
	return " EXISTS"
}

func NotExists() string {
	return " NOT EXISTS"
}

// left parenthesis
func Lp() string {
	return " ("
}

// right parenthesis
func Rp() string {
	return " )"
}

// GROUP BY column1, column2, ...
func GroupBy(columns ...string) string {
	return " GROUP BY " + mergeColumns(columns)
}

func Having(c condition) string {
	return " HAVING " + string(c)
}
