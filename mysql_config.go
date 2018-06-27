package main

import "database/sql"

type MysqlConfig struct {
}

var mysqlResolverImpl = MysqlHandlerImpl{}.NewMysqlResolver()

func (c MysqlConfig) Open() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:33306)/")
	if err != nil {
		panic(err)
	}
	defer db.Close()
}

func (c MysqlConfig) Close() {

}

func (c MysqlConfig) addSchema(databaseName string)  {
	mysqlResolverImpl.Insert("CREATE DATABASE " + databaseName)
}
