package main

import (
	_ "github.com/go-sql-driver/mysql"
)

type EmbeddedMysqlHandlerImpl struct {
	errorHandler *Error
}

func (e EmbeddedMysqlHandlerImpl) NewEmbeddedMysqlHandlerImpl() EmbeddedMysqlHandler {
	return &EmbeddedMysqlHandlerImpl{}
}

func (e EmbeddedMysqlHandlerImpl) CreateDatabase(query string) {
	_, err := db.Exec(query)
	if err != nil {
		e.errorHandler.ErrorMessage(
			"mysql create database failed.",
			err,
		)
	}
}

func (e EmbeddedMysqlHandlerImpl) CreateTable(query string) {
	_, err := db.Exec(query)
	if err != nil {
		e.errorHandler.ErrorMessage(
			"mysql create table failed.",
			err,
		)
	}
}

func (e EmbeddedMysqlHandlerImpl) DropTable(query string) {
	_, err := db.Exec(query)
	if err != nil {
		e.errorHandler.ErrorMessage(
			"mysql drop table failed.",
			err,
		)
	}
}

func (e EmbeddedMysqlHandlerImpl) InsertQuery(query string) {
	_, err := db.Exec(query)
	if err != nil {
		e.errorHandler.ErrorMessage(
			"mysql insert query failed.",
			err,
		)
	}
}

func (e EmbeddedMysqlHandlerImpl) UpdateQuery(query string) {
	_, err := db.Exec(query)
	if err != nil {
		e.errorHandler.ErrorMessage(
			"mysql update query failed.",
			err,
		)
	}
}

func (e EmbeddedMysqlHandlerImpl) DeleteQuery(query string) {
	_, err := db.Exec(query)
	if err != nil {
		e.errorHandler.ErrorMessage(
			"mysql delete query failed.",
			err,
		)
	}
}
