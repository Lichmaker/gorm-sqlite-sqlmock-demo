package main

import (
	"database/sql"

	sqlitemock "github.com/lichmaker/sqlite-mock"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupDatabaseMock(mockDB *sql.DB) {
	myMock := sqlitemock.Open(mockDB)

	dbSqlite, err := gorm.Open(myMock, &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic("failed to connect mock database :" + err.Error())
	}
	DB = dbSqlite
}
