package api

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func init() {
	var sql_create string
	var create_db bool = false
	if _, err := os.Stat("forum.db"); err != nil {
		sql_raw, err := os.ReadFile("forum.sql")
		if err != nil {
			fmt.Println(err)
			return
		}
		sql_example, err := os.ReadFile("example.sql")
		if err != nil {
			fmt.Println(err)
			return
		}
		sql_create = string(sql_raw) + string(sql_example)
		create_db = true
	}
	var err error
	DB, err = sql.Open("sqlite3", "forum.db")
	if err != nil {
		panic(err)
	}
	//defer DB.Close()
	//DB.SetMaxOpenConns(10)
	//DB.SetMaxIdleConns(5)
	if create_db {
		_, err := DB.Exec(sql_create)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
