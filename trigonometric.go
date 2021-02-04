package gooq

// Sin SIN( number )
func (f *Function) Sin(number string) string {
	return "SIN(" + number + ")"
}

// Asin ASIN( number )
func (f *Function) Asin(number string) string {
	return "ASIN(" + number + ")"
}

// Cos COS( number )
func (f *Function) Cos(number string) string {
	return "COS(" + number + ")"
}

// Acos ACOS( number )
func (f *Function) Acos(number string) string {
	return "ACOS(" + number + ")"
}

// Tan TAN( number )
func (f *Function) Tan(number string) string {
	return "TAN(" + number + ")"
}

// Atan ATAN( number )
func (f *Function) Atan(number string) string {
	return "ATAN(" + number + ")"
}

// Cot COT( number )
func (f *Function) Cot(number string) string {
	return "COT(" + number + ")"
}

// Degree DEGREE( number )
func (f *Function) Degree(number string) string {
	return "DEGREE(" + number + ")"
}

// Radians RADIANS( number )
func (f *Function) Radians(number string) string {
	return "RADIANS(" + number + ")"
}
