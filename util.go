package gooq

import (
	"fmt"
	"strings"
)

func mergeValues(vs []interface{}) string {
	s := ""
	for _, v := range vs {
		s += convert2String(v) + ", "
	}
	return strings.TrimSuffix(s, ", ")
}

func convert2String(value interface{}) string {
	switch value.(type) {
	case string:
		if value == "?" {
			return fmt.Sprintf("%v", value)
		}
		return "'" + fmt.Sprintf("%v", value) + "'"

	default:
		return fmt.Sprintf("%v", value)
	}
}
