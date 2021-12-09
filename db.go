package main

import (
	"database/sql"
	"log"
	"strconv"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func addTodo(db *sql.DB, title string) {
	tx, _ := db.Begin()
	current_time := time.Now()
	stmt, _ := tx.Prepare("insert into toDo (title,create_time) values (?,?)")
	_, err := stmt.Exec(title, current_time)
	if err != nil {
		log.Fatal(err)
	}
	tx.Commit()
}

func getTodos(db *sql.DB) [][]string {
	rows, err := db.Query("select * from toDo")
	if err != nil {
		log.Fatal(err)
	}
	todos := [][]string{}
	for rows.Next() {
		var tempTodo toDo_model
		err =
			rows.Scan(&tempTodo.id, &tempTodo.title, &tempTodo.created_at)
		if err != nil {
			log.Fatal(err)
		}

		todos = append(todos, []string{strconv.Itoa(tempTodo.id), tempTodo.title, tempTodo.created_at.Format("2 Jan 2006 15:04")})
	}

	return todos
}
func deleteTodo(db *sql.DB, id string) {
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("DELETE FROM toDo WHERE id=?")
	_, err := stmt.Exec(id)
	if err != nil {
		log.Fatal(err)
	}
	tx.Commit()
}
func deleteAllTodo(db *sql.DB) {
	tx, _ := db.Begin()
	_, err := tx.Exec("DELETE FROM toDo;")
	if err != nil {
		log.Fatal(err)
	}
	tx.Commit()
}
