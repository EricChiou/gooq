// Gooq instance
type Gooq struct {
	SQL  SQL
	Args []interface{}
}

// AddValues add args
func (g *Gooq) AddValues(values ...interface{}) {
	g.Args = append(g.Args, values...)
}
