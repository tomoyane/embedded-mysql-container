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

func (e EmbeddedMysqlHandlerImpl) CreateDatabase(query string) bool {
	_, err := Db.Exec(query)
	if err != nil {
		e.errorHandler.ErrorMessage("mysql create database failed.", err)
		return false
	}

	return true
}

func (e EmbeddedMysqlHandlerImpl) CreateTable(query string) bool {
	_, err := Db.Exec(query)
	if err != nil {
		e.errorHandler.ErrorMessage("mysql create table failed.", err)
		return false
	}

	return true
}

func (e EmbeddedMysqlHandlerImpl) DropTable(query string) bool {
	_, err := Db.Exec(query)
	if err != nil {
		e.errorHandler.ErrorMessage("mysql drop table failed.", err)
		return false
	}

	return true
}

func (e EmbeddedMysqlHandlerImpl) InsertQuery(query string) bool {
	_, err := Db.Exec(query)
	if err != nil {
		e.errorHandler.ErrorMessage("mysql insert query failed.", err)
		return false
	}

	return true
}

func (e EmbeddedMysqlHandlerImpl) UpdateQuery(query string) bool {
	_, err := Db.Exec(query)
	if err != nil {
		e.errorHandler.ErrorMessage("mysql update query failed.", err)
		return false
	}

	return true
}

func (e EmbeddedMysqlHandlerImpl) DeleteQuery(query string) bool {
	_, err := Db.Exec(query)
	if err != nil {
		e.errorHandler.ErrorMessage("mysql delete query failed.", err)
		return false
	}

	return true
}
