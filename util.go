package gooq

import (
	"fmt"
	"strings"
)

func mergeColumns(columns []string) string {
	s := ""
	for _, column := range columns {
		s += column + ", "
	}
	return strings.TrimSuffix(s, ", ")
}

func mergeValues(vs []interface{}) string {
	s := ""
	for _, v := range vs {
		s += convert2String(v) + ", "
	}
	return strings.TrimSuffix(s, ", ")
}

func mergeConditions(conditions []string) string {
	s := ""
	for _, condition := range conditions {
		s += condition + ", "
	}
	return strings.TrimSuffix(s, ", ")
}

func convert2String(value interface{}) string {
	return fmt.Sprintf("%v", value)
}
