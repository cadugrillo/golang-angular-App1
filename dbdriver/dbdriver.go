package dbdriver

import (
	"database/sql"
	"log"
	"os"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var (
	once sync.Once
)

func init() {
	once.Do(initialiseList)
}

func initialiseList() {

}

// Todo data structure for a task with a description of what to do
type Todo struct {
	ID       string `json:"id"`
	Message  string `json:"message"`
	Complete bool   `json:"complete"`
}

func dbConn(dbname string) (db *sql.DB) {
	dbIpaddr := os.Getenv("DBIPADDRESS")
	dbPort := os.Getenv("DBPORT")
	dbDriver := os.Getenv("DBDRIVER")
	dbUser := os.Getenv("DBUSER")
	dbPass := os.Getenv("DBPASSWORD")
	dbName := dbname
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbIpaddr+":"+dbPort+")/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	log.Println("SQL connection opened successfully")
	return db

}

func InitialiseDB() bool {

	dbDriver := os.Getenv("DBDRIVER")
	if dbDriver == "" {
		log.Println("No DB Driver was found. System will use Default storage method")
		return false
	}

	myDBinit := dbConn("")

	AddDB, err := myDBinit.Prepare("CREATE DATABASE IF NOT EXISTS todo")
	if err != nil {
		panic(err.Error())
	}
	AddDB.Exec()
	log.Println("Database todo created successfully")
	myDBinit.Close()

	myDB := dbConn("todo")
	addTable, err := myDB.Prepare("CREATE TABLE IF NOT EXISTS todo_items (ID varchar(36), Message varchar(255), Complete boolean)")
	if err != nil {
		panic(err.Error())
	}
	addTable.Exec()
	log.Println("Table todo_items created successfully")

	myDB.Close()

	return true
}

func DatabaseGet() []Todo {
	var todos []Todo
	db := dbConn("todo")

	rows, err := db.Query("SELECT * FROM todo_items")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	db.Close()

	for rows.Next() {
		var todo Todo
		rows.Scan(&todo.ID, &todo.Message, &todo.Complete)
		todos = append(todos, todo)
	}
	return todos

}

func DatabaseAdd(ID string, Message string, Complete bool) {

	db := dbConn("todo")

	r, err := db.Prepare("INSERT INTO todo_items(ID, Message, Complete) VALUES(?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}
	r.Exec(ID, Message, Complete)
	db.Close()
	log.Println("New item created successfully")
}

func DatabaseComplete(ID string) {

	db := dbConn("todo")

	r, err := db.Prepare("UPDATE todo_items SET Complete=true WHERE ID=?")
	if err != nil {
		panic(err.Error())
	}
	r.Exec(ID)
	db.Close()
	log.Println("Item" + ID + "status completed set successfully")
}

func DatabaseDelete(ID string) {

	db := dbConn("todo")

	r, err := db.Prepare("DELETE FROM todo_items WHERE ID=?")
	if err != nil {
		panic(err.Error())
	}
	r.Exec(ID)
	db.Close()
	log.Println("Item " + ID + " removed successfully")
}
