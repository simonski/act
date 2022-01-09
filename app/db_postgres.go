package app

// import (
// 	"database/sql"
// 	"fmt"

// 	_ "github.com/lib/pq"
// )

// const SQL_SCHEMAX = `

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

// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "postgres"
// 	password = "password"
// 	dbname   = "idendb"
// )

// func main() {
// 	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
// 		"password=%s dbname=%s sslmode=disable",
// 		host, port, user, password, dbname)

// 	db, err := sql.Open("postgres", psqlInfo)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer db.Close()

// 	err = db.Ping()
// 	if err != nil {
// 		panic(err)
// 	}

// }
