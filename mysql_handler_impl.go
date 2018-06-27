package main

import (
	_ "github.com/go-sql-driver/mysql"
)

type MysqlHandlerImpl struct {}

func (m MysqlHandlerImpl) NewMysqlResolver() MysqlHandler {
	return &MysqlHandlerImpl{}
}

func (m MysqlHandlerImpl) CreateDatabase(query string) {

}

func (m MysqlHandlerImpl) CreateTable(query string)  {
}

func (m MysqlHandlerImpl) DropTable(query string) {

}

func (m MysqlHandlerImpl) Insert(query string) {
}

func (m MysqlHandlerImpl) Update(query string) {
}

func (m MysqlHandlerImpl) Delete(query string) {
}

