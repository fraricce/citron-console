package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var database *sql.DB

func InitDatabase() {
	database, _ := sql.Open("sqlite3", "./citron.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS entities (id INTEGER PRIMARY KEY, title TEXT, done INTEGER, entity INTEGER)")
	statement.Exec()
	database.Close()
}

func CreateEntity(item entity) bool {
	database, _ := sql.Open("sqlite3", "./citron.db")

	stmt, _ := database.Prepare("INSERT INTO entities(title, done, entity) values(?,?,?)")
	_, e := stmt.Exec(item.text, 0, item.entity)
	if e != nil {
		log.Printf("Error %s", e)
		return false
	}

	defer database.Close()
	return true
}

func FindAllEntities(eCode int) [][]interface{} {
	database, _ := sql.Open("sqlite3", "./citron.db")

	rows, _ := database.Query("SELECT id, title, done FROM entities where entity = ?", eCode)
	var id int
	var title string
	var done int

	for rows.Next() {
		rows.Scan(&id, &title, &done)
		newVector := []interface{}{id, title, done}
		data = append(data, newVector)
	}

	defer database.Close()
	return data
}

func DeleteEntity(id int) int64 {
	database, _ := sql.Open("sqlite3", "./citron.db")

	stmt, _ := database.Prepare("DELETE FROM entities WHERE id = ?")
	a, e := stmt.Exec(id)
	if e != nil {
		log.Printf("Error %s", e)
	}
	count, _ := a.RowsAffected()

	defer database.Close()
	return count
}

func UpdateEntityStatus(id int, status int) int64 {
	database, _ := sql.Open("sqlite3", "./citron.db")
	stmt, _ := database.Prepare("UPDATE entities SET done = ? WHERE id = ?")
	a, e := stmt.Exec(status, id)
	if e != nil {
		log.Printf("Error %s", e)
	}
	defer database.Close()
	count, _ := a.RowsAffected()
	return count
}
