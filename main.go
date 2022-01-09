package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	_ "github.com/mattn/go-sqlite3"

	"github.com/simonski/act/app"
	goutils "github.com/simonski/goutils"

	// make the terminal so blue
	figure "github.com/common-nighthawk/go-figure"
)

// const USAGE = ` is a terminal tool to track tasks.

// Usage:

//          <command> [arguments]

// Environment variables

// 	ACT_FILE       path to  database file to use
// 	ACT_URL        url of running  db server

// Client commands are:

// 	init             create new  database
// 	add              add a task
// 	update           update a task (-id X -name Y -description Z)

// 	info             lists information about the task db
// 	server           run a  server daemon

// 	status           lists the current tasks you are working on
// 	ls               lists tasks and projects
// 	create           creates a task or project
// 	rm               deletes a task or project
// 	complete         marks a task or project as complete

// 	add              adds a new task
// 	ls

// 	help             print this text
// 	version          prints  version

// Usage " help <command>" for more information.

// `

const USAGE = ` is a terminal tool to track tasks.

Usage:

    act <command> [arguments]

Environment variables

	ACT_FILE       path to  database file to use
	ACT_URL        url of running  db server

Client commands are:

	init             create new  database
	add              add a task
	update           update a task (-id X -name Y -description Z)
	note             add a note to a task
						 note 3 blah blah blah

	info             lists information about the task db

	status           lists the current tasks you are working on		
	ls               lists tasks and projects
	create           creates a task or project

	rm               deletes a task or project
	complete         marks a task or project as complete

Project commands

	project          
		set
		create
		ls


State commands

	workon           general purpose work on a thing list what we are working on
	complete         finish a task
	resume           restart a task
	pause            pause working on a task
	history			 shows reversse order of tasks you've been working on 

Multiplayer commands

	server           run as a server
	server-add		 add a server
	server-update	 update a server
	server-rm		 remove a server
	sync			 merge two or more servers

General commands

	help             print this text
	version          prints  version

Usage " help <command>" for more information.

`

func main() {
	cli := goutils.NewCLI(os.Args)
	command := cli.GetCommand()
	if command == "help" {
		DoLogo(os.Args[0])
		DoUsage(cli)
	} else if command == "init" {
		DoInit(cli)
	} else if command == "add" {
		DoAdd(cli)
	} else if command == "ls" {
		DoList(cli)
	} else if command == "update" {
		DoUpdate(cli)
	} else if command == "rm" {
		DoRm(cli)
	} else if command == "workon" {
		DoWorkOn(cli)

	} else if command == "sql" {
		DoSql(cli)
	} else if isInfo(command) {
		DoInfo(cli)
	} else if isVersion(command) {
		DoVersion(cli)
	} else if command != "" {
		fmt.Printf(" %v: unknown command\n", command)
		fmt.Printf("Run ' help' for usage.\n")
		os.Exit(1)
	} else {
		DoLogo(os.Args[0])
		DoUsage(cli)
	}

	// } else if isVerify(command) {
	// 	DoVerify(cli, true)
	// } else if isList(command) {
	// 	DoList(cli)
	// } else if isClear(command) {
	// 	DoClear(cli)
	// } else if isDescribe(command) {
	// 	DoDescribe(cli)
	// } else if isPut(command, cli) {
	// 	encryptionEnabled := goutils.GetEnvOrDefault(KP_ENCRYPTION, "0") == "1"
	// 	ok := true
	// 	if encryptionEnabled {
	// 		ok = DoVerify(cli, false)
	// 	}
	// 	if ok {
	// 		DoPut(cli)
	// 	}
	// } else if isGet(command, cli) {
	// 	encryptionEnabled := goutils.GetEnvOrDefault(KP_ENCRYPTION, "0") == "1"
	// 	ok := true
	// 	if encryptionEnabled {
	// 		ok = DoVerify(cli, false)
	// 	}
	// 	if ok {
	// 		DoGet(cli)
	// 	}
	// } else if isDelete(command) {
	// 	DoDelete(cli)
}

func isVerify(command string) bool {
	return command == "verify"
}

func isInfo(command string) bool {
	return command == "info"
}

func isDelete(command string) bool {
	return command == "rm"
}

func isVersion(command string) bool {
	return command == "version"
}

func isList(command string) bool {
	return command == "ls"
}

func isClear(command string) bool {
	return command == "clear"
}

func isDescribe(command string) bool {
	return command == "describe"
}

// A 'get' is basically not a a list, delete or a put
func isGet(command string, cli *goutils.CLI) bool {
	return command == "get"
}

func isPut(command string, cli *goutils.CLI) bool {
	return command == "put"
}

// DoVerify performs verification of ~/.KPfile, encryption/decryption using
// specified keys
func DoVerify(cli *goutils.CLI, printFailuresToStdOut bool) bool {
	overallValid := true

	return overallValid
	// fmt.Printf("%v =%v, exists=%v\n", KP_ENCRYPTION_ENABLED, privateKeyExists)
}

func DoInit(cli *goutils.CLI) {
	// filename := GetFileName(cli)
	// if goutils.FileExists(filename) {
	// 	fmt.Printf("> Error, file '%v' exists.\n", filename)
	// 	os.Exit(1)
	// }
	config := app.NewActDBConfig(cli)
	tdb := app.NewActDB(config)
	tdb.Connect()
	tdb.Init()
	// fmt.Printf("> %v created.\n", filename)
	// db, err := sql.Open("sqlite3", "./.db")
	// checkErr(err)

	// stmt, err := db.Prepare(SQL_INIT_1)
	// _, err = stmt.Exec()
	// checkErr(err)

	// stmt, err = db.Prepare(SQL_INIT_2)
	// _, err = stmt.Exec()
	// checkErr(err)
}

func GetFileName(cli *goutils.CLI) string {
	default_filename := goutils.GetEnvOrDefault(app.ACT_FILE, "./actdb")
	return cli.GetStringOrDefault("-file", default_filename)
}

func DoAdd(cli *goutils.CLI) {
	taskName := cli.GetStringOrDie("add")
	fmt.Printf("Task name is '%v'\n", taskName)
	// filename := GetFileName(cli)
	config := app.NewActDBConfig(cli)
	tdb := app.NewActDB(config)
	tdb.Connect()
	// tdb := app.NewActDB(filename)
	tdb.AddTask(taskName)

}

func DoWorkOn(cli *goutils.CLI) {
	x := cli.GetStringOrDie("workon")
	fmt.Printf("Workon '%v'\n", x)
	// fmt.Printf("Task name is '%v'\n", taskName)
	// filename := GetFileName(cli)
	// tdb := NewActDB(filename)
	// tdb.AddTask(taskName)
}

func DoUpdate(cli *goutils.CLI) {
	tdb := GetActDB(cli)
	taskId := cli.GetStringOrDie("-id")
	task := tdb.GetTaskById(taskId)
	if task == nil {
		fmt.Printf("Task id '%v' does not exist.\n", taskId)
		os.Exit(1)
	}
	name := cli.GetStringOrDefault("-name", task.Name)
	task.Name = name
	tdb.Save(task)
}

func DoRm(cli *goutils.CLI) {
	tdb := GetActDB(cli)
	taskId := cli.GetStringOrDie("-id")
	task := tdb.GetTaskById(taskId)
	if task == nil {
		fmt.Printf("Task id '%v' does not exist.\n", taskId)
		os.Exit(1)
	}
	task.Deleted = true
	tdb.Save(task)
}

func GetActDB(cli *goutils.CLI) *app.ActDB {
	filename := GetFileName(cli)
	if !goutils.FileExists(filename) {
		fmt.Printf("> Error, '%v' cannot be found or does not exist (try ' init').\n", filename)
		os.Exit(1)
	}
	config := app.NewActDBConfig(cli)
	tdb := app.NewActDB(config)
	tdb.Connect()
	return tdb
}

func DoList(cli *goutils.CLI) {
	tdb := GetActDB(cli)
	tasks := tdb.ListTasks()
	if len(tasks) == 0 {
		fmt.Printf("> No tasks.\n")
		return
	}

	maxIdLen := 3
	maxNameLen := 4
	maxUpdatedLen := 7
	maxCreatedLen := 7
	// maxStatusLen := 6

	// find nice padding sizes
	for _, task := range tasks {
		maxIdLen = goutils.Max(maxIdLen, len(strconv.Itoa(task.Task_id)))
		maxNameLen = goutils.Max(maxNameLen, len(task.Name))
		maxUpdatedLen = goutils.Max(maxUpdatedLen, 10)
		maxCreatedLen = goutils.Max(maxCreatedLen, 10)
	}

	TOP_LEFT_CORNER := "0x6C"
	TOP_RIGHT_CORNER := 0x6B
	BOTTOM_RIGHT_CORNER := "┘" // ┘

	fmt.Printf("\ntopleft: %v    topright  %c   bottomleft %v \n", TOP_LEFT_CORNER, TOP_RIGHT_CORNER, BOTTOM_RIGHT_CORNER)

	const (
		CORNER_BOTTOM_RIGHT    = "┘"
		CORNER_TOP_RIGHT       = "┐"
		CORNER_TOP_LEFT        = "┌"
		CORNER_BOTTOM_LEFT     = "└"
		INTERSECT_CROSS        = "┼"
		HORIZONTAL             = "─"
		VERTICAL_RIGHT_JOIN    = "├"
		VERTICAL_LEFT_JOIN     = "┤"
		HORIZONTAL_TOP_JOIN    = "┴"
		HORIZONTAL_BOTTOM_JOIN = "┬"
		VERTICAL               = "│"

		COLOR_Reset  = "\033[0m"
		COLOR_Red    = "\033[31m"
		COLOR_Green  = "\033[32m"
		COLOR_Yellow = "\033[33m"
		COLOR_Blue   = "\033[34m"
		COLOR_Purple = "\033[35m"
		COLOR_Cyan   = "\033[36m"
		COLOR_Gray   = "\033[37m"
		COLOR_White  = "\033[97m"
	)

	fmt.Printf("%v%v\n", COLOR_Reset, COLOR_Red)
	fmt.Printf("%v%v%v%v%v%v%v\n", CORNER_TOP_LEFT, HORIZONTAL, HORIZONTAL, HORIZONTAL, HORIZONTAL, HORIZONTAL, CORNER_TOP_RIGHT)
	fmt.Printf("%v%v%v%v%v%v%v\n", CORNER_BOTTOM_LEFT, HORIZONTAL, HORIZONTAL, HORIZONTAL, HORIZONTAL, HORIZONTAL, CORNER_BOTTOM_RIGHT)
	fmt.Printf("%v\n", COLOR_Reset)

	// a = sprintf("%c", 0x6C) + # ┌
	// sprintf("%c", 0x71) + # ─
	// sprintf("%c", 0x71) + # ─
	// sprintf("%c", 0x71) + # ─
	// sprintf("%c", 0x71) + # ─
	// sprintf("%c", 0x71) + # ─
	// sprintf("%c", 0x71) + # ─
	// sprintf("%c\n", 0x6B) +  # ┐
	// sprintf("%c", 0x78) + # │
	// #print("      ")
	// "      " +
	// sprintf("%c\n", 0x78) + # │
	// sprintf("%c", 0x6D) + # └
	// sprintf("%c", 0x71) + # ─
	// sprintf("%c", 0x71) + # ─
	// sprintf("%c", 0x71) + # ─
	// sprintf("%c", 0x71) + # ─
	// sprintf("%c", 0x71) + # ─
	// sprintf("%c", 0x71) + # ─
	// sprintf("%c", 0x6A)  # ┘

	header := fmt.Sprintf("| %5v | %10v | %10v | %10v | %10v |\n", "ID", "Name", "Updated", "Created", "Status")
	line := strings.Repeat("-", len(header))
	fmt.Println(line)
	fmt.Println(header)
	fmt.Println(line)

	for _, task := range tasks {
		taskLine := fmt.Sprintf("| %5v | %10v | %10v | %10v | %10v |\n", task.Task_id, task.Name, task.Updated.Format("2006-01-02"), task.Created.Format("2006-01-02"), task.State)
		fmt.Println(taskLine)
	}

	fmt.Println(line)

}

func DoSql(cli *goutils.CLI) {
	tdb := GetActDB(cli)
	tdb.Demo()
}

func DoLogo(appName string) {
	appName = strings.ReplaceAll(appName, "./", "")
	f := figure.NewColorFigure(appName, "", "blue", true)
	f.Print()
}

func DoInfo(cli *goutils.CLI) {

	filename := GetFileName(cli)

	fmt.Printf("\nAct is currently using the following values:\n")
	fmt.Printf("\n%v          : %v\n", app.ACT_FILE, filename)
	fmt.Printf("\n")
	// t := NewTerminal()

	// sysInfo := goutils.NewSysInfo()

	// fmt.Printf("RAM         : %v\n", sysInfo.RAM)
	// fmt.Printf("CPU         : %v\n", sysInfo.CPU)
	// fmt.Printf("Cores	    : %v\n", runtime.NumCPU())
	// fmt.Printf("Disk        : %v\n", sysInfo.Disk)
	// fmt.Printf("Hostname    : %v\n", sysInfo.Hostname)
	// fmt.Printf("GOOS        : %v\n", runtime.GOOS)
	// fmt.Printf("GOARCH      : %v\n", runtime.GOARCH)
	// fmt.Printf("GOMAXPROC   : %v\n", runtime.GOMAXPROCS)
	// fmt.Printf("Columns     : %v\n", t.Width())
	// fmt.Printf("IsMacOS     : %v\n", sysInfo.IsMacOS())
	// fmt.Printf("IsLinux     : %v\n", sysInfo.IsLinux())
	// fmt.Printf("IsWindows   : %v\n", sysInfo.IsWindows())

}

func DoUsage(cli *goutils.CLI) {
	fmt.Println(USAGE)
}

func DoVersion(cli *goutils.CLI) {
	fmt.Printf("%v\n", app.VERSION)
}
