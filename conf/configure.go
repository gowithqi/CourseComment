package conf

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	DB *sql.DB
)

func init() {
	DB, _ = sql.Open("mysql", "zzq:zzq_sjtu@tcp(localhost:3306)/CourseComment")
}
