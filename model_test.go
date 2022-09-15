package main

import (
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

var (
	mock sqlmock.Sqlmock
)

func TestMain(m *testing.M) {
	var db *sql.DB
	var err error
	db, mock, err = sqlmock.New()
	if err != nil {
		panic(err)
	}
	SetupDatabaseMock(db)
	m.Run()
}

func TestInsert(t *testing.T) {
	menuModel := Menu{}
	menuModel.ParentId = 0
	menuModel.Title = "test"
	menuModel.Name = "testName"
	menuModel.Sort = 1
	menuModel.CreatedAt = time.Now()
	menuModel.UpdatedAt = time.Now()

	mock.ExpectExec("INSERT INTO `menus`").WithArgs( menuModel.ParentId, menuModel.Title, menuModel.Name, menuModel.Sort, menuModel.Route, menuModel.Component, menuModel.Icon, AnyTime{}, AnyTime{}, nil).WillReturnResult(sqlmock.NewResult(0, 0))

	// now we execute our method
	if err := Insert(&menuModel); err != nil {
		t.Errorf("models.NewTMenuModel().Insert() error : %s", err.Error())
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
