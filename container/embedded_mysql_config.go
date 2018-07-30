package container

import (
	"database/sql"
)

type EmbeddedMysqlConfig struct {
}

var (
	driverName = "mysql"
	dataSourceName = "root:root@tcp(127.0.0.1:33306)/"
	dataBaseName string
	Db *sql.DB
)

func (c EmbeddedMysqlConfig) ConnectMysql(dbName string) {
	if dataBaseName == "" {
		errContainer := ErrorContainer{
			msg: "Need to set database name",
			error: nil,
		}

		errContainer.ErrorMessage()

		panic(nil)
	}

	db, err := sql.Open(driverName, dataSourceName + dataBaseName)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	Db = db
}

func (c EmbeddedMysqlConfig) CloseMysql() {
	Db.Close()
}

func (c EmbeddedMysqlConfig) SetDbName(dbName string) {
	dataBaseName = dbName
}

func (c EmbeddedMysqlConfig) GetDriverName() string {
	return driverName
}

func (c EmbeddedMysqlConfig) GetDataSourceName() string {
	return dataSourceName + dataBaseName
}