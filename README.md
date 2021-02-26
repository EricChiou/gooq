## gooq
Simply SQL ORM. Inspired by [JOOQ](https://www.jooq.org/).  

## How to use
#### Simple Example
``` go
import (
	"fmt"

	"github.com/EricChiou/gooq"
)

func main() {
	// INSERT
	var sql gooq.SQL
	sql.Insert("user", "acc", "pwd", "name").Values("my_account", 123456, "user_name")

	fmt.Println(sql.GetSQL())
	// sql: INSERT INTO user ( acc, pwd, name ) VALUES ( 'my_account', 123456, 'user_name' )

	// UPDATE
	sql = gooq.SQL{}
	sql.Update("user").Set("pwd").Eq(111111).C("name").Eq("user_name2").Where("id").Eq(1)

	fmt.Println(sql.GetSQL())
	// sql: UPDATE user SET pwd=111111, name='user_name2' WHERE id=1

	// DELETE
	sql = gooq.SQL{}
	sql.Delete("user").Where("id").Gt(1)

	fmt.Println(sql.GetSQL())
	// sql: DELETE FROM user WHERE id>1

	// Query
	sql = gooq.SQL{}
	sql.Select("*").From("user").Where("id").Lt(100)

	fmt.Println(sql.GetSQL())
	// sql: SELECT * FROM user WHERE id<100

	sql = gooq.SQL{}
	sql.Select("acc", "pwd", "name").From("user").Where("name").Like("%Eric%")

	fmt.Println(sql.GetSQL())
	// sql: SELECT acc, pwd, name FROM user WHERE name LIKE '%Eric%'

	// INNER JOIN, Sub Query, ORDER BY, LIMIT (mysql), USING
	sql = gooq.SQL{}
	var subSQL gooq.SQL
	sql.Select("id", "acc", "name").From("user")
	sql.Join(subSQL.Lp().Select("id").From("user").OrderBy("id").Limit("0", "10").Rp().As("t").GetSQL())
	sql.Using("id").Where("status").Eq("active")

	fmt.Println(sql.GetSQL())
	// sql: SELECT id, acc, name FROM user INNER JOIN ( SELECT id FROM user ORDER BY id LIMIT 0, 10 ) t USING (id) WHERE status='active'

	// COUNT
	sql = gooq.SQL{}
	sql.Select().Count("*").From("user").Where("status").Eq("active")

	fmt.Println(sql.GetSQL())
	// sql: SELECT COUNT(*) FROM user WHERE status='active'
}
```
  
#### You can also use "?" as variable and use g.AddValues(args...) to set variable's value.
```go
import (
	"database/sql"
	"fmt"

	"github.com/EricChiou/gooq"
)

func main() {
	db, _ := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/dbName")
	tx, _ := db.Begin()

	var g gooq.Gooq
	g.SQL.Insert("user", "acc", "pwd", "name").Values("?", "?", "?")
	g.AddValues("my_account", 123456, "user_name")

	fmt.Println(g.SQL.GetSQL())
	fmt.Println(g.Args)
	// sql: INSERT INTO user ( acc, pwd, name ) VALUES ( ?, ?, ? )
	// arg: [my_account 123456 user_name]

	tx.Exec(g.SQL.GetSQL(), g.Args...)
}
```