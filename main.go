package main

import (
	"gopkg.in/alecthomas/kingpin.v2"

	"fmt"
)

var (
	add       = kingpin.Command("add", "Add something new.")
	addEntity = add.Arg("what", "entity type (task, note or secret)").Required().String()
	addText   = add.Arg("text", "Description").Required().String()
)

func main() {
	switch kingpin.Parse() {

	case "add":
		fmt.Printf("Going to add this %s: %s", *addEntity, *addText)
	}
}
