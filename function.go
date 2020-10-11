package gooq

type Function struct{}

func (f *Function) Avg(column string) string {
	return "AVG(" + column + ")"
}

func (f *Function) Count(column string) string {
	return "COUNT(" + column + ")"
}

func (f *Function) Max(column string) string {
	return "MAX(" + column + ")"
}

func (f *Function) Min(column string) string {
	return "MIN(" + column + ")"
}

func (f *Function) Sum(column string) string {
	return "SUM(" + column + ")"
}

func (f *Function) Ascii(column string) string {
	return "ASCII('" + column + "')"
}

func (f *Function) Char(column string) string {
	return "CHAR(" + column + ")"
}

func (f *Function) Concat(columns ...string) string {
	return "CONCAT(" + mergeColumns(columns) + ")"
}

func (f *Function) Length(column string) string {
	return "LENGTH(" + column + ")"
}

func (f *Function) Replace(column, from, to string) string {
	return "REPLACE(" + column + ", '" + from + "', '" + to + "')"
}

func (f *Function) Ucase(column string) string {
	return "UCASE(" + column + ")"
}

func (f *Function) Lcase(column string) string {
	return "LCASE(" + column + ")"
}

// MID(column, start, [,length])
func (f *Function) Mid(column, start string, length ...string) string {
	l := ""
	if len(length) > 0 {
		l = ", " + length[0]
	}
	return "MID(" + column + ", " + start + l + ")"
}

func (f *Function) Abs(number string) string {
	return "ABS(" + number + ")"
}

func (f *Function) Ceil(number string) string {
	return "CEIL(" + number + ")"
}

func (f *Function) Floor(number string) string {
	return "FLOOR(" + number + ")"
}

func (f *Function) Power(number, power string) string {
	return "POWER(" + number + ", " + power + ")"
}

func (f *Function) Round(column, decimals string) string {
	return "ROUND(" + column + ", " + decimals + ")"
}

func (f *Function) Sqrt(number string) string {
	return "SQRT(" + number + ")"
}

func (f *Function) Pi() string {
	return "PI()"
}

func (f *Function) Exp(number string) string {
	return "EXP(" + number + ")"
}

// LOG(base, number)
func (f *Function) Log(base, number string) string {
	return "LOG(" + base + ", " + number + ")"
}

func (f *Function) Log2(number string) string {
	return "LOG2(" + number + ")"
}

func (f *Function) Log10(number string) string {
	return "LOG10(" + number + ")"
}

func (f *Function) Ln(number string) string {
	return "LN(" + number + ")"
}
