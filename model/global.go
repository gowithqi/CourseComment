package model

import (
	"database/sql"
	"github.com/CourseComment/conf"
	_ "github.com/go-sql-driver/mysql"
	//"os"
	//"time"
)

type idtype int32

var (
	db *sql.DB
)

func init() {
	db = conf.DB
}
