package todo

import (
	"errors"
	"golang-angular/dbdriver"
	"sync"

	"github.com/rs/xid"
)

var (
	list        []dbdriver.Todo
	mtx         sync.RWMutex
	once1       sync.Once
	once2       bool
	extDBexists bool
)

func init() {
	once1.Do(initialiseList)
}

func initialiseList() {
	list = []dbdriver.Todo{}
	once2 = true

}

// Get retrieves all elements from the todo list
func Get() []dbdriver.Todo {

	if once2 {
		extDBexists = dbdriver.InitialiseDB()
		once2 = false
	}

	var dblist = []dbdriver.Todo{}
	if extDBexists {
		dblist = append(dblist, dbdriver.DatabaseGet()...)
		println(dblist)
		println(list)
		return dblist
	}
	return list
}

// Add will add a new todo based on a message
func Add(message string) string {
	t := newTodo(message)
	mtx.Lock()
	list = append(list, t)
	if extDBexists {
		dbdriver.DatabaseAdd(t.ID, t.Message, t.Complete)
	}
	mtx.Unlock()
	return t.ID
}

// Delete will remove a Todo from the Todo list
func Delete(id string) error {
	location, err := findTodoLocation(id)
	if err != nil {
		return err
	}
	removeElementByLocation(location)
	if extDBexists {
		dbdriver.DatabaseDelete(id)
	}
	return nil
}

// Complete will set the complete boolean to true, marking a todo as
// completed
func Complete(id string) error {
	location, err := findTodoLocation(id)
	if err != nil {
		return err
	}
	setTodoCompleteByLocation(location)
	if extDBexists {
		dbdriver.DatabaseComplete(id)
	}
	return nil
}

func newTodo(msg string) dbdriver.Todo {
	return dbdriver.Todo{
		ID:       xid.New().String(),
		Message:  msg,
		Complete: false,
	}
}

func findTodoLocation(id string) (int, error) {
	mtx.RLock()
	defer mtx.RUnlock()
	for i, t := range list {
		if isMatchingID(t.ID, id) {
			return i, nil
		}
	}
	return 0, errors.New("could not find todo based on id")
}

func removeElementByLocation(i int) {
	mtx.Lock()
	list = append(list[:i], list[i+1:]...)
	mtx.Unlock()
}

func setTodoCompleteByLocation(location int) {
	mtx.Lock()
	list[location].Complete = true
	mtx.Unlock()
}

func isMatchingID(a string, b string) bool {
	return a == b
}
