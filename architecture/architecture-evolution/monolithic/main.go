package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type student struct {
	id    string
	sname string
	class string
}

type grade struct {
	id     int
	sid    string
	sname  string
	course int
	score  int
}

const (
	mysqlAddNewStudent = iota
	mysqlAddNewGrade
	mysqlDeleteStudent
	mysqlDeleteGradeBySIDAndCourse
	mysqlModifyScoreBySIDAndCourse
	mysqlQueryGradeBySID
)

var sqlString = []string{
	`INSERT INTO students(id,sname,class) VALUES(?, ?, ?);`,
	`INSERT INTO grade(sid,course,score)VALUES(?, ?, ?);`,
	`DELETE FROM students WHERE id = ?;`,
	`DELETE FROM grade WHERE sid = ? AND course = ?;`,
	`UPDATE grade SET score = ? WHERE sid = ? AND course = ?;`,
	`SELECT sid, sname, course, score FROM grade LEFT JOIN students ON  grade.sid = students.id WHERE sid = ?;`,
}

func main() {
	db, err := sql.Open("mysql", "root:Miufighting.@tcp(127.0.0.1:3306)/demo")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	for {
		var identifier string

		fmt.Println("add new student: 1    add new grade:2")
		fmt.Println("delete student:  3    delete grade:4")
		fmt.Println("modify grade:    5    query grade:6")
		fmt.Println()

		fmt.Printf("please choose:")
		fmt.Scanln(&identifier)

		switch {
		case identifier == "1":
			addNewStudent(db)
		case identifier == "2":
			addNewGradeBySID(db)
		case identifier == "3":
			deleteStudent(db)
		case identifier == "4":
			deleteGradeBySIDAndCourse(db)
		case identifier == "5":
			modifyScoreBySIDAndCourse(db)
		case identifier == "6":
			queryGradeBySID(db)
		case identifier == "end":
			fmt.Println("goodbye!")
			return
		}

	}

}

func addNewStudent(db *sql.DB) {
	var s student

	fmt.Printf("please entry student number:  ")
	fmt.Scanln(&s.id)

	fmt.Printf("please entry student name:  ")
	fmt.Scanln(&s.sname)

	fmt.Printf("please entry student class:  ")
	fmt.Scanln(&s.class)

	db.Exec(sqlString[mysqlAddNewStudent], s.id, s.sname, s.class)

	fmt.Println("add new student success!")
	fmt.Println()
}

func addNewGradeBySID(db *sql.DB) {
	var g grade

	fmt.Printf("please entry student number:  ")
	fmt.Scanln(&g.sid)

	fmt.Printf("please entry student course:  ")
	fmt.Scanln(&g.course)

	fmt.Printf("please entry student score:  ")
	fmt.Scanln(&g.score)

	db.Exec(sqlString[mysqlAddNewGrade], g.sid, g.course, g.score)

	fmt.Println("add student grade success!")
	fmt.Println()
}

func queryGradeBySID(db *sql.DB) {
	var (
		studentGrades []*grade
		sid           string
	)

	fmt.Printf("please entry student number:  ")
	fmt.Scanln(&sid)

	rows, _ := db.Query(sqlString[mysqlQueryGradeBySID], sid)

	for rows.Next() {
		var g grade
		rows.Scan(&g.sid, &g.sname, &g.course, &g.score)
		studentGrades = append(studentGrades, &g)
	}

	for k := range studentGrades {
		fmt.Println(studentGrades[k].sid, studentGrades[k].sname, studentGrades[k].course, studentGrades[k].score)
	}

	fmt.Println()
}

func deleteStudent(db *sql.DB) {
	var sid string

	fmt.Printf("please entry student number:  ")
	fmt.Scanln(&sid)

	db.Exec(sqlString[mysqlDeleteStudent], sid)

	fmt.Println("delete student success!")
	fmt.Println()
}

func deleteGradeBySIDAndCourse(db *sql.DB) {
	var g grade

	fmt.Printf("please entry student number:  ")
	fmt.Scanln(&g.sid)

	fmt.Printf("please entry student course:  ")
	fmt.Scanln(&g.course)

	db.Exec(sqlString[mysqlDeleteGradeBySIDAndCourse], g.sid, g.course)

	fmt.Println("delete student grade success!")
	fmt.Println()
}

func modifyScoreBySIDAndCourse(db *sql.DB) {
	var g grade

	fmt.Printf("please entry student number:  ")
	fmt.Scanln(&g.sid)

	fmt.Printf("please entry student course:  ")
	fmt.Scanln(&g.course)

	fmt.Printf("please entry student score:  ")
	fmt.Scanln(&g.score)

	db.Exec(sqlString[mysqlModifyScoreBySIDAndCourse], g.score, g.sid, g.course)

	fmt.Println("update student grade success!")
	fmt.Println()
}
