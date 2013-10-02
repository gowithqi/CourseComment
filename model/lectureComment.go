package model

import (
	"database/sql"
	//"github.com/CourseComment/conf"
	_ "github.com/go-sql-driver/mysql"
	//"os"
	//"time"
)

type LectureComment struct {
	Id idtype
	Lecture
	User
	Content      string
	Super_number int32
	//Time time.Time
}

func (lc LectureComment) updateDB() {
	var rows *sql.Rows
	_ = rows
	stmt, _ := db.Prepare("update lectureComment set super_number=? where id=?")
	stmt.Exec(lc.Super_number, lc.Id)
}

func (lc *LectureComment) RecordLectureCommentSuper(u User) {
	lc.Super_number++

	stmt, _ := db.Prepare("insert lectureCommentSuperRecord set lecture_id=?, user_id=?")
	stmt.Exec(lc.Id, u.Id)

	lc.updateDB()
}

func (lc LectureComment) AddComment() {
	stmt, _ := db.Prepare("insert lectureComment set lecture_id=?, user_id=?, content=?, super_number=?")
	stmt.Exec(lc.Lecture.Id, lc.User.Id, lc.Content, lc.Super_number)
}
