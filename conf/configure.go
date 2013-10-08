package conf

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

var (
	DB *sql.DB
)

//trade-off sorting parameter
const (
	yearseconds int    = (365 * 24 * 3600)
	Timepara    int    = yearseconds / 20
	TimeLayout  string = "2006-01-02 15:04:05"
)

var (
	TimeStart time.Time
)

func init() {
	DB, _ = sql.Open("mysql", "zzq:zzq_sjtu@tcp(localhost:3306)/CourseComment")

	//trade-off parameter init
	TimeStart, _ = time.Parse(TimeLayout, "2013-10-08 22:28:30")
}
