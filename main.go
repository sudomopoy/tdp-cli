package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/fatih/structs"
)

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
		if contains(cm.ADD_TODO_ITEM, flag) {
			addTodo(db, args[1])
			fmt.Println(cli_SUCSESSFUL_ADDED_TODO)
		} else if contains(cm.DELETE_TODO_ITEM, flag) {
			deleteTodo(db, args[1])
			fmt.Println(cli_SUCSESSFUL_REMOVED_TODO)
		} else {
			fmt.Println(cli_NOT_FOUND_ERROR)
			return
		}
	} else if len(args) == 1 {
		flag := args[0]
		if contains(cm.TODO_LIST_ITEMS, flag) {
			toDos := getTodos(db)
			showTable(toDos)
		} else if contains(cm.DELETE_ALL_TODOS_ITEMS, flag) {
			var answer string
			fmt.Println("drop all to dos, sure? (Y or N)")
			fmt.Print(":")
			fmt.Scan(&answer)
			if contains(cm.SURE_COMMAND, answer) {
				deleteAllTodo(db)
			}
		} else if contains(cm.HELP_COMMAND, flag) {
			for _, v := range structs.Values(cm) {
				fmt.Println(v)
			}
		}
	} else {
		fmt.Println(cli_NOT_FOUND_ERROR)
		return
	}
}
