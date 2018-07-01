package main

import (
	"database/sql"
)

type EmbeddedMysqlConfig struct {
}

var (
	mysqlResolverImpl = EmbeddedMysqlHandlerImpl{}.NewMysqlResolver()
	driverName = "mysql"
	dataSourceName = "root:root@tcp(127.0.0.1:33306)/"
	db *sql.DB
)

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
