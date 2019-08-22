package repositories

import (
	"Web-Spid-Demo/datamodels"
	"Web-Spid-Demo/datasources/pgsql_conn"
	"log"
)

func User() {
	pgsql_conn.PgSql.Create(&datamodels.User{
		User:"dollarkiller",
		Password:"123456",
	})

	// 读取
	var user datamodels.User
	err := pgsql_conn.PgSql.First(&user, 1).Error
	if err != nil {
		panic(err.Error())
	}
	log.Println(user)

	e := pgsql_conn.PgSql.Model(&user).Update("password", "asdadad").Error
	if e != nil {
		panic(e.Error())
	}

	err = pgsql_conn.PgSql.First(&user, "user = ?", "dollarkillers").Error
	if err != nil {
		panic(err.Error())
	}
	log.Println(user)



}
