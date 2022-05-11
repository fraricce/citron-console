package main

import (
	"database/sql"
	"fmt"

	"github.com/alexeyco/simpletable"
	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/alecthomas/kingpin.v2"
)

type entity struct {
	text string
	done bool
}

var tasks []entity

var (
	add        = kingpin.Command("add", "Add something new.")
	addEntity  = add.Arg("what", "entity type (task, note or secret)").Required().String()
	addText    = add.Arg("text", "Description").Required().String()
	list       = kingpin.Command("list", "List")
	listEntity = list.Arg("whatToList", "entity type (task, note or secret)").Required().String()
	data       = [][]interface{}{}
)

func main() {

	fmt.Println(`
	_______ __                      ______                       __   
       / ____(_) /__________  ____     / ____/___  ____  _________  / /__ 
      / /   / / __/ ___/ __ \/ __ \   / /   / __ \/ __ \/ ___/ __ \/ / _ \
     / /___/ / /_/ /  / /_/ / / / /  / /___/ /_/ / / / (__  ) /_/ / /  __/
     \____/_/\__/_/   \____/_/ /_/   \____/\____/_/ /_/____/\____/_/\___/ 
                                                                    	
	`)

	database, _ := sql.Open("sqlite3", "./citron.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS tasks (id INTEGER PRIMARY KEY, title TEXT, done INTEGER)")
	statement.Exec()

	switch kingpin.Parse() {

	case "list":
		rows, _ := database.Query("SELECT * FROM tasks")
		var id int
		var title string
		var done int

		for rows.Next() {
			rows.Scan(&id, &title, &done)
			newVector := []interface{}{id, title, done}
			data = append(data, newVector)
		}

		// data       = [][]interface{}{
		// 	{1, "Newton G. Goetz", 532.7},
		// }

		table := simpletable.New()

		table.Header = &simpletable.Header{
			Cells: []*simpletable.Cell{
				{Align: simpletable.AlignCenter, Text: "#"},
				{Align: simpletable.AlignCenter, Text: "Task"},
				{Align: simpletable.AlignCenter, Text: "Done"},
			},
		}

		for _, row := range data {
			r := []*simpletable.Cell{
				{Align: simpletable.AlignRight, Text: fmt.Sprintf("%d", row[0].(int))},
				{Text: row[1].(string)},
				{Align: simpletable.AlignRight, Text: fmt.Sprintf("%d", row[2].(int))},
			}

			table.Body.Cells = append(table.Body.Cells, r)
		}

		table.SetStyle(simpletable.StyleCompactLite)
		fmt.Println(table.String())

		break
	case "add":
		fmt.Printf("Going to add this %s: %s", *addEntity, *addText)
		task := entity{text: *addText, done: false}
		tasks = append(tasks, task)
		fmt.Printf("len=%d cap=%d %v\n", len(tasks), cap(tasks), tasks)
		stmt, _ := database.Prepare("INSERT INTO tasks(title, done) values(?,?)")
		stmt.Exec(task.text, 0)
		break

	}

	database.Close()
}
