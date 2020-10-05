package gooq

import "strings"

type SQL struct {
	sentence string
}

// get full sql string
func (sql *SQL) GetSQL(columns ...string) string {
	return strings.TrimSpace(sql.sentence)
}

// Add other sql string
func (sql *SQL) S(s string) *SQL {
	sql.sentence += " " + s
	return sql
}

// INSERT INTO table ( column1, column2, ... )
func (sql *SQL) Insert(table string, columns ...string) *SQL {
	sql.sentence += Insert(table, columns...)
	return sql
}

// VALUES ( value1, value2, ... )
func (sql *SQL) Values(values ...interface{}) *SQL {
	sql.sentence += Values(values...)
	return sql
}

func (sql *SQL) Update(table string) *SQL {
	sql.sentence += Update(table)
	return sql
}

// SET condition1, condition2, ...
func (sql *SQL) Set(conditions ...condition) *SQL {
	sql.sentence += Set(conditions...)
	return sql
}

// DELETE FROM table
func (sql *SQL) Delete(table string) *SQL {
	sql.sentence += Delete(table)
	return sql
}

// INTO table
func (sql *SQL) INTO(table string) *SQL {
	sql.sentence += INTO(table)
	return sql
}

// SELECT column1, column2, ...
func (sql *SQL) Select(columns ...string) *SQL {
	sql.sentence += Select(columns...)
	return sql
}

// SELECT DISTINCT column1, column2, ...
func (sql *SQL) SelectDistinct(columns ...string) *SQL {
	sql.sentence += SelectDistinct(columns...)
	return sql
}

// SQL Server, SELECT TOP number nil|PERCENT column1, column2, ...
func (sql *SQL) SelectTop(number string, percent bool, columns ...string) *SQL {
	sql.sentence += SelectTop(number, percent, columns...)
	return sql
}

// WITH table
func (sql *SQL) With(table string) *SQL {
	sql.sentence += With(table)
	return sql
}

func (sql *SQL) From(table string) *SQL {
	sql.sentence += From(table)
	return sql
}

func (sql *SQL) Where(c condition) *SQL {
	sql.sentence += Where(c)
	return sql
}

func (sql *SQL) And(c condition) *SQL {
	sql.sentence += And(c)
	return sql
}

func (sql *SQL) Or(c condition) *SQL {
	sql.sentence += Or(c)
	return sql
}

// ORDER BY column1 ASC|DESC, column2 ASC|DESC, ...
func (sql *SQL) OrderBy(columns ...string) *SQL {
	sql.sentence += OrderBy(columns...)
	return sql
}

// MySQL, LIMIT index, number
func (sql *SQL) Limit(index, number string) *SQL {
	sql.sentence += Limit(index, number)
	return sql
}

// INNER JOIN table
func (sql *SQL) Join(table string) *SQL {
	sql.sentence += Join(table)
	return sql
}

func (sql *SQL) LeftJoin(table string) *SQL {
	sql.sentence += LeftJoin(table)
	return sql
}

func (sql *SQL) RightJoin(table string) *SQL {
	sql.sentence += RightJoin(table)
	return sql
}

func (sql *SQL) FullJoin(table string) *SQL {
	sql.sentence += FullJoin(table)
	return sql
}

func (sql *SQL) CrossJoin(table string) *SQL {
	sql.sentence += CrossJoin(table)
	return sql
}

func (sql *SQL) NaturalJoin(table string) *SQL {
	sql.sentence += NaturalJoin(table)
	return sql
}

func (sql *SQL) On(c condition) *SQL {
	sql.sentence += On(c)
	return sql
}

func (sql *SQL) Using(column string) *SQL {
	sql.sentence += Using(column)
	return sql
}

func (sql *SQL) Union() *SQL {
	sql.sentence += Union()
	return sql
}

func (sql *SQL) UnionAll() *SQL {
	sql.sentence += UnionAll()
	return sql
}

func (sql *SQL) Intersect() *SQL {
	sql.sentence += Intersect()
	return sql
}

func (sql *SQL) Minus() *SQL {
	sql.sentence += Minus()
	return sql
}

func (sql *SQL) Exists() *SQL {
	sql.sentence += Exists()
	return sql
}

func (sql *SQL) NotExists() *SQL {
	sql.sentence += NotExists()
	return sql
}

// GROUP BY column1, column2, ...
func (sql *SQL) GroupBy(columns ...string) *SQL {
	sql.sentence += GroupBy(columns...)
	return sql
}

func (sql *SQL) Having(c condition) *SQL {
	sql.sentence += Having(c)
	return sql
}

// left parenthesis
func (sql *SQL) Lp() *SQL {
	sql.sentence += Lp()
	return sql
}

// right parenthesis
func (sql *SQL) Rp() *SQL {
	sql.sentence += Rp()
	return sql
}
