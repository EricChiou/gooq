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
