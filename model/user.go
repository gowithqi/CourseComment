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

//time.Time need to be understand
type User struct {
	Id       idtype
	Account  string
	Name     string
	Password string
	//Join_time time.Time
}

func GetUser(key string, value interface{}) *User {
	var rows *sql.Rows
	res := new(User)

	switch key {
	case "Id":
		id := value.(idtype)
		//extract from the database
		rows, _ = db.Query("select id, account, name, password from user where id=?", id)
	case "Account":
		account := value.(string)
		//extract from the database
		rows, _ = db.Query("select id, account, name, password from user where account=?", account)
	case "Name":
		name := value.(string)
		//extract from the database
		rows, _ = db.Query("select id, account, name, password from user where name=?", name)
	}
	if rows.Next() {
		rows.Scan(&res.Id, &res.Account, &res.Name, &res.Password)
	}

	return res
}

func (user User) RegisterCheck() (bool, string) {
	var (
		rows  *sql.Rows
		count int
	)
	rows, _ = db.Query("select count(*) from user, registerUser where (user.name=?) or (user.name=?)",
		user.Name, user.Name)
	rows.Scan(&count)
	if count > 0 {
		return false, "name"
	}

	rows, _ = db.Query("select count(*) from user, registerUser where (user.account=?) or (user.account=?)",
		user.Account, user.Account)
	rows.Scan(&count)
	if count > 0 {
		return false, "account"
	}
	return true, "success"
}

func (user User) Register() int64 {
	stmt, _ := db.Prepare("insert registeringUser set name=?, account=?, password=?, status=?")
	res, _ := stmt.Exec(user.Name,
		user.Account,
		user.Password,
		"registering")
	resId, _ := res.LastInsertId()
	return resId
}

func RegisterSuccess(id idtype) {
	var (
		name     string
		account  string
		password string
	)
	rows, _ := db.Query("select name, account, password from registeringUser where id=?",
		id)
	rows.Scan(&name, &account, &password)

	stmt, _ := db.Prepare("insert user set name=?, account=?, password=?")
	stmt.Exec(name, account, password)

	go registerSuccess(id)
}

func registerSuccess(id idtype) {
	stmt, _ := db.Prepare("delete from registeringUser where id=?")
	stmt.Exec(id)
}
