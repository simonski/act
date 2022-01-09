package app

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/simonski/goutils"
)

// )

type Task struct {
	Task_id     int
	Created     *time.Time
	Updated     *time.Time
	Due         *time.Time
	User_id     int
	Project_id  int
	State       string
	Name        string
	Description string
	Deleted     bool
	Archived    bool
}

type Project struct {
	project_id  int
	created     *time.Time
	updated     *time.Time
	user_id     int
	state       string
	name        string
	description string
	deleted     bool
	archived    bool
}

type User struct {
	user_id  int
	created  *time.Time
	updated  *time.Time
	username string
}

type Config struct {
	project_id int
	name       string
	value      string
}

type TaskComment struct {
	comment_id  int
	created     *time.Time
	user_id     int
	task_id     int
	comment     string
	description string
	deleted     bool
}

type ProjectComment struct {
	comment_id  int
	created     *time.Time
	user_id     int
	project_id  int
	comment     string
	description string
	deleted     bool
}

// KPDB helper struct holds the data and keys
type ActDB struct {
	db     *sql.DB
	Config *ActDBConfig
}

type ActDBConfig struct {
	IsSqlite       bool
	IsPostgres     bool
	SqliteFilename string
	PgHost         string
	PgPort         int
	PgUser         string
	PgPassword     string
	PgDbName       string
}

func NewActDBConfig(cli *goutils.CLI) *ActDBConfig {
	dbType := cli.GetStringOrDie("-type")
	if dbType == "postgres" {
		PgHost := cli.GetStringOrDie("-host")
		PgPort, _ := strconv.Atoi(cli.GetStringOrDie("-port"))
		PgUser := cli.GetStringOrDie("-user")
		PgPassword := cli.GetStringOrDefault("-password", "")
		PgDbName := cli.GetStringOrDie("-name")
		config := ActDBConfig{IsPostgres: true, PgHost: PgHost, PgPort: PgPort, PgUser: PgUser, PgPassword: PgPassword, PgDbName: PgDbName}
		return &config

	} else if dbType == "sqlite" {
		SqliteFilename := cli.GetStringOrDie("-file")
		config := ActDBConfig{IsSqlite: true, SqliteFilename: SqliteFilename}
		return &config
	} else {
		fmt.Printf("-type can be 'postgres' or 'sqlite'\n")
		return nil
	}

}

func NewActDB(config *ActDBConfig) *ActDB {
	return &ActDB{Config: config}
}

func (adb *ActDB) NewProject() *Project {
	pc := Project{}
	return &pc
}
func (adb *ActDB) NewTask(project *Project) *Task {
	t := Task{}
	t.Project_id = project.project_id
	return &t
}
func (adb *ActDB) NewConfig(project *Project) *Config {
	c := Config{}
	c.project_id = project.project_id
	return &c
}
func (adb *ActDB) NewUser() *User {
	u := User{}
	return &u
}
func (adb *ActDB) NewProjectComment(project *Project) *ProjectComment {
	pc := ProjectComment{}
	return &pc
}
func (adb *ActDB) NewTaskComment(task *Task) *TaskComment {
	tc := TaskComment{}
	return &tc
}

func (adb *ActDB) Disconnect() bool {
	err := adb.db.Close()
	if err != nil {
		panic(err)
	}
	return true
}

func (adb *ActDB) ConnectNoDb() bool {
	if adb.Config.IsSqlite {
		db, err := sql.Open("sqlite3", adb.Config.SqliteFilename)
		checkErr(err)
		adb.db = db
	} else if adb.Config.IsPostgres {
		var psqlInfo string
		if adb.Config.PgPassword != "" {
			psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
				"password=%s sslmode=disable",
				adb.Config.PgHost, adb.Config.PgPort, adb.Config.PgUser, adb.Config.PgPassword)
		} else {
			psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
				"sslmode=disable",
				adb.Config.PgHost, adb.Config.PgPort, adb.Config.PgUser)

		}
		db, err := sql.Open("postgres", psqlInfo)
		adb.db = db
		if err != nil {
			panic(err)
		}
	}
	return true
}

// Load populates the db with the file
func (adb *ActDB) Connect() bool {
	if adb.Config.IsSqlite {
		db, err := sql.Open("sqlite3", adb.Config.SqliteFilename)
		checkErr(err)
		adb.db = db
	} else if adb.Config.IsPostgres {
		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=disable",
			adb.Config.PgHost, adb.Config.PgPort, adb.Config.PgUser, adb.Config.PgPassword, adb.Config.PgDbName)
		db, err := sql.Open("postgres", psqlInfo)
		adb.db = db
		if err != nil {
			panic(err)
		}
	}
	// defer adb.db.Close()

	// cdb.Filename = goutils.EvaluateFilename(filename)
	// cdb.PublicKeyFilename = goutils.EvaluateFilename(pubKey)
	// cdb.PrivateKeyFilename = goutils.EvaluateFilename(privKey)
	// cdb.EncryptionEnabled = encryptionEnabled

	// if !goutils.FileExists(cdb.Filename) {
	// 	db := DB{}
	// 	db.Entries = make(map[string]DBEntry)
	// 	cdb.data = db
	// } else {
	// 	jsonFile, err := os.Open(cdb.Filename)
	// 	if err != nil {
	// 		fmt.Printf("ERR %v\n", err)
	// 		db := DB{}
	// 		db.Entries = make(map[string]DBEntry)
	// 		cdb.data = db
	// 		// panic(err)
	// 	} else {
	// 		db := DB{}
	// 		bytes, _ := ioutil.ReadAll(jsonFile)
	// 		var data map[string]DBEntry
	// 		json.Unmarshal(bytes, &data)
	// 		db.Entries = data
	// 		cdb.data = db
	// 	}
	// }

	return true
}

func (adb *ActDB) Init() bool {
	db := adb.db
	var sqls []string
	if adb.Config.IsSqlite {
		sqls = strings.Split(SQL_SCHEMA_SQLITE, ";")
	} else {

		// need to create the db first
		// dbName := adb.Config.PgDbName
		adb.Disconnect()
		adb.ConnectNoDb()
		db := adb.db
		sql := fmt.Sprintf("create database %v;", adb.Config.PgDbName)
		_, err := db.Exec(sql)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

		// stmt, err := db.Prepare(sql)
		// // fmt.Printf("SQL: %v\n", value)
		// msg, err := stmt.Exec()
		// fmt.Println(msg)
		// checkErr(err)
		// adb.Config.PgDbName = dbName

		sqls = strings.Split(SQL_SCHEMA_POSTGRES, ";")
	}
	for _, value := range sqls {
		// value = strings.ReplaceAll(value, "\n", " ")
		if strings.Trim(value, " \n") == "" {
			continue
		}
		value = strings.Trim(value, " ") + ";"
		value = strings.ReplaceAll(value, "\n", "")
		if value != "" && strings.Index(value, "--") != 0 {
			// fmt.Printf("Not a comment, -- index = %v\n", strings.Index(value, "--"))
			// fmt.Printf("\n%v\n", value)
			stmt, err := db.Prepare(value)
			fmt.Printf("SQL: %v\n", value)
			_, err = stmt.Exec()

			checkErr(err)
		}
	}

	adb.AddConfig("created", time.Now().Format(time.RFC3339Nano))
	adb.AddConfig("version", VERSION)
	return true
}

func (adb *ActDB) AddConfig(name string, value string) {
	db := adb.db
	sql := fmt.Sprintf("insert into config (name, value) values (\"%v\", \"%v\");", name, value)
	fmt.Printf("SQL: '%v'\n", sql)
	_, err := db.Exec(sql)
	// fmt.Printf("%v\n", result)
	checkErr(err)
}

// func DoSql(cli *goutils.CLI) {
// 	db, err := sql.Open("sqlite3", "./foo.db")
// 	checkErr(err)

// 	stmt, err := db.Prepare(SQL_CREATE)
// 	res, err := stmt.Exec()
// 	checkErr(err)

// 	db, err := sql.Open("sqlite3", filename)
// 	checkErr(err)
// 	adb.db = db

// }

// Clear empties the db (without saving it)
func (adb *ActDB) Clear() {
	// cdb.data.Entries = make(map[string]DBEntry)
}

func (adb *ActDB) AddTask(name string) {
	db := adb.db
	stmt, err := db.Prepare("INSERT INTO tasks(user_id, project_id, created, updated, state, name, description, deleted, archived) values(?,?,?,?,?,?,?,?, ?)")
	checkErr(err)

	user_id := 1
	project_id := 1
	created := time.Now().Format(time.RFC3339Nano)
	updated := time.Now().Format(time.RFC3339Nano)

	fmt.Printf("New task created %v, updated %v\n", created, updated)

	state := "created"
	// name := name
	description := ""
	deleted := false
	archived := false
	_, err = stmt.Exec(user_id, project_id, created, updated, state, name, description, deleted, archived)
	checkErr(err)

}

func (adb *ActDB) ListTasks() []*Task {
	db := adb.db
	rows, err := db.Query("SELECT task_id, project_id, created, updated, due, name, state FROM tasks")
	checkErr(err)
	var task_id int
	var project_id int
	var created string
	var updated string
	var due string
	var name string
	var state string

	var results []*Task
	for rows.Next() {
		err = rows.Scan(&task_id, &project_id, &created, &updated, &due, &name, &state)
		fmt.Printf("created %v\n", created)
		checkErr(err)
		t := &Task{}
		t.Task_id = task_id
		t.Project_id = project_id

		ti, _ := time.Parse(time.RFC3339Nano, created)
		t.Created = &ti

		up, _ := time.Parse(time.RFC3339Nano, updated)
		t.Updated = &up

		due_dt, _ := time.Parse(time.RFC3339Nano, due)
		t.Due = &due_dt

		t.State = state

		// t.updated = updated
		t.Name = name
		results = append(results, t)
	}

	rows.Close() //good habit to close
	return results

}

func (adb *ActDB) GetTaskById(taskId string) *Task {
	db := adb.db
	rows, err := db.Query("SELECT task_id, project_id, created, name, state FROM tasks where task_id=?", taskId)
	checkErr(err)
	var task_id int
	var project_id int
	var created string
	var name string
	var state string
	rows.Next()
	err = rows.Scan(&task_id, &project_id, &created, &name, &state)
	if err != nil {
		return nil
	}
	checkErr(err)
	t := &Task{}
	t.Task_id = task_id
	t.Project_id = project_id
	t.State = state

	cr, _ := time.Parse(time.RFC3339Nano, created)
	t.Created = &cr

	t.Name = name
	rows.Close() //good habit to close
	return t

}

func (adb *ActDB) Save(task *Task) {
	db := adb.db
	if task.Task_id == 0 {
		adb.AddTask(task.Name)
		return
	}
	updated := time.Now().Format(time.RFC3339Nano)

	stmt, err := db.Prepare("UPDATE tasks SET name=?, updated=? WHERE task_id = ?")
	checkErr(err)
	_, err = stmt.Exec(task.Name, updated, task.Task_id)
	checkErr(err)
}

func (adb *ActDB) Demo() bool {

	db := adb.db

	// insert
	stmt, err := db.Prepare("INSERT INTO userinfo(username, departname, created) values(?,?,?)")
	checkErr(err)

	res, err := stmt.Exec("astaxie", "研发部门", "2012-12-09")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)
	// update
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)

	res, err = stmt.Exec("astaxieupdate", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	// query
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)
	var uid int
	var username string
	var department string
	var created time.Time

	for rows.Next() {
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}

	rows.Close() //good habit to close

	// delete
	stmt, err = db.Prepare("delete from userinfo where uid=?")
	checkErr(err)

	res, err = stmt.Exec(id)
	checkErr(err)

	affect, err = res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	db.Close()

	// trashSQL, err := database.Prepare("update task set is_deleted='Y',last_modified_at=datetime() where id=?")
	// if err != nil {
	//     fmt.Println(err)
	// }
	// tx, err := database.Begin()
	// if err != nil {
	//     fmt.Println(err)
	// }
	// _, err = tx.Stmt(trashSQL).Exec(id)
	// if err != nil {
	//     fmt.Println("doing rollback")
	//     tx.Rollback()
	// } else {
	//     tx.Commit()
	// }

	return true
}
