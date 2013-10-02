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
	Id     idtype
	Name   string
	Number string

	Professors []Professor
}

func GetCourse(key string, value interface{}) *Course {
	var rows *sql.Rows
	res := new(Course)

	switch key {
	case "Id":
		id := value.(idtype)
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

func (c *Course) GetProfessorOfThisCourse() {
	c.Professors = make([]Professor, 0)

	rows, _ := db.Query("select professor_id from lecture where cource_id=?", c.Id)
	for rows.Next() {
		var id idtype
		rows.Scan(&id)
		tmp := GetProfessor("Id", id)
		c.Professors = append(c.Professors, *tmp)
	}
}
