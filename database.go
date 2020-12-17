package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)
import "github.com/jinzhu/gorm"

var Db *gorm.DB

func database() {
	db, err := gorm.Open("mysql", "root:222@/message?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&User{}, &Msg{})
	Db = db
}
