package gooq

func Avg(column string) string {
	return "AVG(" + column + ")"
}

func Count(column string) string {
	return "COUNT(" + column + ")"
}

func Max(column string) string {
	return "MAX(" + column + ")"
}

func Min(column string) string {
	return "MIN(" + column + ")"
}

func Sum(column string) string {
	return "SUM(" + column + ")"
}

func Ascii(column string) string {
	return "ASCII('" + column + "')"
}

func Char(column string) string {
	return "CHAR(" + column + ")"
}

func Concat(columns ...string) string {
	return "CONCAT(" + mergeColumns(columns) + ")"
}

func Length(column string) string {
	return "LENGTH(" + column + ")"
}

func Replace(column, from, to string) string {
	return "REPLACE(" + column + ", '" + from + "', '" + to + "')"
}

func Ucase(column string) string {
	return "UCASE(" + column + ")"
}

func Lcase(column string) string {
	return "LCASE(" + column + ")"
}

// MID(column, start, [,length])
func Mid(column, start string, length ...string) string {
	l := ""
	if len(length) > 0 {
		l = ", " + length[0]
	}
	return "MID(" + column + ", " + start + l + ")"
}

func Abs(number string) string {
	return "ABS(" + number + ")"
}

func Ceil(number string) string {
	return "CEIL(" + number + ")"
}

func Floor(number string) string {
	return "FLOOR(" + number + ")"
}

func Power(number, power string) string {
	return "POWER(" + number + ", " + power + ")"
}

func Round(column, decimals string) string {
	return "ROUND(" + column + ", " + decimals + ")"
}

func Sqrt(number string) string {
	return "SQRT(" + number + ")"
}

func Pi() string {
	return "PI()"
}

func Exp(number string) string {
	return "EXP(" + number + ")"
}

// LOG(base, number)
func Log(base, number string) string {
	return "LOG(" + base + ", " + number + ")"
}

func Log2(number string) string {
	return "LOG2(" + number + ")"
}

func Log10(number string) string {
	return "LOG10(" + number + ")"
}

func Ln(number string) string {
	return "LN(" + number + ")"
}
