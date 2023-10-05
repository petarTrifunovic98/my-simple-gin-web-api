package repository

import (
	"my-simple-gin-web-api/dbdriver"
	"my-simple-gin-web-api/entities"
)

type UserRepository interface {
	Save(user entities.User)
	FindAll() []entities.User
	FindByID(id any) entities.User
}

type database struct {
	db_client *dbdriver.MySimpleDatabaseClient
}

func NewUserRepository(client *dbdriver.MySimpleDatabaseClient) UserRepository {
	return &database{
		db_client: client,
	}
}

func (db *database) Save(user entities.User) {
	dbdriver.IssueInsertCommand[entities.User](db.db_client, user)
}

func (db *database) FindAll() []entities.User {
	return dbdriver.IssueSelectCommand[entities.User](db.db_client)
}

func (db *database) FindByID(id any) entities.User {
	return dbdriver.IssueSelectOneCommand[entities.User](db.db_client, id)
}
