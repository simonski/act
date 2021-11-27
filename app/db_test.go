package app

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func NewTestDB() *TaskyDB {
	tempfile, _ := ioutil.TempFile("", "test-act.db")
	filename := tempfile.Name()
	os.Remove(filename)
	tdb := NewTaskyDB(filename)
	return tdb
}

func TestDBEmpty(t *testing.T) {
	tdb := NewTestDB()
	tdb.Init()
	tasks := tdb.ListTasks()
	if len(tasks) != 0 {
		fmt.Printf("failed.")
		os.Exit(1)
	}

}

func TestDBNotEmpty(t *testing.T) {
	tdb := NewTestDB()
	tdb.Init()
	tdb.AddTask("fred")
	tasks := tdb.ListTasks()
	if len(tasks) != 1 {
		fmt.Printf("failed.")
		os.Exit(1)
	}

	// command := "fooo"
	// cli := goutils.NewCLI(os.Args)
	// if command == "test" {
	// 	db := LoadDB()
	// 	value := cli.GetStringOrDie(command)
	// 	valueEnc := db.Encrypt(value)
	// 	fmt.Printf("Encrypt('%v')=\n%v\n", value, valueEnc)
	// 	valueDec := db.Decrypt(valueEnc)
	// 	fmt.Printf("\n\nDecrypt\n '%v'\n", valueDec)
	// 	os.Exit(0)
	// }

}

func TestDBCanUpdateName(t *testing.T) {
	tdb := NewTestDB()
	tdb.Init()
	tdb.AddTask("fred")
	tdb.AddTask("jack")
	tasks := tdb.ListTasks()
	if len(tasks) != 2 {
		fmt.Printf("failed.")
		os.Exit(1)
	}

	fredTask := tdb.GetTaskById("1")
	fredTask.name = "jim"
	tdb.Save(fredTask)

	t2 := tdb.GetTaskById("1")
	if t2.name != "jim" {
		t.Log("cannot update name.")
		t.Fail()
	}
}

func TestDBConfigCRUD(t *testing.T) {
	tdb := NewTestDB()
	tdb.Init()
	tdb.AddTask("fred")
	tdb.AddTask("jack")
	tasks := tdb.ListTasks()
	if len(tasks) != 2 {
		fmt.Printf("failed.")
		os.Exit(1)
	}

	fredTask := tdb.GetTaskById("1")
	fredTask.name = "jim"
	tdb.Save(fredTask)

	t2 := tdb.GetTaskById("1")
	if t2.name != "jim" {
		t.Log("cannot update name.")
		t.Fail()
	}
}
