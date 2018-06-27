package main

import (
	_ "github.com/go-sql-driver/mysql"
)

type EmbeddedMysqlHandlerImpl struct {}

func (m EmbeddedMysqlHandlerImpl) NewMysqlResolver() EmbeddedMysqlHandler {
	return &EmbeddedMysqlHandlerImpl{}
}

func (m EmbeddedMysqlHandlerImpl) CreateDatabase(query string) {

}

func (m EmbeddedMysqlHandlerImpl) CreateTable(query string)  {
}

func (m EmbeddedMysqlHandlerImpl) DropTable(query string) {

}

func (m EmbeddedMysqlHandlerImpl) Insert(query string) {
}

func (m EmbeddedMysqlHandlerImpl) Update(query string) {
}

func (m EmbeddedMysqlHandlerImpl) Delete(query string) {
}

