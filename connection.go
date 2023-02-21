package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	DB_HOST     = "buzzwomendatabase-new.cixgcssswxvx.ap-south-1.rds.amazonaws.com"
	DB_USER     = "bdms_staff_admin"
	DB_PASSWORD = "sfhakjfhyiqundfgs3765827635"
	DB_NAME     = "bdms_staff"
)

func Connection() {
	db, err := sql.Open("mysql", DB_USER+":"+DB_PASSWORD+"@tcp("+DB_HOST+")/"+DB_NAME)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer db.Close()

	fmt.Println("Successfully connected to database!")
}
