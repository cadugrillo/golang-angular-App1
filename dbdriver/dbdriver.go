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
	//mailIdint := strings.Replace(mail, "@", "_", 1)
	//mailId := strings.Replace(mailIdint, ".", "_", 1)

	//myDB := dbConn("todo")
	//addTable, err := myDB.Prepare("CREATE TABLE IF NOT EXISTS " + mailId + " (ID varchar(36), Message varchar(255), Complete boolean)")
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

	mailIdint := strings.Replace(userId, "@", "_", 1)
	mailId := strings.Replace(mailIdint, ".", "_", 1)

	myDB := dbConn()
	addTable, err := myDB.Prepare("CREATE TABLE IF NOT EXISTS " + mailId + " (ID varchar(36), Message varchar(255), Complete boolean)")
	if err != nil {
		panic(err.Error())
	}
	addTable.Exec()
	log.Println("Table " + mailId + " created successfully")

	rows, err := db.Query("SELECT * FROM " + mailId)
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

	mailIdint := strings.Replace(userId, "@", "_", 1)
	mailId := strings.Replace(mailIdint, ".", "_", 1)

	r, err := db.Prepare("INSERT INTO " + mailId + "(ID, Message, Complete) VALUES(?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}
	r.Exec(ID, Message, Complete)
	db.Close()
	log.Println("New item created successfully")
}

func DatabaseComplete(userId string, ID string) {

	db := dbConn()

	mailIdint := strings.Replace(userId, "@", "_", 1)
	mailId := strings.Replace(mailIdint, ".", "_", 1)

	r, err := db.Prepare("UPDATE " + mailId + " SET Complete=true WHERE ID=?")
	if err != nil {
		panic(err.Error())
	}
	r.Exec(ID)
	db.Close()
	log.Println("Item" + ID + "from " + mailId + " set to completed status")
}

func DatabaseDelete(userId string, ID string) {

	db := dbConn()

	mailIdint := strings.Replace(userId, "@", "_", 1)
	mailId := strings.Replace(mailIdint, ".", "_", 1)

	r, err := db.Prepare("DELETE FROM " + mailId + " WHERE ID=?")
	if err != nil {
		panic(err.Error())
	}
	r.Exec(ID)
	db.Close()
	log.Println("Item " + ID + " from " + mailId + " removed successfully")
}
