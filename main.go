package main

import (
	"database/sql"
	"fmt"
	"os"
)

const cli_HELLO = "Hello :)"
const cli_NOT_FOUND_ERROR = "Connot undrestand command. \n -h for help"
const cli_HELP = "-h for help"
const cli_ENTER_A_VALID_COMMAND = "please enter a valid command."
const cli_SUCSESSFUL_ADDED_TODO = "ToDo Added to List"
const cli_SUCSESSFUL_REMOVED_TODO = "ToDo Removed from List"
const cli_NOTHING_TO_DO = "nothing to do! \nStart with: add 'todoTitle'"
const help = `-h
	add [title]  add todo to todo list
	remove [id]  remove todo from todo list
	ls , l       todo list`

func main() {
	arg := os.Args
	if len(arg) < 2 {
		fmt.Println(cli_HELLO)
		fmt.Println(cli_HELP)
		return
	}
	db, _ := sql.Open("sqlite3", "F:\\tdp\\database\\tdp.db")
	tdp_intro(arg[1:], db)
}
func tdp_intro(args []string, db *sql.DB) {
	db.Exec(`create table if not exists toDo 
	(id integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		title text, 
		create_time datetime
	)`)
	if len(args) == 2 {
		flag := args[0]
		if flag == "add" {
			addTodo(db, args[1])
			fmt.Println(cli_SUCSESSFUL_ADDED_TODO)
		} else if flag == "del" {
			deleteTodo(db, args[1])
			fmt.Println(cli_SUCSESSFUL_REMOVED_TODO)
		} else {
			fmt.Println(cli_NOT_FOUND_ERROR)
			return
		}
	} else if len(args) == 1 {
		flag := args[0]
		if flag == "ls" || flag == "l" {
			toDos := getTodos(db)
			showTable(toDos)
		} else if flag == "-h" {
			fmt.Println(help)
		}
	} else {
		fmt.Println(cli_NOT_FOUND_ERROR)
		return
	}
}
