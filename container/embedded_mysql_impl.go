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

	_, err = db.Exec("CREATE DATABASE " + name)
	if err != nil {
		m.errContainer.msg = "mysql  query failed."
		m.errContainer.error = cliErr
		m.errContainer.ErrorMessage()
	}

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

	_, err = db.Exec(query)
	if err != nil {
		m.errContainer.msg = "mysql query failed."
		m.errContainer.error = cliErr
		m.errContainer.ErrorMessage()
	}

	db.Close()
}
