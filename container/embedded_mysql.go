package container

import "database/sql"

func AddSchema(name string) {

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:33306)/")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_,err = db.Exec("CREATE DATABASE " + name)
	if err != nil {
		panic(err)
	}
}