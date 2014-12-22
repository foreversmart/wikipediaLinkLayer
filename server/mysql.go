package server

import (
	"database/sql"
    _ "github.com/go-sql-driver/mysql"
    "log"
    "fmt"
)

var db *sql.DB

const (
	userTable string = "user"
    wordTable string = "word"
    pageTable string = "page"
)

func DatabaseInit(){
	db1, err := sql.Open("mysql", "root:33a55c67@tcp(127.0.0.1:3306)/cohesion")
    if err != nil {
        checkErr(err)  // Just for example purpose. You should use proper error handling instead of panic
    }
    if db1 == nil {
        fmt.Println("数据库链接不到")
    }
    db = db1
}

func checkErr(err error) {
    if err != nil {
    	log.Println(err)
        fmt.Println(err)
        panic(err)
    }
}
