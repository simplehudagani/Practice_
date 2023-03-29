package connection

import (
	"database/sql"
	"fmt"
	"log"
)

var db *sql.DB

func Connection() *sql.DB {
	// Create a new MySQL database connection
	db, err := sql.Open("mysql", "bdms_staff_admin:sfhakjfhyiqundfgs3765827635@tcp(buzzwomendatabase-new.cixgcssswxvx.ap-south-1.rds.amazonaws.com:3306)/bdms_staff?charset=utf8")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MySQL database")
	//defer db.Close()
	return db

}
