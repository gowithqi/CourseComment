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

type Lecture struct {
	Id idtype
	Course
	Professor
	Student_score        float32
	Level                float32
	Student_score_number int32
	Level_number         int32

	Comments []LectureComment
}

func GetLecture(c *Course, p *Professor) *Lecture {
	var rows *sql.Rows
	res := &Lecture{Course: *c, Professor: *p}

	rows, _ = db.Query("select id, student_score, level, student_score_number, level_number from lecture where course_id=? and professor_id=?",
		c.Id, p.Id)
	if !rows.Next() {
		return nil
	}

	rows.Scan(&res.Id, &res.Student_score, &res.Level, &res.Student_score_number, &res.Level_number)

	return res
}

func (l Lecture) updateDB() {
	stmt, _ := db.Prepare("update lecture set student_score=?, level=?, student_score_number=?, level_number=? where id=?")
	stmt.Exec(l.Student_score, l.Level, l.Student_score_number, l.Level_number, l.Id)
}

func (l *Lecture) RecordStudentScore(u User, score int8) {
	l.Student_score = (l.Student_score*float32(l.Student_score_number) + float32(score)) / float32(l.Student_score_number+1)
	l.Student_score_number++

	stmt, _ := db.Prepare("insert lectureStudentScoreRecord set user_id=?, lecture_id=?, socre=?")
	stmt.Exec(u.Id, l.Id, score)

	l.updateDB()
}

func (l *Lecture) RecordLevel(u User, level int8) {
	l.Level = (l.Level*float32(l.Level_number) + float32(level)) / float32(l.Level_number+1)
	l.Level_number++

	stmt, _ := db.Prepare("insert lectureLevelRecord set user_id=?, lecture_id=?, level=?")
	stmt.Exec(u.Id, l.Id, level)

	l.updateDB()
}

func (l *Lecture) GetComments() {
	l.Comments = make([]LectureComment, 0)

	//TODO
	//need a time
	rows, _ := db.Query("select id, user_id, content, super_number from lectureComment where lecture_id=?", l.Id)

	for rows.Next() {
		var (
			id           idtype
			user_id      idtype
			content      string
			super_number int32
			tmp          LectureComment
		)
		rows.Scan(&id, &user_id, &content, &super_number)
		tmp = LectureComment{Id: id,
			Lecture:      *l,
			User:         *GetUser("id", user_id),
			Content:      content,
			Super_number: super_number}
		l.Comments = append(l.Comments, tmp)
	}
}
