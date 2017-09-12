package main

import (
	"fmt"
	"net/http"

	"./config"
	"./errorcode"
	"./handler"
	"./router"
	"./libs/apierror"
	"./libs/logger"
	validator "gopkg.in/go-playground/validator.v9"

	"github.com/facebookgo/inject"
	"github.com/jmoiron/sqlx"
)

func setupDB() *sqlx.DB {
	conf := config.GetConfig()
	db, err := sqlx.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", conf.MysqlUsername, conf.MysqlPassword, conf.MysqlHost, conf.MysqlDatabase))
	if err != nil {
		logger.Err("%v", err)
	}
	db.SetMaxOpenConns(conf.MysqlConnectionLimit)
	return db
}

func main() {
	// Prepare logger, and error
	conf := config.GetConfig()
	logger.Info("%v", conf)
	apierror.Setup(errorcode.GetErrorToHTPPStatusMap())

	// Setup dependency injection
	var rh handler.RootHandler
	validate := validator.New()
	db := setupDB()
	err := inject.Populate(db, validate, &rh)
	if err != nil {
		logger.Err("%v", err)
	}

	// Setup router
	r := router.CreateRouter(rh)

	// Serve
	logger.Info("Friend management started in Port: " + conf.Port)
	err = http.ListenAndServe(":"+conf.Port, r)
	if err != nil {
		logger.Err("%v", err)
	}
}
