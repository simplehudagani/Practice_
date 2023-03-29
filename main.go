package main

import (
	// dbs "sum/connection"
	"net/http"
	a "sum/list"
)

// var db = dbs.Connection()

func main() {
	// Create a new MySQL database connection
	// db, err := sql.Open("mysql", "bdms_staff_admin:sfhakjfhyiqundfgs3765827635@tcp(buzzwomendatabase-new.cixgcssswxvx.ap-south-1.rds.amazonaws.com:3306)/bdms_staff?charset=utf8")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Connected to MySQL database")
	// defer db.Close()

	http.HandleFunc("/spoorthilist", a.Spoorthilist)
	http.HandleFunc("/buzzlist", a.Buzzlist)
	http.HandleFunc("/greenlist", a.Greenlist)
	http.HandleFunc("/spoorthiattendance", a.Spoorthiattendace)
	http.HandleFunc("/buzzattendance", a.Buzzattendance)
	http.HandleFunc("/greenattendance", a.Greenattendance)
	//http.HandleFunc("/dashbord/vypar", a.GetReportingOpsManagers)
	//http.HandleFunc("/greenlist", a.Buzzlist)
	http.ListenAndServe(":5002", nil)

}
