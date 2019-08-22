package pgsql_conn

import (
	"Web-Spid-Demo/datamodels"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	PgSql *gorm.DB
	err   error
)

func init() {
	PgSql, err = gorm.Open("postgres", "host=127.0.0.1 port=5432 user=dollarkiller dbname=one password=123456 sslmode=disable")
	if err != nil {
		panic(err.Error())
	}

	// 数据库映射
	mapping()
}

func mapping() {
	PgSql.AutoMigrate(
		&datamodels.User{},
	)
}
