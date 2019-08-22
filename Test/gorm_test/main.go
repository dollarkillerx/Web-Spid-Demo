package main

import (
	"Web-Spid-Demo/repositories"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
	gorm.Model
	User     string
	Password string
	Salt     string
}

func main() {
	//Pgsql, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=dollarkiller dbname=one password=123456 sslmode=disable")
	//if err != nil {
	//	panic(err.Error())
	//}
	//
	//Pgsql.AutoMigrate(&User{})
	repositories.User()
}
