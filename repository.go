package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Repo struct {
	database *sql.DB
	dbdriver string
	dbpath   string
}

func (r Repo) InitDatabase() {
	database, _ := sql.Open(r.dbdriver, r.dbpath)
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS entities (id INTEGER PRIMARY KEY, title TEXT, done INTEGER, entity INTEGER)")
	statement.Exec()
	database.Close()
}

func (r Repo) CreateEntity(item entity) bool {
	stmt, _ := r.database.Prepare("INSERT INTO entities(title, done, entity) values(?,?,?)")
	_, e := stmt.Exec(item.text, 0, item.entity)
	if e != nil {
		log.Printf("Error %s", e)
		return false
	}

	defer r.database.Close()
	return true
}

func (r Repo) FindAllEntities(eCode int) [][]interface{} {
	rows, _ := r.database.Query("SELECT id, title, done FROM entities where entity = ?", eCode)
	var id int
	var title string
	var done int

	for rows.Next() {
		rows.Scan(&id, &title, &done)
		newVector := []interface{}{id, title, done}
		data = append(data, newVector)
	}

	defer r.database.Close()
	return data
}

func (r Repo) DeleteEntity(id int) int64 {
	stmt, _ := r.database.Prepare("DELETE FROM entities WHERE id = ?")
	a, e := stmt.Exec(id)
	if e != nil {
		log.Printf("Error %s", e)
	}
	count, _ := a.RowsAffected()

	defer r.database.Close()
	return count
}

func (r Repo) UpdateEntityStatus(id int, status int) int64 {
	stmt, _ := r.database.Prepare("UPDATE entities SET done = ? WHERE id = ?")
	a, e := stmt.Exec(status, id)
	if e != nil {
		log.Printf("Error %s", e)
	}
	defer r.database.Close()
	count, _ := a.RowsAffected()
	return count
}
