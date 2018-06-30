package main

import "database/sql"

type EmbeddedMysqlConfig struct {
}

var mysqlResolverImpl = EmbeddedMysqlHandlerImpl{}.NewMysqlResolver()
var db *sql.DB
var driverName = "mysql"
var dataSourceName = "root:root@tcp(127.0.0.1:33306)/"

func (c EmbeddedMysqlConfig) ConnectMysql() {
	db, openErr := sql.Open(driverName, dataSourceName)
	if openErr != nil {
		panic(openErr)
	}
	defer db.Close()
}

func (c EmbeddedMysqlConfig) CloseMysql() {
	db.Close()
}

func (c EmbeddedMysqlConfig) addSchema(databaseName string)  {
	mysqlResolverImpl.Insert("CREATE DATABASE " + databaseName)
}	
