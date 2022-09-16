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

	mock.ExpectExec("INSERT INTO `menus`").WithArgs(menuModel.ParentId, menuModel.Title, menuModel.Name, menuModel.Sort, menuModel.Route, menuModel.Component, menuModel.Icon, AnyTime{}, AnyTime{}, nil).WillReturnResult(sqlmock.NewResult(0, 0))

	// now we execute our method
	if err := Insert(&menuModel); err != nil {
		t.Errorf("models.NewTMenuModel().Insert() error : %s", err.Error())
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetAllByParentId(t *testing.T) {
	mock.ExpectQuery("SELECT (.+) FROM `menus`").WithArgs(0).WillReturnRows(sqlmock.NewRows([]string{
		"id", "parentId", "title", "name", "sort", "route", "component", "icon", "createdAt", "updatedAt", "deletedAt",
	}).AddRow(
		1, 0, "menu1", "一级菜单1", 1, "/index", "", "", time.Now(), time.Now(), nil,
	).AddRow(
		2, 0, "menu2", "一级菜单2", 2, "/index", "", "", time.Now(), time.Now(), nil,
	))

	data, err := GetAllByParentId(0)
	if err != nil {
		t.Errorf("查询错误 %s", err)
	}
	if len(data) != 2 {
		t.Errorf("查询数据异常!")
	}
}
