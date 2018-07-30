package container

import (
	_ "github.com/go-sql-driver/mysql"
)

type EmbeddedMysqlHandlerImpl struct {
}

func (e EmbeddedMysqlHandlerImpl) NewEmbeddedMysqlHandlerImpl() EmbeddedMysqlHandler {
	return &EmbeddedMysqlHandlerImpl{}
}

func (e EmbeddedMysqlHandlerImpl) CreateDatabase(query string) bool {
	_, err := Db.Exec(query)
	if err != nil {
		errContainer := ErrorContainer{
			msg: "mysql create database failed.",
			error: err,
		}

		errContainer.ErrorMessage()

		return false
	}

	return true
}

func (e EmbeddedMysqlHandlerImpl) CreateTable(query string) bool {
	_, err := Db.Exec(query)
	if err != nil {
		errContainer := ErrorContainer{
			msg: "mysql create table failed.",
			error: err,
		}

		errContainer.ErrorMessage()

		return false
	}

	return true
}

func (e EmbeddedMysqlHandlerImpl) DropTable(query string) bool {
	_, err := Db.Exec(query)
	if err != nil {
		errContainer := ErrorContainer{
			msg: "mysql drop table failed.",
			error: err,
		}

		errContainer.ErrorMessage()

		return false
	}

	return true
}

func (e EmbeddedMysqlHandlerImpl) InsertQuery(query string) bool {
	_, err := Db.Exec(query)
	if err != nil {
		errContainer := ErrorContainer{
			msg: "mysql insert query failed.",
			error: err,
		}

		errContainer.ErrorMessage()

		return false
	}

	return true
}

func (e EmbeddedMysqlHandlerImpl) UpdateQuery(query string) bool {
	_, err := Db.Exec(query)
	if err != nil {
		errContainer := ErrorContainer{
			msg: "mysql update query failed.",
			error: err,
		}

		errContainer.ErrorMessage()

		return false
	}

	return true
}

func (e EmbeddedMysqlHandlerImpl) DeleteQuery(query string) bool {
	_, err := Db.Exec(query)
	if err != nil {
		errContainer := ErrorContainer{
			msg: "mysql delete query failed.",
			error: err,
		}

		errContainer.ErrorMessage()

		return false
	}

	return true
}
