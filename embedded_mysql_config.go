package main

import "database/sql"

type EmbeddedMysqlConfig struct {
}

var mysqlResolverImpl = EmbeddedMysqlHandlerImpl{}.NewMysqlResolver()
var db *sql.DB

func (c EmbeddedMysqlConfig) Open() {
	db, openErr := sql.Open("mysql", "root:root@tcp(127.0.0.1:33306)/")
	if openErr != nil {
		panic(openErr)
	}
	defer db.Close()
}

func (c EmbeddedMysqlConfig) Close() {
	db.Close()
}

func (c EmbeddedMysqlConfig) addSchema(databaseName string)  {
	mysqlResolverImpl.Insert("CREATE DATABASE " + databaseName)
}
