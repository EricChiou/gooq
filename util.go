package gooq

import (
	"fmt"
	"reflect"
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

func mergeConditions(conditions []condition) string {
	s := ""
	for _, condition := range conditions {
		s += string(condition) + ", "
	}
	return strings.TrimSuffix(s, ", ")
}

func convert2String(value interface{}) string {
	if value != nil && reflect.TypeOf(value).Kind() == reflect.String {
		return fmt.Sprintf("'%s'", value)
	}
	return fmt.Sprintf("%v", value)
}
