package app

// import (
// 	"database/sql"
// 	"fmt"
// 	"strings"
// 	"time"

// 	_ "github.com/mattn/go-sqlite3"
// )

// const SQL_SCHEMA = `

// DROP TABLE IF EXISTS config;
// CREATE TABLE IF NOT EXISTS config (
// 	name STRING UNIQUE NOT NULL,
// 	value STRING NOT NULL
// );

// -- DROP TABLE IF EXISTS users;

// DROP TABLE IF EXISTS users;
// CREATE TABLE IF NOT EXISTS users (
// 	user_id  INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,
// 	created TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
// 	updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
// 	username VARCHAR(64) NULL UNIQUE
// );
// -- INSERT INTO USERS (username) values  ('hi');

// DROP TABLE IF EXISTS projects;
// CREATE TABLE IF NOT EXISTS projects (
// 	project_id INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,
// 	created TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
// 	updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
// 	name STRING NOT NULL UNIQUE,
// 	user_id INTEGER,
// 	description STRING NULL,
// 	state STRING NOT NULL,
// 	deleted BOOLEAN NOT NULL DEFAULT false,
// 	archived BOOLEAN NOT NULL DEFAULT false,
// 	FOREIGN KEY (user_id) REFERENCES users(user_id)
// );

// DROP TABLE IF EXISTS tasks;
// CREATE TABLE IF NOT EXISTS tasks (
// 	task_id INTEGER PRIMARY KEY AUTOINCREMENT,
// 	created TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
// 	updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
// 	due TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
// 	user_id INTEGER,
// 	project_id INTEGER,
// 	state STRING NOT NULL,
// 	name STRING NOT NULL,
// 	description STRING NULL,
// 	deleted BOOLEAN DEFAULT FALSE,
// 	archived BOOLEAN DEFAULT FALSE,
// 	FOREIGN KEY (user_id) REFERENCES users(user_id),
// 	FOREIGN KEY (project_id) REFERENCES projects(project_id)
// );

// DROP TABLE IF EXISTS project_comments;
// CREATE TABLE IF NOT EXISTS project_comments (
// 	comment_id INTEGER PRIMARY KEY AUTOINCREMENT,
// 	created TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
// 	user_id INTEGER,
// 	project_id INTEGER,
// 	comment STRING NOT NULL,
// 	description STRING NULL,
// 	deleted BOOLEAN DEFAULT FALSE,
// 	archived BOOLEAN DEFAULT FALSE,
// 	FOREIGN KEY (user_id) REFERENCES users(user_id),
// 	FOREIGN KEY (project_id) REFERENCES projects(project_id)
// );

// DROP TABLE IF EXISTS task_comments;
// CREATE TABLE IF NOT EXISTS task_comments (
// 	comment_id INTEGER PRIMARY KEY AUTOINCREMENT,
// 	created TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
// 	user_id INTEGER,
// 	task_id INTEGER,
// 	comment STRING NOT NULL,
// 	description STRING NULL,
// 	deleted BOOLEAN DEFAULT FALSE,
// 	archived BOOLEAN DEFAULT FALSE,
// 	FOREIGN KEY (user_id) REFERENCES users(user_id),
// 	FOREIGN KEY (task_id) REFERENCES tasks(task_id)
// );
// `

// // type DBObject interface {
// // 	GetId() int
// // 	Save(tdb *ActDB)
// // }

// type Task struct {
// 	Task_id     int
// 	Created     *time.Time
// 	Updated     *time.Time
// 	Due         *time.Time
// 	User_id     int
// 	Project_id  int
// 	State       string
// 	Name        string
// 	Description string
// 	Deleted     bool
// 	Archived    bool
// }

// func NewTask() *Task {
// 	t := Task{}
// 	return &t
// }

// type Project struct {
// 	project_id  int
// 	created     *time.Time
// 	updated     *time.Time
// 	user_id     int
// 	state       string
// 	name        string
// 	description string
// 	deleted     bool
// 	archived    bool
// }

// func NewProject() *Project {
// 	p := Project{}
// 	return &p
// }

// type User struct {
// 	user_id  int
// 	created  *time.Time
// 	updated  *time.Time
// 	username string
// }

// func NewUser() *User {
// 	u := User{}
// 	return &u
// }

// type Config struct {
// 	project_id int
// 	name       string
// 	value      string
// }

// func NewConfig() *Config {
// 	c := Config{}
// 	return &c
// }

// type TaskComment struct {
// 	comment_id  int
// 	created     *time.Time
// 	user_id     int
// 	task_id     int
// 	comment     string
// 	description string
// 	deleted     bool
// }

// type ProjectComment struct {
// 	comment_id  int
// 	created     *time.Time
// 	user_id     int
// 	project_id  int
// 	comment     string
// 	description string
// 	deleted     bool
// }

// // KPDB helper struct holds the data and keys
// type ActDB struct {
// 	db *sql.DB
// }

// // NewKPDB constructor
// func NewActDB(filename string) *ActDB {
// 	tdb := ActDB{}
// 	tdb.Load(filename)
// 	return &tdb
// }

// func (tbh *ActDB) NewProject() *Project {
// 	pc := Project{}
// 	return &pc
// }
// func (tbh *ActDB) NewTask(project *Project) *Task {
// 	t := Task{}
// 	t.Project_id = project.project_id
// 	return &t
// }
// func (tbh *ActDB) NewConfig(project *Project) *Config {
// 	c := Config{}
// 	c.project_id = project.project_id
// 	return &c
// }
// func (tbh *ActDB) NewUser() *User {
// 	u := User{}
// 	return &u
// }
// func (tbh *ActDB) NewProjectComment(project *Project) *ProjectComment {
// 	pc := ProjectComment{}
// 	return &pc
// }
// func (tbh *ActDB) NewTaskComment(task *Task) *TaskComment {
// 	tc := TaskComment{}
// 	return &tc
// }

// // Load populates the db with the file
// func (tdb *ActDB) Load(filename string) bool {
// 	db, err := sql.Open("sqlite3", filename)
// 	checkErr(err)
// 	tdb.db = db

// 	// cdb.Filename = goutils.EvaluateFilename(filename)
// 	// cdb.PublicKeyFilename = goutils.EvaluateFilename(pubKey)
// 	// cdb.PrivateKeyFilename = goutils.EvaluateFilename(privKey)
// 	// cdb.EncryptionEnabled = encryptionEnabled

// 	// if !goutils.FileExists(cdb.Filename) {
// 	// 	db := DB{}
// 	// 	db.Entries = make(map[string]DBEntry)
// 	// 	cdb.data = db
// 	// } else {
// 	// 	jsonFile, err := os.Open(cdb.Filename)
// 	// 	if err != nil {
// 	// 		fmt.Printf("ERR %v\n", err)
// 	// 		db := DB{}
// 	// 		db.Entries = make(map[string]DBEntry)
// 	// 		cdb.data = db
// 	// 		// panic(err)
// 	// 	} else {
// 	// 		db := DB{}
// 	// 		bytes, _ := ioutil.ReadAll(jsonFile)
// 	// 		var data map[string]DBEntry
// 	// 		json.Unmarshal(bytes, &data)
// 	// 		db.Entries = data
// 	// 		cdb.data = db
// 	// 	}
// 	// }

// 	return true
// }

// func (tdb *ActDB) Init() bool {
// 	db := tdb.db
// 	sqls := strings.Split(SQL_SCHEMA, ";")
// 	for _, value := range sqls {
// 		// value = strings.ReplaceAll(value, "\n", " ")
// 		if strings.Trim(value, " \n") == "" {
// 			continue
// 		}
// 		value = strings.Trim(value, " ") + ";"
// 		value = strings.ReplaceAll(value, "\n", "")
// 		if value != "" && strings.Index(value, "--") != 0 {
// 			// fmt.Printf("Not a comment, -- index = %v\n", strings.Index(value, "--"))
// 			// fmt.Printf("\n%v\n", value)
// 			stmt, err := db.Prepare(value)
// 			_, err = stmt.Exec()
// 			checkErr(err)
// 		}
// 	}

// 	tdb.AddConfig("created", time.Now().Format(time.RFC3339Nano))
// 	tdb.AddConfig("version", VERSION)
// 	return true
// }

// func (tdb *ActDB) AddConfig(name string, value string) {
// 	db := tdb.db
// 	sql := fmt.Sprintf("insert into config (name, value) values (\"%v\", \"%v\");", name, value)
// 	_, err := db.Exec(sql)
// 	// fmt.Printf("%v\n", result)
// 	checkErr(err)
// }

// // func DoSql(cli *goutils.CLI) {
// // 	db, err := sql.Open("sqlite3", "./foo.db")
// // 	checkErr(err)

// // 	stmt, err := db.Prepare(SQL_CREATE)
// // 	res, err := stmt.Exec()
// // 	checkErr(err)

// // 	db, err := sql.Open("sqlite3", filename)
// // 	checkErr(err)
// // 	tdb.db = db

// // }

// // Clear empties the db (without saving it)
// func (tdb *ActDB) Clear() {
// 	// cdb.data.Entries = make(map[string]DBEntry)
// }

// func (tdb *ActDB) AddTask(name string) {
// 	db := tdb.db
// 	stmt, err := db.Prepare("INSERT INTO tasks(user_id, project_id, created, updated, state, name, description, deleted, archived) values(?,?,?,?,?,?,?,?, ?)")
// 	checkErr(err)

// 	user_id := 1
// 	project_id := 1
// 	created := time.Now().Format(time.RFC3339Nano)
// 	updated := time.Now().Format(time.RFC3339Nano)

// 	fmt.Printf("New task created %v, updated %v\n", created, updated)

// 	state := "created"
// 	// name := name
// 	description := ""
// 	deleted := false
// 	archived := false
// 	_, err = stmt.Exec(user_id, project_id, created, updated, state, name, description, deleted, archived)
// 	checkErr(err)

// }

// func (tdb *ActDB) ListTasks() []*Task {
// 	db := tdb.db
// 	rows, err := db.Query("SELECT task_id, project_id, created, updated, due, name, state FROM tasks")
// 	checkErr(err)
// 	var task_id int
// 	var project_id int
// 	var created string
// 	var updated string
// 	var due string
// 	var name string
// 	var state string

// 	var results []*Task
// 	for rows.Next() {
// 		err = rows.Scan(&task_id, &project_id, &created, &updated, &due, &name, &state)
// 		fmt.Printf("created %v\n", created)
// 		checkErr(err)
// 		t := NewTask()
// 		t.Task_id = task_id
// 		t.Project_id = project_id

// 		ti, _ := time.Parse(time.RFC3339Nano, created)
// 		t.Created = &ti

// 		up, _ := time.Parse(time.RFC3339Nano, updated)
// 		t.Updated = &up

// 		due_dt, _ := time.Parse(time.RFC3339Nano, due)
// 		t.Due = &due_dt

// 		t.State = state

// 		// t.updated = updated
// 		t.Name = name
// 		results = append(results, t)
// 	}

// 	rows.Close() //good habit to close
// 	return results

// }

// func (tdb *ActDB) GetTaskById(taskId string) *Task {
// 	db := tdb.db
// 	rows, err := db.Query("SELECT task_id, project_id, created, name, state FROM tasks where task_id=?", taskId)
// 	checkErr(err)
// 	var task_id int
// 	var project_id int
// 	var created string
// 	var name string
// 	var state string
// 	rows.Next()
// 	err = rows.Scan(&task_id, &project_id, &created, &name, &state)
// 	if err != nil {
// 		return nil
// 	}
// 	checkErr(err)
// 	t := NewTask()
// 	t.Task_id = task_id
// 	t.Project_id = project_id
// 	t.State = state

// 	cr, _ := time.Parse(time.RFC3339Nano, created)
// 	t.Created = &cr

// 	t.Name = name
// 	rows.Close() //good habit to close
// 	return t

// }

// func (tdb *ActDB) Save(task *Task) {
// 	db := tdb.db
// 	if task.Task_id == 0 {
// 		tdb.AddTask(task.Name)
// 		return
// 	}
// 	updated := time.Now().Format(time.RFC3339Nano)

// 	stmt, err := db.Prepare("UPDATE tasks SET name=?, updated=? WHERE task_id = ?")
// 	checkErr(err)
// 	_, err = stmt.Exec(task.Name, updated, task.Task_id)
// 	checkErr(err)
// }

// func (tdb *ActDB) Demo() bool {

// 	db := tdb.db

// 	// insert
// 	stmt, err := db.Prepare("INSERT INTO userinfo(username, departname, created) values(?,?,?)")
// 	checkErr(err)

// 	res, err := stmt.Exec("astaxie", "研发部门", "2012-12-09")
// 	checkErr(err)

// 	id, err := res.LastInsertId()
// 	checkErr(err)

// 	fmt.Println(id)
// 	// update
// 	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
// 	checkErr(err)

// 	res, err = stmt.Exec("astaxieupdate", id)
// 	checkErr(err)

// 	affect, err := res.RowsAffected()
// 	checkErr(err)

// 	fmt.Println(affect)

// 	// query
// 	rows, err := db.Query("SELECT * FROM userinfo")
// 	checkErr(err)
// 	var uid int
// 	var username string
// 	var department string
// 	var created time.Time

// 	for rows.Next() {
// 		err = rows.Scan(&uid, &username, &department, &created)
// 		checkErr(err)
// 		fmt.Println(uid)
// 		fmt.Println(username)
// 		fmt.Println(department)
// 		fmt.Println(created)
// 	}

// 	rows.Close() //good habit to close

// 	// delete
// 	stmt, err = db.Prepare("delete from userinfo where uid=?")
// 	checkErr(err)

// 	res, err = stmt.Exec(id)
// 	checkErr(err)

// 	affect, err = res.RowsAffected()
// 	checkErr(err)

// 	fmt.Println(affect)

// 	db.Close()

// 	// trashSQL, err := database.Prepare("update task set is_deleted='Y',last_modified_at=datetime() where id=?")
// 	// if err != nil {
// 	//     fmt.Println(err)
// 	// }
// 	// tx, err := database.Begin()
// 	// if err != nil {
// 	//     fmt.Println(err)
// 	// }
// 	// _, err = tx.Stmt(trashSQL).Exec(id)
// 	// if err != nil {
// 	//     fmt.Println("doing rollback")
// 	//     tx.Rollback()
// 	// } else {
// 	//     tx.Commit()
// 	// }

// 	return true
// }
