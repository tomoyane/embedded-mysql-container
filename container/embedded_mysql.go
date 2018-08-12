package container

type MysqlConfig interface {
	AddSchema(dbName string)

	CreateTable(query string)
}
