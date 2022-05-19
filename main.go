package main

import (
	"fmt"
	"log"

	"github.com/alexeyco/simpletable"
	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	add            = kingpin.Command("add", "Add something new.")
	addEntity      = add.Arg("what", "entity type (task, note or secret)").Required().String()
	addText        = add.Arg("text", "Description").Required().String()
	list           = kingpin.Command("list", "List")
	listEntity     = list.Arg("whatToList", "entity type (task, note or secret)").Required().String()
	done           = kingpin.Command("done", "Done")
	doneEntityId   = done.Arg("whatIdToSetAsDone", "entity id").Required().String()
	undone         = kingpin.Command("undone", "Undone")
	undoneEntityId = undone.Arg("whatIdToSetAsUndone", "entity id").Required().String()
	delete         = kingpin.Command("del", "Delete")
	deleteEntityId = delete.Arg("whatIdToDel", "entity id").Required().String()
	data           = [][]interface{}{}
)

func main() {

	InitDatabase()

	fmt.Println(`
	  ______ __                  _____                   __
	 / ___(_) /________  ___    / ___/__  ___  ___ ___  / /__ 
	/ /__/ / __/ __/ _ \/ _ \  / /__/ _ \/ _ \(_-</ _ \/ / -_)
	\___/_/\__/_/  \___/_//_/  \___/\___/_//_/___/\___/_/\__/ 
	`)

	switch kingpin.Parse() {

	case "list":

		data = ListEntities(*listEntity)

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

		res, eName := AddEntity(*addEntity, *addText)
		if res {
			fmt.Printf("-> Entity %s has been created\n", eName)
		} else {
			fmt.Printf("-> Entity %s could not be created\n", eName)
		}
		break

	case "del":
		count := RemoveEntity(*deleteEntityId)
		if count == 0 {
			fmt.Printf("-> Cannot find an entity with id %d.\n", deleteEntityId)
		} else {
			fmt.Printf("-> Entity id %d has been deleted.\n", deleteEntityId)
		}

		break

	case "done":
		count := SetEntityStatus(*doneEntityId, 1)
		if count == 0 {
			fmt.Printf("-> Cannot find an entity with id %s.\n", *doneEntityId)
		} else {
			fmt.Printf("-> Entity id %s has been set to done.\n", *doneEntityId)
		}
		break

	case "undone":
		count := SetEntityStatus(*undoneEntityId, 0)
		if count == 0 {
			fmt.Printf("-> Cannot find an entity with id %s.\n", *undoneEntityId)
		} else {
			fmt.Printf("-> Entity id %s has been set to undone.\n", *undoneEntityId)
		}

		break

	default:
		log.Print("None")
		break

	}

	fmt.Println("")

}
