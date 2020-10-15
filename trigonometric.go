package gooq

func (f *Function) Sin(number string) string {
	return "SIN(" + number + ")"
}

func (f *Function) Asin(number string) string {
	return "ASIN(" + number + ")"
}

func (f *Function) Cos(number string) string {
	return "COS(" + number + ")"
}

func (f *Function) Acos(number string) string {
	return "ACOS(" + number + ")"
}

func (f *Function) Tan(number string) string {
	return "TAN(" + number + ")"
}

func (f *Function) Atan(number string) string {
	return "ATAN(" + number + ")"
}

func (f *Function) Cot(number string) string {
	return "COT(" + number + ")"
}

func (f *Function) Degree(number string) string {
	return "DEGREE(" + number + ")"
}

func (f *Function) Radians(number string) string {
	return "RADIANS(" + number + ")"
}
