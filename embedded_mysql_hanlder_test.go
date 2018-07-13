package main

import "testing"

var mysqlHandler = EmbeddedMysqlHandlerImpl{}.NewEmbeddedMysqlHandlerImpl()

func TestCreateDatabase(t *testing.T) {
	query := "CREATE DATABASE test"
	result := mysqlHandler.CreateDatabase(query)
	if !result {
		t.Failed()
	}
}

func TestCreateTable(t *testing.T) {
	query := "CREATE TABLE test_table"
	result := mysqlHandler.CreateTable(query)
	if !result {
		t.Failed()
	}
}

func TestInsert(t *testing.T) {
}
