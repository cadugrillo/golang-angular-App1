package dbdriver

import (
	"database/sql"
	"log"
	"os"
	"strings"
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

func dbConn() (db *sql.DB) {
	dbIpaddr := os.Getenv("DBIPADDRESS")
	dbPort := os.Getenv("DBPORT")
	dbDriver := os.Getenv("DBDRIVER")
	dbUser := os.Getenv("DBUSER")
	dbPass := os.Getenv("DBPASSWORD")
	dbName := "todo"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbIpaddr+":"+dbPort+")/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	log.Println("SQL connection opened successfully")
	return db

}

func InitialiseDB() bool {

	//dbDriver := os.Getenv("DBDRIVER")
	//if dbDriver == "" {
	//	log.Println("No DB Driver was found. System will use Default storage method")
	//	return false
	//}

	//myDBinit := dbConn("")

	//AddDB, err := myDBinit.Prepare("CREATE DATABASE IF NOT EXISTS todo")
	//if err != nil {
	//	panic(err.Error())
	//}
	//AddDB.Exec()
	//log.Println("Database todo created successfully")
	//myDBinit.Close()

	//mail := "cadugrillo@gmail.com"
	//userIdModint := strings.Replace(mail, "@", "_", 1)
	//userIdMod := strings.Replace(userIdModint, ".", "_", 1)

	//myDB := dbConn("todo")
	//addTable, err := myDB.Prepare("CREATE TABLE IF NOT EXISTS " + userIdMod + " (ID varchar(36), Message varchar(255), Complete boolean)")
	//if err != nil {
	//	panic(err.Error())
	//}
	//addTable.Exec()
	//log.Println("Table todo_items created successfully")

	//myDB.Close()

	return false
}

func DatabaseGet(userId string) []Todo {
	var todos []Todo
	db := dbConn()

	userIdMod := strings.Replace(userId, "-", "_", 4)

	myDB := dbConn()
	addTable, err := myDB.Prepare("CREATE TABLE IF NOT EXISTS " + userIdMod + " (ID varchar(46), Message varchar(255), Complete boolean)")
	if err != nil {
		panic(err.Error())
	}
	addTable.Exec()
	log.Println("Table " + userIdMod + " created successfully")

	rows, err := db.Query("SELECT * FROM " + userIdMod)
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

func DatabaseAdd(userId string, ID string, Message string, Complete bool) {

	db := dbConn()

	userIdMod := strings.Replace(userId, "-", "_", 4)

	r, err := db.Prepare("INSERT INTO " + userIdMod + "(ID, Message, Complete) VALUES(?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}
	r.Exec(ID, Message, Complete)
	db.Close()
	log.Println("New item created successfully")
}

func DatabaseComplete(userId string, ID string) {

	db := dbConn()

	userIdMod := strings.Replace(userId, "-", "_", 4)

	r, err := db.Prepare("UPDATE " + userIdMod + " SET Complete=true WHERE ID=?")
	if err != nil {
		panic(err.Error())
	}
	r.Exec(ID)
	db.Close()
	log.Println("Item" + ID + "from " + userIdMod + " set to completed status")
}

func DatabaseDelete(userId string, ID string) {

	db := dbConn()

	userIdMod := strings.Replace(userId, "-", "_", 4)

	r, err := db.Prepare("DELETE FROM " + userIdMod + " WHERE ID=?")
	if err != nil {
		panic(err.Error())
	}
	r.Exec(ID)
	db.Close()
	log.Println("Item " + ID + " from " + userIdMod + " removed successfully")
}
