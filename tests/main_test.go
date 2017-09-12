package main

import (
    "testing"
    "fmt"
    "github.com/karneles/friend-management/handler"
	validator "gopkg.in/go-playground/validator.v9"
    
    "github.com/facebookgo/inject"
    "github.com/jmoiron/sqlx"
)

var (
    rh handler.RootHandler
)

func setupDB() *sqlx.DB {
	db, err := sqlx.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", "ssi", "teramakuro", "devbox-carneles.dev.svc.cluster.local", "test"))
	if err != nil {
		fmt.Print(err)
	}
	db.SetMaxOpenConns(40)
	return db
}

func TestMain(m *testing.M) {
	validate := validator.New()
	db := setupDB()
	err := inject.Populate(db, validate, &rh)
	if err != nil {
		fmt.Print(err)
	} else {
		m.Run()
	}
}