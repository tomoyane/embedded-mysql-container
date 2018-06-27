package main

type MysqlHandler interface {
	CreateDatabase(query string)

	CreateTable(query string)

	DropTable(query string)

	Insert(query string)

	Update(query string)

	Delete(query string)
}
