package gooq

// Function SQL Function
type Function struct{}

// Avg AVG(column)
func (f *Function) Avg(column string) string {
	return "AVG(" + column + ")"
}

// Count COUNT(column)
func (f *Function) Count(column string) string {
	return "COUNT(" + column + ")"
}

// Max MAX(column)
func (f *Function) Max(column string) string {
	return "MAX(" + column + ")"
}

// Min MIN(column)
func (f *Function) Min(column string) string {
	return "MIN(" + column + ")"
}

// Sum SUM(column)
func (f *Function) Sum(column string) string {
	return "SUM(" + column + ")"
}

// ASCII ASCII('column')
func (f *Function) ASCII(column string) string {
	return "ASCII('" + column + "')"
}

// Char CHAR(column)
func (f *Function) Char(column string) string {
	return "CHAR(" + column + ")"
}

// Concat CONCAT(column1, column2, ...)
func (f *Function) Concat(columns ...string) string {
	return "CONCAT(" + mergeColumns(columns) + ")"
}

// Length LENGTH(column)
func (f *Function) Length(column string) string {
	return "LENGTH(" + column + ")"
}

// Replace REPLACE(column, 'from', 'to')
func (f *Function) Replace(column, from, to string) string {
	return "REPLACE(" + column + ", '" + from + "', '" + to + "')"
}

// Ucase UCASE(column)
func (f *Function) Ucase(column string) string {
	return "UCASE(" + column + ")"
}

// Lcase LCASE(column)
func (f *Function) Lcase(column string) string {
	return "LCASE(" + column + ")"
}

// Mid MID(column, start, [,length])
func (f *Function) Mid(column, start string, length ...string) string {
	l := ""
	if len(length) > 0 {
		l = ", " + length[0]
	}
	return "MID(" + column + ", " + start + l + ")"
}

// Abs ABS(number)
func (f *Function) Abs(number string) string {
	return "ABS(" + number + ")"
}

// Ceil CEIL(number)
func (f *Function) Ceil(number string) string {
	return "CEIL(" + number + ")"
}

// Floor FLOOR(number)
func (f *Function) Floor(number string) string {
	return "FLOOR(" + number + ")"
}

// Power POWER(number, power)
func (f *Function) Power(number, power string) string {
	return "POWER(" + number + ", " + power + ")"
}

// Round ROUND(column, decimals)
func (f *Function) Round(column, decimals string) string {
	return "ROUND(" + column + ", " + decimals + ")"
}

// Sqrt SQRT(number)
func (f *Function) Sqrt(number string) string {
	return "SQRT(" + number + ")"
}

// Pi PI()
func (f *Function) Pi() string {
	return "PI()"
}

// Exp EXP(number)
func (f *Function) Exp(number string) string {
	return "EXP(" + number + ")"
}

// Log LOG(base, number)
func (f *Function) Log(base, number string) string {
	return "LOG(" + base + ", " + number + ")"
}

// Log2 LOG2(number)
func (f *Function) Log2(number string) string {
	return "LOG2(" + number + ")"
}

// Log10 LOG10(number)
func (f *Function) Log10(number string) string {
	return "LOG10(" + number + ")"
}

// Ln LN(number)
func (f *Function) Ln(number string) string {
	return "LN(" + number + ")"
}
