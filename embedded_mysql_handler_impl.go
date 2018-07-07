package main

import (
	_ "github.com/go-sql-driver/mysql"
)

type EmbeddedMysqlHandlerImpl struct {
}

func (e EmbeddedMysqlHandlerImpl) NewEmbeddedMysqlHandlerImpl() EmbeddedMysqlHandler {
	return &EmbeddedMysqlHandlerImpl{}
}

func (e EmbeddedMysqlHandlerImpl) CreateDatabase(query string) {
	_, queryErr := db.Exec(query)
	if queryErr != nil {
		errorHandler.ErrorMessage(
			"mysql create database failed.",
			queryErr,
		)
	}
}

func (e EmbeddedMysqlHandlerImpl) CreateTable(query string) {
	_, queryErr := db.Exec(query)
	if queryErr != nil {
		errorHandler.ErrorMessage(
			"mysql create table failed.",
			queryErr,
		)
	}
}

func (e EmbeddedMysqlHandlerImpl) DropTable(query string) {
	_, queryErr := db.Exec(query)
	if queryErr != nil {
		errorHandler.ErrorMessage(
			"mysql drop table failed.",
			queryErr,
		)
	}
}

func (e EmbeddedMysqlHandlerImpl) InsertQuery(query string) {
	_, queryErr := db.Exec(query)
	if queryErr != nil {
		errorHandler.ErrorMessage(
			"mysql insert query failed.",
			queryErr,
		)
	}
}

func (e EmbeddedMysqlHandlerImpl) UpdateQuery(query string) {
	_, queryErr := db.Exec(query)
	if queryErr != nil {
		errorHandler.ErrorMessage(
			"mysql update query failed.",
			queryErr,
		)
	}
}

func (e EmbeddedMysqlHandlerImpl) DeleteQuery(query string) {
	_, queryErr := db.Exec(query)
	if queryErr != nil {
		errorHandler.ErrorMessage(
			"mysql delete query failed.",
			queryErr,
		)
	}
}
