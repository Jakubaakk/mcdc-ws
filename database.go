package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"mcdc-ws/users"
)

func ConnectToDB() *gorm.DB {
	//dbUrl := "postgres://mcdc:mcdcpassword@localhost:55432/mcdc"
	dbUrl := "mcdc:mcdcpassword@tcp(localhost:3306)/mcdc?charset=utf8"
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
