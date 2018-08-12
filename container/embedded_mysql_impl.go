package container

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type MysqlConfigImpl struct {
	errContainer ErrorContainer
}

func (m MysqlConfigImpl) New() MysqlConfig {
	return &MysqlConfigImpl{}
}

func (m MysqlConfigImpl) AddSchema(name string) {

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:33306)/")
	if err != nil {
		m.errContainer.msg = "mysql connection failed."
		m.errContainer.error = cliErr
		m.errContainer.ErrorMessage()
	}
	defer db.Close()

	db.Exec("CREATE DATABASE " + name)
	db.Close()
}

func (m MysqlConfigImpl) CreateTable(query string) {

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:33306)/")
	if err != nil {
		m.errContainer.msg = "mysql connection failed."
		m.errContainer.error = cliErr
		m.errContainer.ErrorMessage()
	}
	defer db.Close()

	db.Exec(query)
	db.Close()
}
