package main

type EmbeddedMysqlHandler interface {
	CreateDatabase(query string) bool

	CreateTable(query string) bool

	DropTable(query string) bool

	InsertQuery(query string) bool

	UpdateQuery(query string) bool

	DeleteQuery(query string) bool
}
