package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"mcdc-ws/configuration"
	"mcdc-ws/users"
)

func ConnectToDB(config *configuration.McdcConfiguration) *gorm.DB {
	//dbUrl := "postgres://mcdc:mcdcpassword@localhost:55432/mcdc"
	//dbUrl := "mcdc:mcdcpassword@tcp(localhost:3306)/mcdc?charset=utf8"
	dbUrl := getDbUrl(config)
	db, err := gorm.Open(mysql.Open(dbUrl), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("Cannot connect to DB")
	}
	err = db.AutoMigrate(&users.User{})
	if err != nil {
		panic(err)
	}

	return db
}

func getDbUrl(config *configuration.McdcConfiguration) string {
	db := config.Database
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", db.User, db.Password, db.Host, db.Port, db.Database)
}
