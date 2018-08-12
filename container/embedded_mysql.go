package container

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func AddSchema(name string) {

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:33306)/")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	time.Sleep(10000)
	db.Exec("CREATE DATABASE " + name)

	db.Close()
}
