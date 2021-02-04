## gooq
Simply SQL ORM. Inspired by [JOOQ](https://www.jooq.org/).  

## How to use
#### Simple Example
```go
import (
	_sql "database/sql"
	"fmt"

	"github.com/EricChiou/gooq"
)

func main() {
	var f gooq.Function
	c := gooq.Column

	// INSERT
	var sql gooq.SQL
	sql.Insert("user", "acc", "pwd", "name").Values("my_account", 123456, "user_name")
	fmt.Println(sql.GetSQL())
	// result: INSERT INTO user ( acc, pwd, name ) VALUES ( 'my_account', 123456, 'user_name' )

	// UPDATE
	sql = gooq.SQL{}
	sql.Update("user").Set(c("pwd").Eq(111111), c("name").Eq("user_name2")).Where(c("id").Eq(1))
	fmt.Println(sql.GetSQL())
	// result: UPDATE user SET pwd = 111111, name = 'user_name2' WHERE id = 1

	// DELETE
	sql = gooq.SQL{}
	sql.Delete("user").Where(c("id").Eq(1))
	fmt.Println(sql.GetSQL())
	// result: DELETE FROM user WHERE id = 1

	// Query
	sql = gooq.SQL{}
	sql.Select("*").From("user").Where(c("id").Eq(1))
	fmt.Println(sql.GetSQL())
	// result: SELECT * FROM user WHERE id = 1

	sql = gooq.SQL{}
	sql.Select("acc", "pwd", "name").From("user").Where(c("id").Eq(1))
	fmt.Println(sql.GetSQL())
	// result: SELECT acc, pwd, name FROM user WHERE id = 1

	// INNER JOIN, Sub Query, ORDER BY, LIMIT (mysql), USING
	sql = gooq.SQL{}
	var subSql gooq.SQL
	sql.Select("id", "acc", "name").From("user")
	sql.Join(subSql.Lp().Select("id").From("user").OrderBy("id").Limit("0", "10").Rp().As("t").GetSQL())
	sql.Using("id").Where(c("status").Eq("active"))

	fmt.Println(sql.GetSQL())
	// result: SELECT id, acc, name FROM user INNER JOIN ( SELECT id FROM user ORDER BY id LIMIT 0, 10 ) t USING (id) WHERE status = 'active'

	// COUNT
	sql = gooq.SQL{}
	sql.Select(f.Count("*")).From("user").Where(c("status").Eq("active"))
	fmt.Println(sql.GetSQL())
	// result: SELECT COUNT(*) FROM user WHERE status = 'active'

	// You can also use "?" as variable then use g.AddValues(...) to set variable's value.
	var g gooq.Gooq
	g.SQL.Insert("user", "?", "?", "?").Values("my_account", 123456, "user_name")
	g.AddValues("my_account", 123456, "user_name")

	db, _ := _sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/dbName")
	tx, _ := db.Begin()
	tx.Exec(g.SQL.GetSQL(), g.Args...)
}
```
  
#### You can also use "?" as variable then use g.AddValues(args...) to set variable's value.
```go
import (
	"database/sql"

	"github.com/EricChiou/gooq"
)

func main() {
	db, _ := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/dbName")
	tx, _ := db.Begin()

	var g gooq.Gooq
	g.SQL.Insert("user", "acc", "pwd", "name").Values("?", "?", "?")
	g.AddValues("my_account", 123456, "user_name")

	tx.Exec(g.SQL.GetSQL(), g.Args...)
}
```