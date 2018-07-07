package main

type EmbeddedMysqlHandler interface {
	CreateDatabase(query string)

	CreateTable(query string)

	DropTable(query string)

	InsertQuery(query string)

	UpdateQuery(query string)

	DeleteQuery(query string)
}
