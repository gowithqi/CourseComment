package model

import (
	"database/sql"
	//"github.com/CourseComment/conf"
	_ "github.com/go-sql-driver/mysql"
	//"os"
	//"time"
)

// var (
// 	db *sql.DB
// )

// func init() {
// 	db = conf.DB
// }

type Course struct {
	Id     int
	Name   string
	Number string
}

func GetCourse(key string, value interface{}) *Course {
	var rows *sql.Rows
	res := new(Course)

	switch key {
	case "Id":
		id := value.(int)
		//extract from the database
		rows, _ = db.Query("select id, name, number from course where id=?", id)
	case "Name":
		name := value.(string)
		//extract from the database
		rows, _ = db.Query("select id, name, number from course where name=?", name)
	case "Number":
		number := value.(string)
		//extract from the database
		rows, _ = db.Query("select id, name, number from course where number=?", number)
	}
	if rows.Next() {
		rows.Scan(&res.Id, &res.Name, &res.Number)
	}

	return res
}
