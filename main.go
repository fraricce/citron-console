package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"github.com/alexeyco/simpletable"
	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/alecthomas/kingpin.v2"
)

type entity struct {
	text string
	done bool
}

var (
	add            = kingpin.Command("add", "Add something new.")
	addEntity      = add.Arg("what", "entity type (task, note or secret)").Required().String()
	addText        = add.Arg("text", "Description").Required().String()
	list           = kingpin.Command("list", "List")
	listEntity     = list.Arg("whatToList", "entity type (task, note or secret)").Required().String()
	done           = kingpin.Command("done", "Done")
	doneEntityId   = done.Arg("whatIdToSetAsDone", "entity id").Required().String()
	undone         = kingpin.Command("undone", "Undone")
	undoneEntityId = undone.Arg("whatIdToSetAsDone", "entity id").Required().String()
	delete         = kingpin.Command("del", "Delete")
	deleteEntityId = delete.Arg("whatIdToDel", "entity id").Required().String()
	data           = [][]interface{}{}
)

func main() {

	fmt.Println(`
	  ______ __                  _____                   __
	 / ___(_) /________  ___    / ___/__  ___  ___ ___  / /__ 
	/ /__/ / __/ __/ _ \/ _ \  / /__/ _ \/ _ \(_-</ _ \/ / -_)
	\___/_/\__/_/  \___/_//_/  \___/\___/_//_/___/\___/_/\__/ 
	`)

	database, _ := sql.Open("sqlite3", "./citron.db")
	// entity is 0 for task, 1 for note
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS entities (id INTEGER PRIMARY KEY, title TEXT, done INTEGER, entity INTEGER)")
	statement.Exec()

	switch kingpin.Parse() {

	case "list":

		var eCode int

		switch *listEntity {
		case "task":
			eCode = 0
			break
		case "note":
			eCode = 1
			break
		}

		rows, _ := database.Query("SELECT id, title, done FROM entities where entity = ?", eCode)
		var id int
		var title string
		var done int

		for rows.Next() {
			rows.Scan(&id, &title, &done)
			newVector := []interface{}{id, title, done}
			data = append(data, newVector)
		}

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
		var eCode int
		var eName string

		switch *addEntity {
		case "task":
			eCode = 0
			eName = "task"
			break
		case "note":
			eCode = 1
			eName = "string"
			break
		}

		task := entity{text: *addText, done: false}
		stmt, _ := database.Prepare("INSERT INTO entities(title, done, entity) values(?,?,?)")
		_, e := stmt.Exec(task.text, 0, eCode)
		if e != nil {
			log.Printf("Error %s", e)
		}
		fmt.Printf("-> %s has been added.\n", eName)
		break

	case "del":
		eId, _ := strconv.Atoi(*deleteEntityId)
		stmt, _ := database.Prepare("DELETE FROM entities WHERE id = ?")
		a, e := stmt.Exec(eId)
		if e != nil {
			log.Printf("Error %s", e)
		}
		count, _ := a.RowsAffected()
		if count == 0 {
			fmt.Printf("-> Cannot find an entity with id %d.\n", eId)
		} else {
			fmt.Printf("-> Entity id %d has been deleted.\n", eId)
		}

		break

	case "done":
		eId, _ := strconv.Atoi(*doneEntityId)
		stmt, _ := database.Prepare("UPDATE entities SET done = 1 WHERE id = ?")
		a, e := stmt.Exec(eId)
		if e != nil {
			log.Printf("Error %s", e)
		}
		count, _ := a.RowsAffected()
		if count == 0 {
			fmt.Printf("-> Cannot find an entity with id %d.\n", eId)
		} else {
			fmt.Printf("-> Entity id %d has been set to done.\n", eId)
		}

		break

	case "undone":
		eId, _ := strconv.Atoi(*undoneEntityId)
		stmt, _ := database.Prepare("UPDATE entities SET done = 0 WHERE id = ?")
		a, e := stmt.Exec(eId)
		if e != nil {
			log.Printf("Error %s", e)
		}
		count, _ := a.RowsAffected()
		if count == 0 {
			fmt.Printf("-> Cannot find an entity with id %d.\n", eId)
		} else {
			fmt.Printf("-> Entity id %d has been set to undone.\n", eId)
		}

		break

	}

	fmt.Printf("\n\n")
	database.Close()
}
