package main

import (
	_ "github.com/go-sql-driver/mysql"
)

type EmbeddedMysqlHandlerImpl struct {}

var errorHandler = Error{}

func (m EmbeddedMysqlHandlerImpl) NewMysqlResolver() EmbeddedMysqlHandler {
	return &EmbeddedMysqlHandlerImpl{}
}

func (m EmbeddedMysqlHandlerImpl) CreateDatabase(query string) {
	_, queryErr := db.Exec(query)
	if queryErr != nil {
		errorHandler.ErrorMessage(
			"mysql create database failed.",
			queryErr,
		)
	}
}

func (m EmbeddedMysqlHandlerImpl) CreateTable(query string)  {
	_, queryErr := db.Exec(query)
	if queryErr != nil {
		errorHandler.ErrorMessage(
			"mysql create table failed.",
			queryErr,
		)
	}
}

func (m EmbeddedMysqlHandlerImpl) DropTable(query string) {
	_, queryErr := db.Exec(query)
	if queryErr != nil {
		errorHandler.ErrorMessage(
			"mysql drop table failed.",
			queryErr,
		)
	}
}

func (m EmbeddedMysqlHandlerImpl) Insert(query string) {
	_, queryErr := db.Exec(query)
	if queryErr != nil {
		errorHandler.ErrorMessage(
			"mysql create database failed.",
			queryErr,
		)
	}
}

func (m EmbeddedMysqlHandlerImpl) Update(query string) {
}

func (m EmbeddedMysqlHandlerImpl) Delete(query string) {
}

