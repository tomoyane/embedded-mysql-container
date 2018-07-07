package main

import (
	"database/sql"
)

var (
	driverName = "mysql"
	dataSourceName = "root:root@tcp(127.0.0.1:33306)/"
	db *sql.DB
)

func ConnectMysql() {
	db, openErr := sql.Open(driverName, dataSourceName)
	if openErr != nil {
		panic(openErr)
	}
	defer db.Close()
}

func CloseMysql() {
	db.Close()
}

func GetDriverName() string {
	return driverName
}

func GetDataSourceName() string {
	return dataSourceName
}